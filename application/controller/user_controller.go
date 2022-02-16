package controller

import (
	"event-bus-demo/application/dto"
	"event-bus-demo/application/error"
	"event-bus-demo/domain/model"
	"event-bus-demo/domain/service"
	"event-bus-demo/infrastructure/event_sourcing"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"net/http"
)

type UserController interface {
	GetUserByID(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUserPassword(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	eventBus         event_sourcing.EventBus
	userReadService  service.UserReadService
	controllerAdvice error.ControllerAdvice
}

func NewUserController(eventBus event_sourcing.EventBus, userReadService service.UserReadService, controllerAdvice error.ControllerAdvice) UserController {
	return &userController{
		eventBus:         eventBus,
		userReadService:  userReadService,
		controllerAdvice: controllerAdvice,
	}
}

func (controller *userController) GetUserByID(ctx *gin.Context) {
	if ID, err := uuid.Parse(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID parameter must be a valid UUID value",
		})
	} else if response, err := controller.userReadService.GetUserByID(model.GetUserByIDEvent{ID: ID}); err != nil {
		httpError := controller.controllerAdvice.TranslateError(err)
		ctx.JSON(httpError.GetCode(), gin.H{
			"message": httpError.GetMessage(),
		})
	} else {
		ctx.JSON(http.StatusOK, response)
	}
}

func (controller *userController) CreateUser(ctx *gin.Context) {
	var request dto.CreateUserRequest
	err := ctx.ShouldBindWith(&request, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else if id, err := uuid.NewRandom(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		controller.eventBus.Publish(request.ToEvent(id))
		response := dto.CreateToDoResponse{
			ID: id,
		}
		ctx.JSON(http.StatusCreated, response)
	}
}

func (controller *userController) UpdateUserPassword(ctx *gin.Context) {
	var request dto.UpdateUserPasswordRequest
	if ID, err := uuid.Parse(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID parameter must be a valid UUID value",
		})
	} else if err := ctx.ShouldBindWith(&request, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else if _, err := controller.userReadService.GetUserByID(model.GetUserByIDEvent{ID: ID}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		controller.eventBus.Publish(request.ToEvent(ID))
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}

func (controller *userController) DeleteUser(ctx *gin.Context) {
	if ID, err := uuid.Parse(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID parameter must be a valid UUID value",
		})
	} else {
		controller.eventBus.Publish(model.DeleteUserEvent{ID: ID})
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}

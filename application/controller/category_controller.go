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
	"log"
	"net/http"
)

type CategoryController interface {
	GetCategories(ctx *gin.Context)
	GetCategoryById(ctx *gin.Context)
	SaveCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
}

type categoryController struct {
	eventChannel        event_sourcing.EventBusChannel
	eventBus            event_sourcing.EventBus
	categoryReadService service.CategoryReadService
	controllerAdvice    error.ControllerAdvice
}

func NewCategoryController(eventBus event_sourcing.EventBus, categoryReadService service.CategoryReadService, controllerAdvice error.ControllerAdvice) CategoryController {
	return &categoryController{
		eventBus:            eventBus,
		categoryReadService: categoryReadService,
		controllerAdvice:    controllerAdvice,
	}
}

func (controller *categoryController) GetCategories(ctx *gin.Context) {
	categories, err := controller.categoryReadService.GetCategories()
	if err != nil {
		log.Println(err)
		httpError := controller.controllerAdvice.TranslateError(err)
		ctx.JSON(httpError.GetCode(), gin.H{
			"message": httpError.GetMessage(),
		})
	} else {
		ctx.JSON(http.StatusOK, categories)
	}
}

func (controller *categoryController) GetCategoryById(ctx *gin.Context) {
	if ID, err := uuid.Parse(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID parameter must be a valid UUID value",
		})
	} else if category, err := controller.categoryReadService.GetCategoryById(model.GetCategoryByIDEvent{ID: ID}); err != nil {
		httpError := controller.controllerAdvice.TranslateError(err)
		ctx.JSON(httpError.GetCode(), gin.H{
			"message": httpError.GetMessage(),
		})
	} else {
		ctx.JSON(http.StatusOK, category)
	}
}

func (controller *categoryController) SaveCategory(ctx *gin.Context) {
	var request dto.CreateCategoryRequest
	err := ctx.ShouldBindWith(&request, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else if ID, err := uuid.NewRandom(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		controller.eventBus.Publish(request.ToEvent(ID))
		response := dto.CreateCategoryResponse{
			ID: ID,
		}
		ctx.JSON(http.StatusCreated, response)
	}
}

func (controller *categoryController) UpdateCategory(ctx *gin.Context) {
	var request dto.UpdateCategoryRequest
	if ID, err := uuid.Parse(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID parameter must be a valid UUID value",
		})
	} else if err := ctx.ShouldBindWith(&request, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else if _, err := controller.categoryReadService.GetCategoryById(model.GetCategoryByIDEvent{ID: ID}); err != nil {
		httpError := controller.controllerAdvice.TranslateError(err)
		ctx.JSON(httpError.GetCode(), gin.H{
			"message": httpError.GetMessage(),
		})
	} else {
		controller.eventBus.Publish(request.ToEvent(ID))
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}

func (controller *categoryController) DeleteCategory(ctx *gin.Context) {
	if ID, err := uuid.Parse(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID parameter must be a valid UUID value",
		})
	} else {
		event := model.DeleteCategoryEvent{
			ID: ID,
		}
		controller.eventBus.Publish(event)
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}

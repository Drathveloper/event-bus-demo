package controller

import (
	"event-bus-demo/application/dto"
	"event-bus-demo/application/error"
	"event-bus-demo/domain/model"
	"event-bus-demo/domain/service"
	"event-bus-demo/infrastructure/event_sourcing"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"net/http"
)

type ToDoController interface {
	GetToDoList(ctx *gin.Context)
	GetToDoById(ctx *gin.Context)
	SaveToDo(ctx *gin.Context)
	UpdateToDo(ctx *gin.Context)
	DeleteToDo(ctx *gin.Context)
	RemoveCategoriesFromToDo(ctx *gin.Context)
	AddCategoriesIntoToDo(ctx *gin.Context)
}

type toDoController struct {
	eventBus            event_sourcing.EventBus
	toDoReadService     service.ToDoReadService
	categoryReadService service.CategoryReadService
	controllerAdvice    error.ControllerAdvice
}

func NewTodoController(eventBus event_sourcing.EventBus, toDoReadService service.ToDoReadService, categoryReadService service.CategoryReadService, controllerAdvice error.ControllerAdvice) ToDoController {
	return &toDoController{
		eventBus:            eventBus,
		toDoReadService:     toDoReadService,
		categoryReadService: categoryReadService,
		controllerAdvice:    controllerAdvice,
	}
}

func (controller *toDoController) GetToDoList(ctx *gin.Context) {
	todoList, err := controller.toDoReadService.GetAllToDo()
	if err != nil {
		httpError := controller.controllerAdvice.TranslateError(err)
		ctx.JSON(httpError.GetCode(), gin.H{
			"message": httpError.GetMessage(),
		})
	} else {
		ctx.JSON(http.StatusOK, todoList)
	}
}

func (controller *toDoController) GetToDoById(ctx *gin.Context) {
	if ID, err := uuid.Parse(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID parameter must be a valid UUID value",
		})
	} else if response, err := controller.toDoReadService.GetToDo(model.GetToDoEvent{ID: ID}); err != nil {
		httpError := controller.controllerAdvice.TranslateError(err)
		ctx.JSON(httpError.GetCode(), gin.H{
			"message": httpError.GetMessage(),
		})
	} else {
		ctx.JSON(http.StatusOK, response)
	}
}

func (controller *toDoController) SaveToDo(ctx *gin.Context) {
	var request dto.CreateToDoRequest
	err := ctx.ShouldBindWith(&request, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else if id, err := uuid.NewRandom(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else if _, err := controller.categoryReadService.GetCategoriesByIds(request.Categories); err != nil {
		appErr := controller.controllerAdvice.TranslateError(err)
		ctx.JSON(appErr.GetCode(), gin.H{
			"message": appErr.GetMessage(),
		})
	} else {
		controller.eventBus.Publish(request.ToEvent(id))
		response := dto.CreateToDoResponse{
			ID: id,
		}
		ctx.JSON(http.StatusCreated, response)
	}
}

func (controller *toDoController) UpdateToDo(ctx *gin.Context) {
	var request dto.UpdateToDoRequest
	if ID, err := uuid.Parse(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID parameter must be a valid UUID value",
		})
	} else if err := ctx.ShouldBindWith(&request, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else if _, err := controller.toDoReadService.GetToDo(model.GetToDoEvent{ID: ID}); err != nil {
		httpError := controller.controllerAdvice.TranslateError(err)
		ctx.JSON(httpError.GetCode(), gin.H{
			"message": httpError.GetMessage(),
		})
	} else {
		controller.eventBus.Publish(request.ToEvent(ID))
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}

func (controller *toDoController) DeleteToDo(ctx *gin.Context) {
	if ID, err := uuid.Parse(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID parameter must be a valid UUID value",
		})
	} else {
		event := model.DeleteToDoEvent{
			ID: ID,
		}
		controller.eventBus.Publish(event)
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}

func (controller *toDoController) RemoveCategoriesFromToDo(ctx *gin.Context) {
	var request dto.RemoveCategoriesFromToDoRequest
	if ID, err := uuid.Parse(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID parameter must be a valid UUID value",
		})
	} else if err := ctx.ShouldBindWith(&request, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else if ok, err := controller.toDoReadService.IsToDoAlreadyInCategories(ID, request.CategoriesID); err != nil || !ok {
		if err != nil {
			appErr := controller.controllerAdvice.TranslateError(err)
			ctx.JSON(appErr.GetCode(), gin.H{
				"message": appErr.GetMessage(),
			})
		} else {
			ctx.JSON(http.StatusPreconditionFailed, gin.H{
				"message": fmt.Sprintf("given ToDo is not registered in one or more of given categories"),
			})
		}
	} else if _, err := controller.categoryReadService.GetCategoriesByIds(request.CategoriesID); err != nil {
		appErr := controller.controllerAdvice.TranslateError(err)
		ctx.JSON(appErr.GetCode(), gin.H{
			"message": appErr.GetMessage(),
		})
	} else {
		event := model.RemoveCategoriesFromToDoEvent{
			ToDoID:     ID,
			Categories: request.CategoriesID,
		}
		controller.eventBus.Publish(event)
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}

func (controller *toDoController) AddCategoriesIntoToDo(ctx *gin.Context) {
	var request dto.AddCategoriesFromToDoRequest
	if ID, err := uuid.Parse(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID parameter must be a valid UUID value",
		})
	} else if err := ctx.ShouldBindWith(&request, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else if ok, err := controller.toDoReadService.IsToDoAlreadyInCategories(ID, request.CategoriesID); err != nil || ok {
		if err != nil {
			appErr := controller.controllerAdvice.TranslateError(err)
			ctx.JSON(appErr.GetCode(), gin.H{
				"message": appErr.GetMessage(),
			})
		} else {
			ctx.JSON(http.StatusPreconditionFailed, gin.H{
				"message": fmt.Sprintf("given ToDo is already registered in one or more of given categories"),
			})
		}
	} else if _, err := controller.categoryReadService.GetCategoriesByIds(request.CategoriesID); err != nil {
		appErr := controller.controllerAdvice.TranslateError(err)
		ctx.JSON(appErr.GetCode(), gin.H{
			"message": appErr.GetMessage(),
		})
	} else {
		event := model.AddCategoriesFromToDoEvent{
			ToDoID:     ID,
			Categories: request.CategoriesID,
		}
		controller.eventBus.Publish(event)
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}

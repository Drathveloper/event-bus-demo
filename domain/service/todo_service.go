package service

import (
	"event-bus-demo/application/dto"
	"event-bus-demo/domain/error"
	"event-bus-demo/domain/mapper"
	"event-bus-demo/domain/model"
	"event-bus-demo/infrastructure/database/service"
	"event-bus-demo/infrastructure/util"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

type ToDoReadService interface {
	GetAllToDo() (dto.GetAllToDoResponse, error.DomainError)
	GetToDo(event model.GetToDoEvent) (dto.GetToDoResponse, error.DomainError)
	IsToDoAlreadyInCategories(ID uuid.UUID, categories []uuid.UUID) (bool, error.DomainError)
}

type ToDoWriteService interface {
	AddToDo(event model.CreateToDoEvent) error.DomainError
	UpdateToDo(event model.UpdateToDoEvent) error.DomainError
	DeleteToDo(event model.DeleteToDoEvent) error.DomainError
	AddCategoriesIntoToDo(event model.AddCategoriesFromToDoEvent) error.DomainError
	RemoveCategoriesFromToDo(event model.RemoveCategoriesFromToDoEvent) error.DomainError
}

type toDoReadService struct {
	logger              *zap.Logger
	toDoDatabaseService service.ToDoDatabaseService
	domainAdvice        error.DomainAdvice
}

type toDoWriteService struct {
	logger              *zap.Logger
	toDoDatabaseService service.ToDoDatabaseService
	domainAdvice        error.DomainAdvice
}

func NewToDoReadService(toDoDatabaseService service.ToDoDatabaseService, domainAdvice error.DomainAdvice, logger *zap.Logger) ToDoReadService {
	return &toDoReadService{
		toDoDatabaseService: toDoDatabaseService,
		logger:              logger,
		domainAdvice:        domainAdvice,
	}
}

func NewToDoWriteService(toDoDatabaseService service.ToDoDatabaseService, domainAdvice error.DomainAdvice, logger *zap.Logger) ToDoWriteService {
	return &toDoWriteService{
		toDoDatabaseService: toDoDatabaseService,
		logger:              logger,
		domainAdvice:        domainAdvice,
	}
}

func (service *toDoReadService) GetAllToDo() (dto.GetAllToDoResponse, error.DomainError) {
	toDoList, err := service.toDoDatabaseService.GetAllToDoList()
	return mapper.NewGetAllToDoResponseFromDomainModel(toDoList), service.domainAdvice.TranslateError(err)
}

func (service *toDoReadService) GetToDo(event model.GetToDoEvent) (dto.GetToDoResponse, error.DomainError) {
	toDo, err := service.toDoDatabaseService.GetToDo(event.ID)
	return mapper.NewGetToDoResponseFromDomainModel(toDo), service.domainAdvice.TranslateError(err)
}

func (service *toDoReadService) IsToDoAlreadyInCategories(ID uuid.UUID, categories []uuid.UUID) (bool, error.DomainError) {
	toDo, err := service.toDoDatabaseService.GetToDo(ID)
	if err != nil {
		return true, service.domainAdvice.TranslateError(err)
	}
	for _, category := range toDo.Categories {
		if util.Contains[uuid.UUID](categories, category.ID) {
			return true, nil
		}
	}
	return false, nil
}

func (service *toDoWriteService) AddToDo(event model.CreateToDoEvent) error.DomainError {
	categories := make([]model.Category, 0)
	for _, categoryId := range event.Categories {
		categories = append(categories, model.Category{
			ID: categoryId,
		})
	}
	toDo := model.ToDo{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		CreatedAt:   &event.CreatedAt,
		Categories:  categories,
	}
	err := service.toDoDatabaseService.CreateToDo(toDo)
	return service.domainAdvice.TranslateError(err)
}

func (service *toDoWriteService) UpdateToDo(event model.UpdateToDoEvent) error.DomainError {
	updatedAt := time.Now()
	toDo := model.ToDo{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		UpdatedAt:   &updatedAt,
	}
	err := service.toDoDatabaseService.UpdateToDo(toDo)
	return service.domainAdvice.TranslateError(err)
}

func (service *toDoWriteService) DeleteToDo(event model.DeleteToDoEvent) error.DomainError {
	err := service.toDoDatabaseService.DeleteToDo(event.ID)
	return service.domainAdvice.TranslateError(err)
}

func (service *toDoWriteService) AddCategoriesIntoToDo(event model.AddCategoriesFromToDoEvent) error.DomainError {
	err := service.toDoDatabaseService.AddCategoriesIntoToDo(event)
	return service.domainAdvice.TranslateError(err)
}

func (service *toDoWriteService) RemoveCategoriesFromToDo(event model.RemoveCategoriesFromToDoEvent) error.DomainError {
	err := service.toDoDatabaseService.RemoveCategoriesFromToDo(event)
	return service.domainAdvice.TranslateError(err)
}

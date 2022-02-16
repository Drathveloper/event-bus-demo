package service

import (
	"event-bus-demo/application/dto"
	"event-bus-demo/domain/error"
	"event-bus-demo/domain/mapper"
	"event-bus-demo/domain/model"
	"event-bus-demo/infrastructure/database/service"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CategoryReadService interface {
	GetCategories() (dto.GetCategoriesResponse, error.DomainError)
	GetCategoryById(event model.GetCategoryByIDEvent) (dto.GetCategoryResponse, error.DomainError)
	GetCategoriesByIds([]uuid.UUID) (dto.GetCategoriesResponse, error.DomainError)
}

type CategoryWriteService interface {
	AddUser(event model.CreateCategoryEvent) error.DomainError
	UpdateUser(event model.UpdateCategoryNameEvent) error.DomainError
	DeleteUser(event model.DeleteCategoryEvent) error.DomainError
}

type categoryReadService struct {
	logger                  *zap.Logger
	categoryDatabaseService service.CategoryDatabaseService
	domainAdvice            error.DomainAdvice
}

type categoryWriteService struct {
	logger                  *zap.Logger
	categoryDatabaseService service.CategoryDatabaseService
	domainAdvice            error.DomainAdvice
}

func NewCategoryReadService(categoryDatabaseService service.CategoryDatabaseService, domainAdvice error.DomainAdvice, logger *zap.Logger) CategoryReadService {
	return &categoryReadService{
		logger:                  logger,
		categoryDatabaseService: categoryDatabaseService,
		domainAdvice:            domainAdvice,
	}
}

func NewCategoryWriteService(categoryDatabaseService service.CategoryDatabaseService, domainAdvice error.DomainAdvice, logger *zap.Logger) CategoryWriteService {
	return &categoryWriteService{
		logger:                  logger,
		categoryDatabaseService: categoryDatabaseService,
		domainAdvice:            domainAdvice,
	}
}

func (service *categoryReadService) GetCategories() (dto.GetCategoriesResponse, error.DomainError) {
	categories, err := service.categoryDatabaseService.GetAllCategories()
	return mapper.NewGetCategoriesResponseFromDomainModelList(categories), service.domainAdvice.TranslateError(err)
}

func (service *categoryReadService) GetCategoryById(event model.GetCategoryByIDEvent) (dto.GetCategoryResponse, error.DomainError) {
	category, err := service.categoryDatabaseService.GetCategory(event.ID)
	return mapper.NewGetCategoryResponseFromDomainModel(category), service.domainAdvice.TranslateError(err)
}

func (service *categoryReadService) GetCategoriesByIds(categoriesID []uuid.UUID) (dto.GetCategoriesResponse, error.DomainError) {
	response := make([]dto.GetCategoryResponse, 0)
	for _, categoryID := range categoriesID {
		foundCategory, err := service.GetCategoryById(model.GetCategoryByIDEvent{ID: categoryID})
		if err != nil {
			return dto.GetCategoriesResponse{}, err
		}
		response = append(response, foundCategory)
	}
	return dto.GetCategoriesResponse{
		Categories: response,
	}, nil
}

func (service *categoryWriteService) AddUser(event model.CreateCategoryEvent) error.DomainError {
	category := model.Category{
		ID:   event.ID,
		Name: event.Name,
	}
	err := service.categoryDatabaseService.CreateCategory(category)
	return service.domainAdvice.TranslateError(err)
}

func (service *categoryWriteService) UpdateUser(event model.UpdateCategoryNameEvent) error.DomainError {
	category := model.Category{
		ID:   event.ID,
		Name: event.Name,
	}
	err := service.categoryDatabaseService.UpdateCategory(category)
	return service.domainAdvice.TranslateError(err)
}

func (service *categoryWriteService) DeleteUser(event model.DeleteCategoryEvent) error.DomainError {
	err := service.categoryDatabaseService.DeleteCategory(event.ID)
	return service.domainAdvice.TranslateError(err)
}

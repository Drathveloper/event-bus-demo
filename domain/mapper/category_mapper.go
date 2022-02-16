package mapper

import (
	"event-bus-demo/application/dto"
	domainModel "event-bus-demo/domain/model"
)

func NewGetCategoryResponseFromDomainModel(category domainModel.Category) dto.GetCategoryResponse {
	return dto.GetCategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func NewGetCategoriesResponseFromDomainModelList(categories []domainModel.Category) dto.GetCategoriesResponse {
	dtoList := make([]dto.GetCategoryResponse, 0)
	for _, category := range categories {
		dtoList = append(dtoList, NewGetCategoryResponseFromDomainModel(category))
	}
	return dto.GetCategoriesResponse{
		Categories: dtoList,
	}
}

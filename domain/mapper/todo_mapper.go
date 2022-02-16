package mapper

import (
	"event-bus-demo/application/dto"
	"event-bus-demo/domain/model"
)

func NewGetToDoResponseFromDomainModel(model model.ToDo) dto.GetToDoResponse {
	return dto.GetToDoResponse{
		ID:          model.ID,
		Title:       model.Title,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		Categories:  NewGetCategoriesResponseFromDomainModelList(model.Categories).Categories,
	}
}

func NewGetAllToDoResponseFromDomainModel(modelList []model.ToDo) dto.GetAllToDoResponse {
	responseList := make([]dto.GetToDoResponse, 0)
	for _, toDo := range modelList {
		responseList = append(responseList, NewGetToDoResponseFromDomainModel(toDo))
	}
	return dto.GetAllToDoResponse{
		ToDos: responseList,
	}
}

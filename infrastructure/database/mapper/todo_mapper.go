package mapper

import (
	domainModel "event-bus-demo/domain/model"
	dbModel "event-bus-demo/infrastructure/database/model"
	"event-bus-demo/infrastructure/database/sqlc"
	"time"
)

func NewToDoEntityFromToDoModel(model domainModel.ToDo) dbModel.ToDoEntity {
	return dbModel.ToDoEntity{
		ID:          model.ID,
		Title:       model.Title,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		Categories:  NewCategoryEntityListFromCategoryListModel(model.Categories),
	}
}

func NewToDoEntityFromSQLModel(sqlModel sqlc.Todo) dbModel.ToDoEntity {
	var updatedAt *time.Time
	updatedAt = nil
	if sqlModel.UpdatedAt.Valid {
		updatedAt = &sqlModel.UpdatedAt.Time
	}
	createdAt := sqlModel.CreatedAt.Local()
	return dbModel.ToDoEntity{
		ID:          sqlModel.ID,
		Title:       sqlModel.Title,
		Description: sqlModel.Description,
		CreatedAt:   &createdAt,
		UpdatedAt:   updatedAt,
	}
}

func NewToDoEntityListFromSQLModelList(sqlModelList []sqlc.Todo) []dbModel.ToDoEntity {
	entities := make([]dbModel.ToDoEntity, 0)
	for _, sqlModel := range sqlModelList {
		entity := NewToDoEntityFromSQLModel(sqlModel)
		entities = append(entities, entity)
	}
	return entities
}

func NewToDoFromEntity(entity dbModel.ToDoEntity) domainModel.ToDo {
	return domainModel.ToDo{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		Categories:  NewCategoriesListFromEntity(entity.Categories),
	}
}

func NewToDoListFromEntityList(entities []dbModel.ToDoEntity) []domainModel.ToDo {
	models := make([]domainModel.ToDo, 0)
	for _, entity := range entities {
		models = append(models, NewToDoFromEntity(entity))
	}
	return models
}

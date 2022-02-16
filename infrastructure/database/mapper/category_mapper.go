package mapper

import (
	domainModel "event-bus-demo/domain/model"
	dbModel "event-bus-demo/infrastructure/database/model"
	"event-bus-demo/infrastructure/database/sqlc"
)

func NewCategoryEntityFromCategoryModel(model domainModel.Category) dbModel.CategoryEntity {
	return dbModel.CategoryEntity{
		ID:   model.ID,
		Name: model.Name,
	}
}

func NewCategoryEntityListFromCategoryListModel(model []domainModel.Category) []dbModel.CategoryEntity {
	entities := make([]dbModel.CategoryEntity, 0)
	for _, e := range model {
		entities = append(entities, NewCategoryEntityFromCategoryModel(e))
	}
	return entities
}

func NewCategoryEntityFromSQLModel(sqlModel sqlc.Category) dbModel.CategoryEntity {
	return dbModel.CategoryEntity{
		ID:   sqlModel.ID,
		Name: sqlModel.Name,
	}
}

func NewCategoryEntityListFromSQLModelList(sqlModelList []sqlc.Category) []dbModel.CategoryEntity {
	entities := make([]dbModel.CategoryEntity, 0)
	for _, sqlModel := range sqlModelList {
		entity := NewCategoryEntityFromSQLModel(sqlModel)
		entities = append(entities, entity)
	}
	return entities
}

func NewCategoryFromEntity(entity dbModel.CategoryEntity) domainModel.Category {
	return domainModel.Category{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func NewCategoriesListFromEntity(entities []dbModel.CategoryEntity) []domainModel.Category {
	models := make([]domainModel.Category, 0)
	for _, entity := range entities {
		models = append(models, NewCategoryFromEntity(entity))
	}
	return models
}

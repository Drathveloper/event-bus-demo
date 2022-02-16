package mapper

import (
	domainModel "event-bus-demo/domain/model"
	dbModel "event-bus-demo/infrastructure/database/model"
	"event-bus-demo/infrastructure/database/sqlc"
)

func NewUserEntityFromUserModel(user domainModel.User) dbModel.UserEntity {
	return dbModel.UserEntity{
		ID:       user.ID,
		Username: user.Username,
		Password: user.GetPassword(),
		Role:     user.Role,
	}
}

func NewUserEntityFromSQLModel(sqlModel sqlc.User) dbModel.UserEntity {
	return dbModel.UserEntity{
		ID:       sqlModel.ID,
		Username: sqlModel.Username,
		Password: sqlModel.Password,
	}
}

func NewUserFromEntity(entity dbModel.UserEntity) domainModel.User {
	user := domainModel.User{
		ID:       entity.ID,
		Username: entity.Username,
	}
	user.SetHashedPassword(entity.Password)
	return user
}

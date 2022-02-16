package event

import (
	"event-bus-demo/domain/model"
	"event-bus-demo/domain/service"
	"event-bus-demo/infrastructure/event_sourcing"
	"fmt"
	"go.uber.org/zap"
)

type userEventHandler struct {
	userService service.UserWriteService
	logger      *zap.Logger
}

func NewUserEventHandler(userService service.UserWriteService, logger *zap.Logger) event_sourcing.EventHandler {
	eventHandler := &userEventHandler{
		userService: userService,
		logger:      logger,
	}
	return eventHandler
}

func (handler *userEventHandler) Handle(event event_sourcing.Event) event_sourcing.EventResult {
	var err error
	switch event.(type) {
	case model.CreateUserEvent:
		err = handler.userService.AddUser(event.(model.CreateUserEvent))
	case model.UpdateUserPasswordEvent:
		err = handler.userService.UpdateUserPassword(event.(model.UpdateUserPasswordEvent))
	case model.DeleteUserEvent:
		err = handler.userService.DeleteUser(event.(model.DeleteUserEvent))
	default:
		err = fmt.Errorf("unknown event")
	}
	return HandleError(handler.logger, event, err)
}

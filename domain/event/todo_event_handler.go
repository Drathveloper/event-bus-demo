package event

import (
	"event-bus-demo/domain/model"
	"event-bus-demo/domain/service"
	"event-bus-demo/infrastructure/event_sourcing"
	"fmt"
	"go.uber.org/zap"
)

type toDoEventHandler struct {
	toDoService service.ToDoWriteService
	logger      *zap.Logger
}

func NewToDoEventHandler(toDoService service.ToDoWriteService, logger *zap.Logger) event_sourcing.EventHandler {
	eventHandler := &toDoEventHandler{
		toDoService: toDoService,
		logger:      logger,
	}
	return eventHandler
}

func (handler *toDoEventHandler) Handle(event event_sourcing.Event) event_sourcing.EventResult {
	var err error
	switch event.(type) {
	case model.CreateToDoEvent:
		err = handler.toDoService.AddToDo(event.(model.CreateToDoEvent))
	case model.UpdateToDoEvent:
		err = handler.toDoService.UpdateToDo(event.(model.UpdateToDoEvent))
	case model.DeleteToDoEvent:
		err = handler.toDoService.DeleteToDo(event.(model.DeleteToDoEvent))
	case model.AddCategoriesFromToDoEvent:
		err = handler.toDoService.AddCategoriesIntoToDo(event.(model.AddCategoriesFromToDoEvent))
	case model.RemoveCategoriesFromToDoEvent:
		err = handler.toDoService.RemoveCategoriesFromToDo(event.(model.RemoveCategoriesFromToDoEvent))
	default:
		err = fmt.Errorf("unknown event")
	}
	return HandleError(handler.logger, event, err)
}

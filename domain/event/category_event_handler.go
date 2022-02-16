package event

import (
	"event-bus-demo/domain/model"
	"event-bus-demo/domain/service"
	"event-bus-demo/infrastructure/event_sourcing"
	"fmt"
	"go.uber.org/zap"
)

type categoryEventHandler struct {
	categoryService service.CategoryWriteService
	logger          *zap.Logger
}

func NewCategoryEventHandler(categoryService service.CategoryWriteService, logger *zap.Logger) event_sourcing.EventHandler {
	eventHandler := &categoryEventHandler{
		categoryService: categoryService,
		logger:          logger,
	}
	return eventHandler
}

func (handler *categoryEventHandler) Handle(event event_sourcing.Event) event_sourcing.EventResult {
	var err error
	switch event.(type) {
	case model.CreateCategoryEvent:
		err = handler.categoryService.AddUser(event.(model.CreateCategoryEvent))
	case model.UpdateCategoryNameEvent:
		err = handler.categoryService.UpdateUser(event.(model.UpdateCategoryNameEvent))
	case model.DeleteCategoryEvent:
		err = handler.categoryService.DeleteUser(event.(model.DeleteCategoryEvent))
	default:
		err = fmt.Errorf("unknown event")
	}
	return HandleError(handler.logger, event, err)
}

package main

import (
	"event-bus-demo/application/controller"
	applicationError "event-bus-demo/application/error"
	domainError "event-bus-demo/domain/error"
	"event-bus-demo/domain/event"
	"event-bus-demo/domain/model"
	"event-bus-demo/domain/service"
	"event-bus-demo/infrastructure/configuration"
	"event-bus-demo/infrastructure/constants"
	"event-bus-demo/infrastructure/database/repository"
	dbService "event-bus-demo/infrastructure/database/service"
	"event-bus-demo/infrastructure/event_sourcing"
	"go.uber.org/zap"
)

type RequiredDependencies struct {
	EventBus            event_sourcing.EventBus
	RequiredControllers RequiredControllers
}

type RequiredControllers struct {
	ToDoController     controller.ToDoController
	CategoryController controller.CategoryController
	UserController     controller.UserController
}

func initializeDependencies(config configuration.ApplicationConfiguration) (RequiredDependencies, error) {
	env, err := constants.NewEnvironment(*config.Gin.Environment)
	if err != nil {
		return RequiredDependencies{}, err
	}
	var logger *zap.Logger
	if env.IsProductiveEnvironment() {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}
	connectionPool, err := configuration.BuildDatabase(*config.Rdbms)
	if err != nil {
		return RequiredDependencies{}, err
	}

	// Event bus
	eventChannel := event_sourcing.NewBufferedEventChannel(*config.Event.ChannelBufferSize)
	eventBus := event_sourcing.NewEventBus(eventChannel, *config.Event.MaxWorkers, logger)

	// Repository
	transactionalRepository := repository.NewTransactionalRepository(logger, connectionPool)
	toDoRepository := repository.NewToDoRepository(logger, connectionPool)
	categoryRepository := repository.NewCategoryRepository(logger, connectionPool)
	userRepository := repository.NewUserRepository(logger, connectionPool)

	// Infrastructure service
	toDoDatabaseService := dbService.NewToDoDatabaseService(transactionalRepository, toDoRepository)
	categoryDatabaseService := dbService.NewCategoryDatabaseService(transactionalRepository, categoryRepository)
	userDatabaseService := dbService.NewUserDatabaseService(transactionalRepository, userRepository)

	// Domain service
	domainAdvice := domainError.NewDomainAdvice()
	toDoReadService := service.NewToDoReadService(toDoDatabaseService, domainAdvice, logger)
	toDoWriteService := service.NewToDoWriteService(toDoDatabaseService, domainAdvice, logger)
	categoryReadService := service.NewCategoryReadService(categoryDatabaseService, domainAdvice, logger)
	categoryWriteService := service.NewCategoryWriteService(categoryDatabaseService, domainAdvice, logger)
	userReadService := service.NewUserReadService(userDatabaseService, domainAdvice, logger)
	userWriteService := service.NewUserWriteService(userDatabaseService, domainAdvice, logger)

	// Event handler
	toDoEventHandler := event.NewToDoEventHandler(toDoWriteService, logger)
	categoryEventHandler := event.NewCategoryEventHandler(categoryWriteService, logger)
	userEventHandler := event.NewUserEventHandler(userWriteService, logger)

	// Event Subscriber
	loggerSubscriber := event.NewEventLoggerSubscriber(logger)

	// Controller
	controllerAdvice := applicationError.NewControllerAdvice()
	toDoController := controller.NewTodoController(eventBus, toDoReadService, categoryReadService, controllerAdvice)
	categoryController := controller.NewCategoryController(eventBus, categoryReadService, controllerAdvice)
	userController := controller.NewUserController(eventBus, userReadService, controllerAdvice)

	// Register handlers on eventBus
	eventBus.RegisterHandler(model.ToDoEventTopic, toDoEventHandler)
	eventBus.RegisterHandler(model.CategoryEventTopic, categoryEventHandler)
	eventBus.RegisterHandler(model.UserEventTopic, userEventHandler)

	// Register subscribers on eventBus
	eventBus.RegisterSubscriber(model.ToDoEventTopic, loggerSubscriber)
	eventBus.RegisterSubscriber(model.CategoryEventTopic, loggerSubscriber)
	eventBus.RegisterSubscriber(model.UserEventTopic, loggerSubscriber)

	return RequiredDependencies{
		EventBus: eventBus,
		RequiredControllers: RequiredControllers{
			ToDoController:     toDoController,
			CategoryController: categoryController,
			UserController:     userController,
		},
	}, nil
}

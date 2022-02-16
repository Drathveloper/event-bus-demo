package main

import (
	"event-bus-demo/infrastructure/util"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(profiles []string, controllers RequiredControllers) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

	v1Group := router.Group("/v1")
	{
		toDoGroup := v1Group.Group("/todo")
		{
			toDoGroup.GET("", controllers.ToDoController.GetToDoList)
			toDoGroup.POST("", controllers.ToDoController.SaveToDo)
			toDoGroup.GET("/:id", controllers.ToDoController.GetToDoById)
			toDoGroup.PUT("/:id", controllers.ToDoController.UpdateToDo)
			toDoGroup.DELETE("/:id", controllers.ToDoController.DeleteToDo)
			toDoGroup.PATCH("/:id/categories", controllers.ToDoController.AddCategoriesIntoToDo)
			toDoGroup.DELETE("/:id/categories", controllers.ToDoController.RemoveCategoriesFromToDo)
		}
		categoryGroup := v1Group.Group("/category")
		{
			categoryGroup.GET("", controllers.CategoryController.GetCategories)
			categoryGroup.POST("", controllers.CategoryController.SaveCategory)
			categoryGroup.GET("/:id", controllers.CategoryController.GetCategoryById)
			categoryGroup.PUT("/:id", controllers.CategoryController.UpdateCategory)
			categoryGroup.DELETE("/:id", controllers.CategoryController.DeleteCategory)
		}
		if util.Contains[string](profiles, "with_users") {
			userGroup := v1Group.Group("/user")
			{
				userGroup.POST("", controllers.UserController.CreateUser)
				userGroup.GET("/:id", controllers.UserController.GetUserByID)
				userGroup.PATCH("/:id", controllers.UserController.UpdateUserPassword)
				userGroup.DELETE("/:id", controllers.UserController.DeleteUser)
			}
		}
	}
	return router
}

package main

import (
	"choreboard/models"
	"github.com/gin-gonic/gin"
)

func main() {

	// Creates a router without any middleware by default
	router := gin.New()

	//Middleware setup
	router.Use(gin.Recovery())

	models.InitDb()

	choreController := registerChoreController()
	userController := registerUserController()

	v1 := router.Group("/v1")
	{
		v1.POST("/users", userController.NewUser)
		v1.POST("/user/login", userController.Login)
		v1.GET("/users/:user_id", userController.GetUser)

		v1.GET("/users/:user_id/chores", choreController.GetUserCreatedChores)
		v1.GET("/users/:user_id/assigned_chores", choreController.GetUserAssignedChores)

		v1.POST("/users/chores", choreController.NewChore)
		v1.PATCH("/users/chores/:chore_id", choreController.UpdateChore)

		v1.PUT("/chores/:chore_id/users/:user_id", choreController.AssignChore)
	}

	router.Run(":3000")
}

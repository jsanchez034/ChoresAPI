package main

import (
	"choreboard/controllers"
	"choreboard/models"
	"choreboard/models/rethinkdb"
)

func registerChoreController() *controllers.ChoreController {
	choreServices := rethinkdb.NewRethinkChoreService(models.Session)
	return controllers.NewChoreController(choreServices)
}

func registerUserController() *controllers.UserController {
	userService := rethinkdb.NewRethinkUserService(models.Session)
	return controllers.NewUserController(userService)
}

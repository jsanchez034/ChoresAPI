package controllers

import (
	"choreboard/controllers/requests"
	"choreboard/controllers/responses"
	"choreboard/models"
	"choreboard/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	dbservice models.UserService
}

func NewUserController(userservice models.UserService) *UserController {

	return &UserController{
		dbservice: userservice,
	}
}

func (c *UserController) NewUser(ctx *gin.Context) {
	//Decode respose body into Chore struct
	var user models.User

	err := ctx.ParseBody(&user)

	if err != nil {
		ctx.JSON(400, &utils.Error{400, err.Error()})
		return
	}

	//Insert user into DB
	err = c.dbservice.CreateUser(&user)

	if err != nil {
		ctx.JSON(err.(*utils.Error).Status, err)
		return
	}

	//Render chore as Json
	ctx.JSON(200, user)

}

func (c *UserController) GetUser(ctx *gin.Context) {

	userid := ctx.Params.ByName("user_id")

	if userid == "" {
		ctx.JSON(400, &utils.Error{404, "User id required"})
		return
	}

	//Loggin a user
	user, err := c.dbservice.GetUser(userid)

	if err != nil {
		ctx.JSON(err.(*utils.Error).Status, err)
		return
	}

	//created nested resource links
	userLinks := make(map[string]string)
	userLinks["chores"] = "/v1/users/" + userid + "/chores"
	userLinks["assigned_chores"] = "/v1/users/" + userid + "/assigned_chores"
	userResponse := responses.UserResponse{
		user,
		userLinks,
	}

	//Render chore as Json
	ctx.JSON(200, gin.H{"user": userResponse})

}

func (c *UserController) Login(ctx *gin.Context) {
	var loginRequest requests.UserLoginRequest

	err := ctx.ParseBody(&loginRequest)

	if err != nil {
		ctx.JSON(400, &utils.Error{400, err.Error()})
		return
	}

	//Loggin a user
	user, err := c.dbservice.LoginUser(loginRequest.Email, loginRequest.Password)

	if err != nil {
		ctx.JSON(err.(*utils.Error).Status, err)
		return
	}

	//Render chore as Json
	ctx.JSON(200, user)

}

package controllers

import (
	"choreboard/models"
	"choreboard/utils"
	"github.com/gin-gonic/gin"
)

type ChoreController struct {
	dbservice models.ChoreService
}

func NewChoreController(choreservice models.ChoreService) *ChoreController {

	return &ChoreController{
		dbservice: choreservice,
	}
}

func (c *ChoreController) GetUserCreatedChores(ctx *gin.Context) {

	//Get chores from DB
	userid := ctx.Params.ByName("user_id")
	chores, err := c.dbservice.GetUserCreatedChores(userid)

	if err != nil {
		ctx.JSON(500, &utils.Error{500, err.Error()})
		return
	}

	//Render chores as Json
	ctx.JSON(200, gin.H{"chores": chores})
}

func (c *ChoreController) GetUserAssignedChores(ctx *gin.Context) {
	//Get chores from DB
	userid := ctx.Params.ByName("user_id")
	chores, err := c.dbservice.GetUserAssignedChores(userid)

	if err != nil {
		ctx.JSON(500, &utils.Error{500, err.Error()})
		return
	}

	//Render chores as Json
	ctx.JSON(200, gin.H{"chores": chores})
}

func (c *ChoreController) NewChore(ctx *gin.Context) {
	//Decode respose body into Chore struct
	var chore models.Chore

	err := ctx.ParseBody(&chore)

	if err != nil {
		ctx.JSON(400, &utils.Error{400, err.Error()})
		return
	}

	//Insert chores into DB
	err = c.dbservice.InsertChore(&chore)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	//Render chore as Json
	ctx.JSON(200, chore)
}

func (c *ChoreController) UpdateChore(ctx *gin.Context) {
	//Decode respose body into Chore struct
	var chore models.Chore

	err := ctx.ParseBody(&chore)

	if err != nil {
		ctx.JSON(400, &utils.Error{400, err.Error()})
		return
	}

	//Set Chore Id to update from uri param
	chore.Id = ctx.Params.ByName("chore_id")

	//Make sure created by, created date, last modified, assigned user
	//are never updated by user
	chore.CreatedBy = nil
	chore.Created = nil
	chore.LastModified = nil
	chore.AssignedUserId = nil

	//Update chore in DB
	chore, err = c.dbservice.UpdateChore(&chore)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	//Render chore as Json
	ctx.JSON(200, chore)
}

func (c *ChoreController) AssignChore(ctx *gin.Context) {
	//Decode respose body into Chore struct
	var chore models.Chore

	//Set User that is to be assigned Chore Id to update from uri param
	chore.Id = ctx.Params.ByName("chore_id")
	chore.AssignedUserId = utils.String(ctx.Params.ByName("user_id"))

	//Update chore in DB
	chore, err := c.dbservice.AssignChore(&chore)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	//Render chore as Json
	ctx.JSON(200, chore)
}

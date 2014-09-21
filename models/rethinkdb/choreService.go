package rethinkdb

import (
	. "choreboard/models"
	"choreboard/utils"
	r "github.com/dancannon/gorethink"
	re "github.com/dancannon/gorethink/encoding"
	"time"
)

type RethinkChoreService struct {
	session *r.Session
}

func NewRethinkChoreService(session *r.Session) *RethinkChoreService {
	return &RethinkChoreService{
		session: session,
	}
}

func (choreService *RethinkChoreService) GetUserCreatedChores(userid string) ([]Chore, error) {
	chores := []Chore{}

	// Fetch all the items from the database
	res, err := r.Table("chores").Filter(r.Row.Field("created_by").Eq(userid)).OrderBy(r.Asc("created")).Run(choreService.session)
	if err != nil {
		return nil, &utils.Error{500, "Could not retrieve chores"}
	}

	// Scan each row into a Chore instance and then add this to the list
	err = res.All(&chores)

	if err != nil {
		return nil, &utils.Error{500, "Could not retrieve chores"}
	}

	return chores, nil

}

func (choreService *RethinkChoreService) GetUserAssignedChores(userid string) ([]Chore, error) {
	chores := []Chore{}

	// Fetch all the items from the database
	res, err := r.Table("chores").Filter(r.Row.Field("assigned_user_id").Eq(userid)).OrderBy(r.Asc("last_modified")).Run(choreService.session)
	if err != nil {
		return nil, &utils.Error{500, "Could not retrieve chores"}
	}

	// Scan each row into a Chore instance and then add this to the list
	err = res.All(&chores)

	if err != nil {
		return nil, &utils.Error{500, "Could not retrieve chores"}
	}

	return chores, nil

}

func (choreService *RethinkChoreService) InsertChore(chore *Chore) error {

	dt := time.Now()
	chore.Created = &dt

	// Insert the chore into the database
	res, err := r.Table("chores").Insert(chore).RunWrite(choreService.session)
	if err != nil {
		return &utils.Error{500, "Could not insert chore"}
	}

	chore.Id = res.GeneratedKeys[0]

	return nil
}

func (choreService *RethinkChoreService) UpdateChore(chore *Chore) (Chore, error) {
	dt := time.Now()
	chore.LastModified = &dt

	// Update the chore in the database
	res, err := r.Table("chores").Get(chore.Id).Update(chore, r.UpdateOpts{ReturnVals: true}).RunWrite(choreService.session)
	if err != nil {
		return Chore{}, &utils.Error{500, "Could not update chore"}
	}

	var newValueChore = new(Chore)
	err = re.Decode(newValueChore, res.NewValue)

	if err != nil {
		return Chore{}, &utils.Error{500, err.Error()}
	}

	return *newValueChore, nil
}

func (choreService *RethinkChoreService) AssignChore(chore *Chore) (Chore, error) {

	return choreService.UpdateChore(&Chore{
		Id:             chore.Id,
		AssignedUserId: chore.AssignedUserId,
	})
}

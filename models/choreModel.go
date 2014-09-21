package models

import "time"

//interface for service object that interacts with model and db
type ChoreService interface {
	GetUserCreatedChores(userid string) ([]Chore, error)
	InsertChore(chore *Chore) error
	UpdateChore(chore *Chore) (Chore, error)
	AssignChore(chore *Chore) (Chore, error)
	GetUserAssignedChores(userid string) ([]Chore, error)
	//DeleteChore(choreid string) error
}

//model
type Chore struct {
	Id             string     `json:"id" gorethink:"id,omitempty"`
	Name           *string    `json:"name" gorethink:"name,omitempty""`
	Description    *string    `json:"description" gorethink:"description,omitempty""`
	CreatedBy      *string    `json:"created_by" gorethink:"created_by,omitempty"`
	Created        *time.Time `json:"created" gorethink:"created,omitempty"`
	LastModified   *time.Time `json:"last_modified" gorethink:"last_modified,omitempty"`
	AssignedUserId *string    `json:"assigned_user_id" gorethink:"assigned_user_id,omitempty"`
}

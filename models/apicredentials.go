package models

import "time"

//interface for service object that interacts with model and db
type ApiCredentialsService interface {
	GenerateUserCredentials(userid string) (ApiCredentials, error)
	RevokeApiUserCredentials(userid string) error
}

//model
type ApiCredentials struct {
	UserId  string     `json:"user_id" gorethink:"user_id,omitempty"`
	Revoked bool       `json:"revoked" gorethink:"revoked,omitempty"`
	Created *time.Time `json:"created" gorethink:"created,omitempty"`
}

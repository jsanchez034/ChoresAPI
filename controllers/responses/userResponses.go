package responses

import "choreboard/models"

type Links map[string]string

type UserResponse struct {
	models.User
	Links `json:"links"`
}

package rethinkdb

import (
	. "choreboard/models"
	"choreboard/utils"
	r "github.com/dancannon/gorethink"
	"time"
)

type RethinkUserService struct {
	session *r.Session
}

func NewRethinkUserService(session *r.Session) *RethinkUserService {
	return &RethinkUserService{
		session: session,
	}
}

func (userService *RethinkUserService) GetUser(userid string) (User, error) {
	var user User

	//Get user by id
	res, err := r.Table("user").Get(userid).Run(userService.session)
	if err != nil || res.IsNil() {
		return User{}, &utils.Error{404, "Could not find user"}
	}

	//Hydrate user object
	err = res.One(&user)

	if err != nil {
		return User{}, &utils.Error{404, "Could not find user"}
	}

	//Make sure to never return password
	user.Password = ""

	return user, nil
}

func (userService *RethinkUserService) UpdateUser(user *User) error {
	return nil
}

func (userService *RethinkUserService) CreateUser(user *User) error {
	var count int
	//Verify user does not already exist
	res, err := r.Table("user").Filter(r.Row.Field("email").Eq(user.Email)).Count().Run(userService.session)

	if err != nil {
		return &utils.Error{500, "Could not create user"}
	}

	err = res.One(&count)

	if count != 0 {
		return &utils.Error{409, "User already exists"}
	}

	dt := time.Now()
	user.Created = &dt

	err = user.SetEncryptPassword(user.Password)

	if err != nil {
		return &utils.Error{500, "Could not create user"}
	}

	//Insert the user into the database
	writeresponse, err := r.Table("user").Insert(user).RunWrite(userService.session)
	if err != nil {
		return &utils.Error{500, "Could not create user"}
	}

	user.Password = ""
	user.Id = writeresponse.GeneratedKeys[0]

	return nil
}

func (userService *RethinkUserService) LoginUser(email string, password string) (*User, error) {
	var user User

	//Get user by email
	res, err := r.Table("user").Filter(r.Row.Field("email").Eq(email)).Run(userService.session)
	if err != nil {
		return nil, &utils.Error{500, err.Error()}
	}

	//Hydrate user object
	err = res.One(&user)

	if err != nil {
		return nil, &utils.Error{401, "Unauthorized"}
	}

	//Verify password matches
	if pwd, err := user.GetDecryptPassword(); err != nil || password != pwd {
		return nil, &utils.Error{401, "Unauthorized"}
	}

	return &user, nil
}

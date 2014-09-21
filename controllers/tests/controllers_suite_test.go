package controllers_test

import (
	. "choreboard/models"
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"

	"testing"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var now time.Time = time.Now()

//depency injection using interface
//for services that get data from db
type fakeChoreServices struct {
}

func (choreService *fakeChoreServices) GetAllCreatedChores(userid string) ([]Chore, error) {

	return []Chore{
		Chore{"123", "test1", "Description1", "guid-user-id", &now},
		Chore{"1234", "test2", "Description2", "guid-user-id", &now},
		Chore{"12345", "test3", "Description3", "guid-user-id", &now},
		Chore{"123456", "test4", "Description4", "guid-user-id", &now},
	}, nil
}

func (choreService *fakeChoreServices) InsertChore(chore *Chore) error {
	return nil
}

type fakeChoreErrorServices struct {
}

func (choreService *fakeChoreErrorServices) GetAllCreatedChores(userid string) ([]Chore, error) {

	return nil, errors.New("DB Error")
}

func (choreService *fakeChoreErrorServices) InsertChore(chore *Chore) error {
	return errors.New("DB Error")
}

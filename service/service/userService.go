package service

import (
	"agenda_api/cli/entity"
)

//UserAtomicService
type UserAtomicService struct{}

//UserService
var UserService = UserAtomicService{}

//Add
func (*UserAtomicService) Add(u *entity.User) error{

}

//Delete
func 
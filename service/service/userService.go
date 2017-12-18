package service

import (
	"agenda_api/cli/server/entity/User"
)

//UserAtomicService
type UserAtomicService struct{}

//UserService
var UserService = UserAtomicService{}

//
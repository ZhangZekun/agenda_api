package service

import (
	"agenda_api/cli/entity"
)

//SessionAtomicService
type SessionAtomicService struct{}

//SessionService
var SessionService = SessionAtomicService{}

//Save
func (*SessionAtomicService) Save(s *entity.Session) error{
	_, err := 
}
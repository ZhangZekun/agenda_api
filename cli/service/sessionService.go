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



	return nil
}

//Find
func (*SessionAtomicService) FindSession(s *entity.Session) bool{

	return false
}

//Delete
func (*SessionAtomicService) DeleteSession(s *entity.Session) error{

	return nil
}
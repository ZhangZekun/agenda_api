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
	_, err := Enginea.Insert(s)

	CheckErr(err)

	return nil
}

//Find
func (*SessionAtomicService) FindSession(s *entity.Session) bool{
	has, err := Enginea.Get(s)
	return has
}

//Delete
func (*SessionAtomicService) DeleteSession(s *entity.Session) error{
	_, err := Enginea.Delete(s)

	CheckErr(err)

	return nil
}
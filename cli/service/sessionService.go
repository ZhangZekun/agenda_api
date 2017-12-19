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

	if (SessionService.FindSession(s) == false){
		_, err := Enginea.Insert(s)
		CheckErr(err)
	}

	return nil
}

//Find
func (*SessionAtomicService) FindSession(s *entity.Session) bool{
	has, err := Enginea.Get(s)
	CheckErr(err)
	return has
}

//Delete
func (*SessionAtomicService) DeleteSession(s *entity.Session) error{
	_, err := Enginea.Where("current_user = ?", s.Currentuser).Delete(s)

	CheckErr(err)

	return nil
}
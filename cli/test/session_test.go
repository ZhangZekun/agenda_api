package test

import (
	"agenda_api/cli/entity"
	"agenda_api/cli/service"
	"testing"
	"fmt"
)

func TestSessionService(t *testing.T) {
	curentUser := "test_session1"
	session := entity.NewSession(entity.Session{Currentuser:curentUser})

	fmt.Println("saveSession_test")
	service.SessionService.Save(session)
	
	fmt.Println("querySession_test")
	query_session := service.SessionService.FindSession(session)
	fmt.Print("isUserLogin:")
	fmt.Print(query_session)

	fmt.Println("deleteSession_test")
	service.SessionService.DeleteSession(session)
	query_session = service.SessionService.FindSession(session)
	fmt.Println("isUserLogin:")
	fmt.Print(query_session)
}
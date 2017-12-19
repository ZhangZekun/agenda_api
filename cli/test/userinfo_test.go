package test

import(
	"testing"
	"agenda_api/cli/service"
	"agenda_api/cli/entity"
)

func TestUserinfoService(t *testing.T){

	user := entity.NewUserInfo(entity.User{Username:"test_user", Password:"test_password", Email:"test_email@163.com", Phone:"12345678901"})

	if err1 := service.UserInfoService.InsertUser(user); err1 != nil {
		t.Fatal(err1)
	}

	if test_user, err2 := service.UserInfoService.FindByUsername(user.Username); err2 != nil{
		t.Fatal(err2)
	}

	session := entity.NewSession(entity.Session{LoginId:"", CurrentUser:"test_user"})
	
	if err3 := service.UserInfoService.LoginInfoInsert(session); err3 != nil{
		t.Fatal(err3)
	}

	if service.UserInfoService.UserHasLogin(session.LoginId) == true{
		login_list, err4 := service.UserInfoService.GetAllUsersInfo(session)
	}
	
	if service.UserInfoService.UserHasLogin(session.LoginId) == true{
		if err5 := service.UserInfoService.LoginInfoDelete(session); err5 != nil{
			t.Fatal(err5)
		}
	
	}
}
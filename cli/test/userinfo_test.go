package test

import(
	"testing"
	"agenda_api/cli/service"
	"agenda_api/cli/entity"
	"database/sql"
)

func TestUserinfoService(t *testing.T){

	user := entity.User{
		Username:"test_user", 
		Password:"test_password", 
		Email:sql.NullString{String:"test_email@163.com", Valid:false}, 
		Phone:sql.NullString{String:"12345678901", Valid:false}}

	if err1 := service.UserInfoService.InsertUser(user); err1 != nil {
		t.Fatal(err1)
	}

	if _, err2 := service.UserInfoService.FindByUsername(user.Username); err2 != nil{
		t.Fatal(err2)
	}

	session := entity.NewSession(entity.Session{LoginId:"", CurrentUser:"test_user"})
	
	if err3 := service.UserInfoService.LoginInfoInsert(*session); err3 != nil{
		t.Fatal(err3)
	}

	if _ , err6 := service.UserInfoService.UserHasLogin(session.LoginId); err6 == nil{
		_, err4 := service.UserInfoService.GetAllUsersInfo(*session)
		if (err4 != nil){
			t.Fatal(err4)
		}
	} else{
		t.Fatal(err6)
	}
	
	if  _ , err7 := service.UserInfoService.UserHasLogin(session.LoginId); err7 == nil{
		if err5 := service.UserInfoService.LoginInfoDelete(*session); err5 != nil{
			t.Fatal(err5)
		}
	} else {
		t.Fatal(err7)
	}
}
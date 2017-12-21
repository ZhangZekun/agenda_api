package server

import (
	"encoding/json"
	"net/http"
	"github.com/unrolled/render"
	"agenda_api/cli/service"
	"agenda_api/cli/entity"
)

var failMsg struct{
	Message string
	Data struct{
		Error string
	}
}
var successMsg struct{
	Message string
}

func isLoginHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("LoginId")
		session,err1 := service.UserInfoService.UserHasLogin(cookie.Value)
		if session != nil {
			user, _ := service.UserInfoService.FindByUsername(session.CurrentUser)
			if user != nil {
				var successMsg struct{
					Username string
				}
				successMsg.Username = user.Username
				formatter.JSON(w, http.StatusOK, successMsg)
			} else {
				var failMsg struct{
					Message string 
				}
				failMsg.Message = "no login"
				formatter.JSON(w, http.StatusUnauthorized, failMsg)
			}
		}
	}
}


func loginHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){

		failMsg.Message = "login fail"
		successMsg.Message = "login success"

		//check if login
		cookie, err := req.Cookie("LoginId")
		session,err1 := service.UserInfoService.UserHasLogin(cookie.Value)
		if session != nil {
			user, _ := service.UserInfoService.FindByUsername(session.CurrentUser)
			if user != nil {
				failMsg.Data.Error = "There is already a user logging in.You should logout first."
				formatter.JSON(w, http.StatusBadRequest, failMsg)
				return 
			}
		}

		var reqBody struct {
			Username string
			Password string
		}

		if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil{
			formatter.JSON(w, http.StatusUnprocessableEntity)
			return
		}

		//login
		user, _ := service.UserInfoService.FindByUsername(reqBody.Username)
		if user!= nil && user.Password == reqBody.Password{
			err2 := service.UserInfoService.InsertUser(user)
			if err2 != nil {
				failMsg.Data.Error = err2.String()
				formatter.JSON(w, http.StatusUnprocessableEntity, failMsg)
				return
			} else {
				var loginId string
				for {
					loginId_uuid := uuid.NewV4()
					loginId = loginId_uuid.String()
					s,_ := service.UserInfoService.UserHasLogin(loginId)
					if s == nil {
						break
					}
				}
				_ := service.UserInfoService.LoginInfoInsert(entity.Session{LoginId:loginId, CurrentUser:user.Username})
				set_cookie := http.Cookie{Name: "LoginId", Value: loginId, Path: "/", MaxAge: 86400}
				http.SetCookie(w, &set_cookie)
				formatter.JSON(w, http.StatusOK, successMsg)
			}
		}
	}
}

func logoutHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		failMsg.Message = "logout fail"
		successMsg.Message = "logout success"
		cookie, err := req.Cookie("LoginId")
		session,err1 := service.UserInfoService.UserHasLogin(cookie.Value)
		if session != nil {
			service.UserInfoService.LoginInfoDelete(session)
			del_cookie := http.Cookie{Name: "LoginId", Path: "/", MaxAge: -1}
			http.SetCookie(w, &del_cookie)
			formatter.JSON(w, http.StatusOK, successMsg)
			return
		} else {
			failMsg.Data.Error = "You don't login."
			formatter.JSON(w, http.StatusUnauthorized, failMsg)
			return
		}
	}
}

func registerHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){

		failMsg.Message = "register fail"
		successMsg.Message = "register success"

		var reqBody struct {
			Username string
			Password string
			Email string
			Phone string
		}

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil{
			formatter.JSON(w, http.StatusUnprocessableEntity, failMsg)
			return 
		}

		var user entity.User{
			Username:reqBody.Username,
			Password:reqBody.Password,
			Email:reqBody.Email,
			Phone:reqBody.Phone
		}

		if err1 := service.UserInfoService.InsertUser(user); err1 == nil {
			formatter.JSON(w, http.StatusOK, successMsg)
		} else {
			failMsg.Data.Error = err1
			formatter.JSON(w, http.StatusBadRequest, failMsg)
		}
	}
}

func deleteUserHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		
	}
}

func queryUsersHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		 
	}
}
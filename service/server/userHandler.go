package server

import (
	"net/http"
	"github.com/unrolled/render"
	"agenda_api/cli/service"
	"agenda_api/cli/entity"
)



func isLoginHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("LoginId")
		service.UserInfoService.UserHasLogin(cookie.Value)
	}
}


func loginHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
	}
}

func logoutHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){

	}
}

func registerHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){

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
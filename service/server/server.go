package server

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {

	//User Handler
	mx.HandleFunc("/api/agenda/user/login", isLoginHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/agenda/user/login", loginHandler(formatter)).Methods("POST")
	mx.HandleFunc("/api/agenda/user/allusers", queryUsersHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/agenda/user/register", registerHandler(formatter)).Methods("POST")
	mx.HandleFunc("/api/agenda/user/logout", logoutHandler(formatter)).Methods("POST")
	mx.HandleFunc("/api/agenda/user/self", deleteUserHandler(formatter)).Methods("DELETE")

	//Meeting Hanndler
//	mx.HandleFunc("/api/agenda/meeting", createMeetingHandler(formatter)).Methods("POST")
	mx.HandleFunc("/api/agenda/meeting/{title}/participators", addParticipatorsHandler(formatter)).Methods("POST")
	mx.HandleFunc("/api/agenda/meeting/{title}/particioators", deleteParticipatorsHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/api/agenda/meeting", queryMeetingHandler(formatter)).Methods("POST")
	mx.HandleFunc("/api/agenda/meeting/sponser/{title}", cncelMeetingHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/api/agenda/meeting/participator/{title}", quitMeetingHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/api/agenda/meeting/sponser/all", deleteAllMeetingHandler(formatter)).Methods("DELETE")
}

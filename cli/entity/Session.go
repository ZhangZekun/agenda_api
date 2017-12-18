package entity

type Session struct{
	Currentuser  string
}

func NewSession(s Session) *Session {
	if len(s.Currentuser) == 0 {
		panic("You have not logged in!")
	}
	return &s
}
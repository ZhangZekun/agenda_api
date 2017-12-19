package entity

type Session struct{
	LoginId  string `xorm:"pk"`
	CurrentUser string
}

func NewSession(s Session) *Session {
	if len(s.LoginId) == 0 {
		panic("You have not logged in!")
	}
	return &s
}
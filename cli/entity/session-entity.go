package entity

type Session struct{
	LoginId  string `xorm:"pk"`
	CurrentUser string
}

func NewSession(s Session) *Session {
	if len(s.CurrentUser) == 0 {
		panic("You have not enter your name!")
	}
	return &s
}
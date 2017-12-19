package entity

import (
//    "time"
	"database/sql"
)


type User struct {
	Username           string
	Password           string
	Email              sql.NullString
	Phone              sql.NullString
	SponsorMeeting     sql.NullString
	ParticipantMeeting sql.NullString
}

// NewUserInfo .
func NewUserInfo(u User) *User {
    if len(u.Username) == 0 {
        panic("Username shold not null!")
    }
    return &u
}
package entity

import (
    "time"
	"database/sql"
)


type Meeting struct {
	Id           	int
	Title        	string
	Sponsor     	string
	Participants 	sql.NullString
	StartTime		time.Time
	EndTime 		time.Time
}


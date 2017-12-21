package entity

import (
    "database/sql"

    _ "github.com/mattn/go-sqlite3"
)


var Mydb *sql.DB

var sqlCreateUserTable string = 
`CREATE TABLE IF NOT EXISTS User(
    Username VARCHAR(40) NULL,
    Password VARCHAR(40) NULL,
    Email    VARCHAR(40) NULL,
    Phone    VARCHAR(40) NULL,
    SponsorMeeting TEXT NULL,
    ParticipantMeeting TEXT NULL
    );`

var sqlCreateSessionTable string = 
`CREATE TABLE IF NOT EXISTS LoginUsers(
    LoginId VARCHAR(40) NULL,
    CurrentUser VARCHAR(40) NULL
    );`

var sqlCreateMeetingTable string = 
`CREATE TABLE IF NOT EXISTS Meeting(
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Title VARCHAR(40) NULL,
    Sponsor    VARCHAR(40) NULL,
    Participants TEXT NULL,
    StartTime DATETIME,
    EndTime DATETIME
    );`


func init() {
    //https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
    db, err := sql.Open("sqlite3", "../data/agenda.db")
    CheckErr(err)

    db.Exec(sqlCreateUserTable)
    db.Exec(sqlCreateSessionTable)
    Mydb = db
}

// SQLExecer interface for supporting sql.DB and sql.Tx to do sql statement
type SQLExecer interface {
    Exec(query string, args ...interface{}) (sql.Result, error)
    Prepare(query string) (*sql.Stmt, error)
    Query(query string, args ...interface{}) (*sql.Rows, error)
    QueryRow(query string, args ...interface{}) (*sql.Row)
}

// DaoSource Data Access Object Source
type DaoSource struct {
    // if DB, each statement execute sql with random conn.
    // if Tx, all statements use the same conn as the Tx's connection
    SQLExecer
}

func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}
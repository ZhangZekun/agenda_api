package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"agenda_api/cli/entity"
	"github.com/go-xorm/xorm")

var Enginea *xorm.Engine

func init() {
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	Enginea, _ = xorm.NewEngine("mysql", "root:zzm15331411@tcp(localhost:3306)/test?charset=utf8")
	err := Enginea.Ping()
	CheckErr(err)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "agenda_")
	Enginea.SetTableMapper(tbMapper) 
	var session entity.Session
	hasSessionTable, _ := Enginea.IsTableExist(&session)
	if (!hasSessionTable){
		Enginea.CreateTables(&session)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

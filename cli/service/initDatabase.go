package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm")

var Enginea *xorm.Engine

func init() {
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	Enginea, _ = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/agenda?charset=utf8")
	err := Enginea.Ping()
	CheckErr(err)
	Enginea.SetMapper(core.SameMapper{})
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

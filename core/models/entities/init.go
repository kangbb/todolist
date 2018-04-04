package entities

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func init() {
	var err error
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3308)/todolist?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	Engine = engine
	if os.Getenv("DEBUG") == "TRUE" {
		Engine.ShowSQL(true)
		Engine.Logger().SetLevel(core.LOG_DEBUG)
	}
}

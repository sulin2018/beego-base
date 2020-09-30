package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var ormClient orm.Ormer

func InitModels() {
	conn := beego.AppConfig.String("MySQLInfo")
	if conn == "" {
		panic("MySQLInfo cannot be empty")
	}
	err := orm.RegisterDataBase("default", "mysql", conn)
	if err != nil {
		logs.Error(err)
	}

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}

	orm.RegisterModel(new(User))
	// create table
	_ = orm.RunSyncdb("default", false, true)

	orm.SetMaxIdleConns("default", 15)
	orm.SetMaxOpenConns("default", 30)

	ormClient = orm.NewOrm()

	logs.Info("Init models success")
}

type User struct {
	Id       int
	UserName string
}

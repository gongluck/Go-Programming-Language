package model

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Pwd  string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "gongluck:gongluck@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

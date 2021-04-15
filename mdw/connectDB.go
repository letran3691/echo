package mdw

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"log"
)

func init()  {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	conectstring := "root:123456@tcp(10.10.0.111:3306)/orm_test"

	orm.RegisterDataBase("default","mysql",conectstring)

	name := "default"

	force := false

	verbose := true

	err := orm.RunSyncdb(name,force,verbose)

	if err != nil {
		log.Fatalf("Failed to connect DB %v",err)
	}
}
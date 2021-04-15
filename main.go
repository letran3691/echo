package main

import (
	"echo/handler"
	"echo/mdw"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

//func init()  {
//
//	connectionString := "root:123456@tcp(10.10.0.111:3306)/orm_test"
//
//	orm.RegisterDriver("mysql", orm.DRMySQL)
//
//	orm.RegisterDataBase("default", "mysql", connectionString)
//
//	// Database alias.
//	name := "default"
//
//	// Drop table and re-create.
//	force := false
//
//	// Print log.
//	verbose := true
//
//	// Error.
//	err := orm.RunSyncdb(name, force, verbose)
//	if err != nil {
//		log.Fatalf("Failed to run sync, error: %v",err)
//	}
//}


func main() {

	server := echo.New()

	server.Use(middleware.Logger())

	islogin := middleware.JWT(mdw.GetSecretKey())
	log.Printf("%v",islogin)
	isAdmin := mdw.IsAdminMdw
	server.GET("/",handler.Hello,islogin)
	server.POST("/login",handler.Login,middleware.BasicAuth(mdw.BacisAuth))
	server.GET("/admin",handler.Hello,islogin,isAdmin)

	//group API
	Groupv2 := server.Group("/v2")
	Groupv2.GET("/hello",handler.Hellov2)

	GroupUser := server.Group("/api/user")
	GroupUser.POST("/adduser",handler.AddUser)
	GroupUser.GET("/getuser",handler.GetUser)
	GroupUser.POST("/updateuser",handler.UpdateUser)
	GroupUser.DELETE("/deleteuser",handler.DeleteUser)

	server.Logger.Fatal(server.Start(":8088"))

}



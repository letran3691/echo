package handler

import (
	"echo/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"log"
	"net/http"

)

func Hello(c echo.Context) error  {

	userinfo := models.ParseUserInfo(c)
	log.Printf("hello %v",userinfo)
	log.Printf("%v",userinfo)

	if userinfo.Role{
		message := fmt.Sprintf("hello %s  is %v",userinfo.Username, userinfo.Role)

		hello := &models.Welcome{
			Text: message,
		}
		return c.JSON(http.StatusOK,hello)
	}else {
		return c.String(http.StatusOK,"is not admin")
	}

}

func Hellov2(c echo.Context)error  {
	return c.String(http.StatusOK,"Hellov2")

}


func AddUser(c echo.Context)error  {
	user := &models.User{}


	if err := c.Bind(user); err != nil{
		log.Printf("Binding user error %v", err)
		return err
	}

	log.Printf("user %v %T", user.Role,user.Role)
	if user.Password == "" || len(user.Password) < 6 {
		return c.String(http.StatusBadRequest,fmt.Sprint("can't password null"))
	}
	name := user.Name

	o := orm.NewOrm()

	exist := o.QueryTable("user").Filter("name",name).Exist()
	log.Printf("Is exist %v", exist)

	if exist == true {
		return c.String(http.StatusBadRequest,fmt.Sprintf("name is exist, enter orther name"))
	}

	hashpassw,_ := models.HashPassword(user.Password)

	userupdate := &models.User{
		Id: user.Id,
		Name: user.Name,
		Age: user.Age,
		Phone: user.Phone,
		Role: user.Role,
		Password: hashpassw,
	}


	id, err := o.Insert(userupdate)
	if err != nil {
		log.Printf("Insert user error %v", err)
		return c.String(http.StatusBadRequest,fmt.Sprint("err"))
		return err
	}

	log.Printf("Insert user  %d", id)

	return c.JSON(http.StatusOK,userupdate)
}


func GetUser(c echo.Context)error  {
	id := cast.ToInt64(c.QueryParam("id"))
	log.Printf("%v",id)
	o := orm.NewOrm()

	user := &models.User{
		Id : id,
	}
	err := o.Read(user)
	if err != nil {
		log.Printf("Get user error %v", err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}
func UpdateUser(c echo.Context)error  {
	user := &models.User{}

	if err := c.Bind(user); err != nil{
		log.Printf("update user error %v", err)
		return err
	}
	log.Printf("update %v",user)

	o := orm.NewOrm()

	_, err := o.Update(user,"name")

	if err != nil {
		log.Printf("update user %s error %v", user.Name, err)
		return err
	}
	userupdate := &models.User{
		Id: user.Id,
		Name: user.Name,
		Age: user.Age,
		Phone: user.Phone,
	}
	o.Read(userupdate)
	return c.JSON(http.StatusOK,userupdate)
}
func DeleteUser(c echo.Context)error  {
	//user := &models.User{}
	id := cast.ToInt64(c.QueryParam("id"))
	log.Printf("%v",id)

	o := orm.NewOrm()
	user := &models.User{
		Id : id,
	}
	log.Printf("%v",user)

	//_,err := o.Delete(&models.User{Id: id})
	_,err := o.Delete(user)


	if err != nil {
		log.Printf("dele user error %v", err)
		return err
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("Deleted id %d user",id))
}

//func GetAlllUser(c echo.Context) error  {
//
//	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSON)
//	c.Response().WriteHeader(http.StatusOK)
//	enc := json.NewEncoder(c.Response())
//
//
//	for _,user := range models.ListUser{
//		if err := enc.Encode(user); err !=nil{
//			return err
//		}
//		c.Response().Flush()
//		time.Sleep(1 * time.Second)
//	}
//	return nil
//}

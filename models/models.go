package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

)

type Welcome struct {
	Text string `json:"text"`
}

type UserLogin struct {
	Username string `json:"username" form:"username"`
	Role bool `json:"role" form:"role"`
}
type User struct {
	Id int64 `orm:"auto" json:"id" form:"id"`
	Name string `json:"name" form:"name" orm:"unique"`
	Password string `json:"password" form:"password" query:"password"`
	Age int 	`json:"age" form:"age"`
	Phone string	`json:"phone" form:"phone"`
	Role bool	`json:"role" form:"role"`
}

func init()  {
	orm.RegisterModel(new(User))
}


//var ListUser= []User{
//	{
//		Name: "trungkv",
//		Age: 33,
//	},
//	{
//		Name: "Huyen",
//		Age: 43,
//	},
//	{
//		Name: "Hang",
//		Age: 53,
//	},	{
//		Name: "Nam",
//		Age: 54,
//	},	{
//		Name: "Linh",
//		Age: 65,
//	},	{
//		Name: "Manh",
//		Age: 32,
//	},	{
//		Name: "Lam",
//		Age: 43,
//	},	{
//		Name: "Huan",
//		Age: 62,
//	},
//}


type LoginResponse struct {
	Token string `json:"token"`
}
//// Pasing token to data
func ParseUserInfo(c echo.Context) UserLogin  {
	user:= c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["name"].(string)
	role := claims["role"].(bool)

	userinfo := UserLogin{
		username, role,
	}
	return userinfo
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}


func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
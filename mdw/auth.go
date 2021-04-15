package mdw

import (
	"echo/models"
	"github.com/labstack/echo/v4"
	"github.com/astaxie/beego/orm"
	//"net/http"
	_"net/http"
	"log"
)

func BacisAuth(name string,password string,c echo.Context) (bool, error) {

	//id := cast.ToInt64(c.QueryParam("id"))
	//log.Printf("%v",id)
	log.Printf("name %v %v",name,password)

	o := orm.NewOrm()
	//
	user := models.User{

		Name: name,

	}

	//log.Printf("User %v",user)
	err := o.Read(&user, "name")
	if err != nil {
		log.Printf("auth user error %v", err)
	}

	log.Printf("dsadsa %v",user)
	//return c.JSON(http.StatusOK, user)
	if models.CheckPasswordHash(password,user.Password){
			c.Set("name",name)
			c.Set("role",user.Role)
			return true,nil
	}

	return false,nil

}

func GetSecretKey() []byte {
	return []byte("sadfaeg4t54yawfarwefef4taaff")
}
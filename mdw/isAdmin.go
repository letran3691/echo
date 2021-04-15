package mdw

import (
	"echo/models"
	"github.com/labstack/echo/v4"
)
//check user have been admin
func IsAdminMdw(next echo.HandlerFunc)echo.HandlerFunc  {
	return func(c echo.Context) error {
		userinfo := models.ParseUserInfo(c)

		if userinfo.Role{
			return next(c)
		}
		return echo.ErrUnauthorized
	}
}
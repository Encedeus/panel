package middleware

import (
	"github.com/labstack/echo/v4"
)

func CORSMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		origin := c.Request().Header.Get("Origin")

		//fmt.Println(origin)

		c.Response().Header().Set("Access-Control-Allow-Origin", origin)
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Accept")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET,HEAD,POST, PUT, DELETE, PATCH")

		if c.Request().Method == "OPTIONS" {
			return c.NoContent(200)
		}

		return next(c)
	}
}

package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"panel/util"
)

// RefreshJWTAuth serves as a middleware for authorization via the refresh token
func RefreshJWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check if the header is empty
		if c.Request().Header.Get("Authorization") == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "unauthorised",
			})
		}

		// extract and validate JWT
		token := util.GetTokenFromHeader(c)
		isValid, _, err := util.ValidateRefreshJWT(token)

		if !isValid || err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "unauthorised",
			})
		}

		return next(c)
	}
}

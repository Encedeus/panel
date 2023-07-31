package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"panel/util"
	"strings"
)

// RefreshJWTAuth serves as a middleware for authorization via the refresh token
func RefreshJWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check if cookie exists
		cookie, err := c.Request().Cookie("encedeus_refreshToken")
		if err != nil || strings.TrimSpace(cookie.Value) == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "unauthorised",
			})
		}

		// extract and validate JWT
		token := cookie.Value
		isValid, refreshToken, err := util.ValidateRefreshJWT(token)

		if !isValid || err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "unauthorised",
			})
		}

		c.Request().Header.Set("UUID", refreshToken.UserId.String())

		return next(c)
	}
}

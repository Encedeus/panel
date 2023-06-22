package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"panel/util"
)

// AccessJWTAuth serves as a middleware for authorization via the access token
func AccessJWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check if the header is empty
		if c.Request().Header.Get("Authorization") == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "unauthorised",
			})
		}

		// extract and validate JWT
		token := util.GetTokenFromHeader(c)
		isValid, _, err := util.ValidateAccessJWT(token)

		if !isValid {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "unauthorised",
			})
		}

		// responds with a status 500 if the error is not an invalid JWT
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "internal server error",
			})
		}

		return next(c)
	}
}

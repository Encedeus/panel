package middleware

import (
    "context"
    "github.com/Encedeus/panel/util"
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "net/http"
    "strings"
)

func ContextWithIDFromRefresh(ctx context.Context, refreshToken util.TokenClaims) context.Context {
    return context.WithValue(ctx, contextKey(1), refreshToken.UserID.String())
}

func IDFromRefreshContext(ctx context.Context) (uuid.UUID, error) {
    return uuid.Parse(ctx.Value(contextKey(1)).(string))
}

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

        c.SetRequest(c.Request().WithContext(ContextWithIDFromRefresh(c.Request().Context(), refreshToken)))

        return next(c)
    }
}

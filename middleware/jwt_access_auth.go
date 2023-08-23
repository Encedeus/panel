package middleware

import (
    "context"
    "github.com/Encedeus/panel/util"
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "net/http"
)

func ContextWithIDFromAccess(ctx context.Context, accessToken util.TokenClaims) context.Context {
    return context.WithValue(ctx, contextKey(2), accessToken.UserID.String())
}

func IDFromAccessContext(ctx context.Context) (uuid.UUID, error) {
    return uuid.Parse(ctx.Value(contextKey(2)).(string))
}

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
        isValid, accessToken, err := util.ValidateAccessJWT(token)

        if !isValid || err != nil {
            return c.JSON(http.StatusUnauthorized, echo.Map{
                "message": "unauthorised",
            })
        }

        c.SetRequest(c.Request().WithContext(ContextWithIDFromAccess(c.Request().Context(), accessToken)))

        return next(c)
    }
}

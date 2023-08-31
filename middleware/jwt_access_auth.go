package middleware

import (
    "context"
    "errors"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/ent/apikey"
    "github.com/Encedeus/panel/services"
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "net/http"
    "slices"
    "strings"
)

func ContextWithIDFromAccess(ctx context.Context, accessToken services.TokenClaims) context.Context {
    return context.WithValue(ctx, contextKey(2), accessToken.Token.UserId.Value)
}

func IDFromAccessContext(ctx context.Context) (uuid.UUID, error) {
    return uuid.Parse(ctx.Value(contextKey(2)).(string))
}

// AccessJWTAuth serves as a middleware for authorization via the access token
func AccessJWTAuth(db *ent.Client, next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // check if the header is empty
        if strings.TrimSpace(c.Request().Header.Get("Authorization")) == "" {
            return c.JSON(http.StatusUnauthorized, echo.Map{
                "message": "unauthorised",
            })
        }

        ctx := c.Request().Context()
        token := services.GetTokenFromHeader(c)

        isValid, apiKey, err := services.ValidateAccountAPIKey(token)
        if err != nil && !errors.Is(err, services.ErrInvalidTokenType) {
            return c.JSON(http.StatusUnauthorized, echo.Map{
                "message": "unauthorised",
            })
        }
        if isValid {
            _, err := db.ApiKey.Query().Where(apikey.KeyEQ(token)).First(ctx)
            if err != nil {
                return c.JSON(http.StatusUnauthorized, echo.Map{
                    "message": "unauthorised",
                })
            }

            ip := strings.Split(c.Request().RemoteAddr, ":")[0]
            if apiKey.IpAddresses != nil && len(apiKey.IpAddresses[0]) > 0 && !slices.Contains(apiKey.IpAddresses, ip) {
                return c.JSON(http.StatusUnauthorized, echo.Map{
                    "message": "unauthorised",
                })
            }
        }

        isValid, accessToken, err := services.ValidateAccessJWT(token)
        if err != nil && !errors.Is(err, services.ErrInvalidTokenType) {
            return c.JSON(http.StatusUnauthorized, echo.Map{
                "message": "unauthorised",
            })
        }

        c.SetRequest(c.Request().WithContext(ContextWithIDFromAccess(ctx, accessToken)))

        return next(c)
    }
}

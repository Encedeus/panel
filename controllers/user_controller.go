package controllers

import (
    "errors"
    "fmt"
    "github.com/Encedeus/panel/config"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/hashing"
    "github.com/Encedeus/panel/middleware"
    "github.com/Encedeus/panel/proto"
    protoapi "github.com/Encedeus/panel/proto/go"
    "github.com/Encedeus/panel/services"
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "github.com/labstack/gommon/log"
    "google.golang.org/protobuf/encoding/protojson"
    "io"
    "net/http"
    "os"
    "strings"
)

type UserController struct {
    Controller
}

func (uc UserController) registerRoutes(srv *Server) {
    userEndpoint := srv.Group("user")
    {
        userEndpoint.Static("/pfp", config.Config.CDN.Directory)

        userEndpoint.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
            return middleware.AccessJWTAuth(srv.DB, next)
        })

        userEndpoint.GET("/:id", func(c echo.Context) error {
            return handleFindUser(c, srv.DB)
        })
        userEndpoint.POST("", func(c echo.Context) error {
            return handleCreateUser(c, srv.DB)
        })
        userEndpoint.PUT("", func(c echo.Context) error {
            return handleSetPfp(c, srv.DB)
        })
        userEndpoint.PATCH("", func(c echo.Context) error {
            return handleUpdateUser(c, srv.DB)
        })
        userEndpoint.DELETE("/:id", func(c echo.Context) error {
            return handleDeleteUser(c, srv.DB)
        })
        userEndpoint.PATCH("/:id/changePassword", func(c echo.Context) error {
            return handleChangePassword(c, srv.DB)
        })
        userEndpoint.PATCH("/:id/changeUsername", func(c echo.Context) error {
            return handleChangeUsername(c, srv.DB)
        })
        userEndpoint.PATCH("/:id/changeEmail", func(c echo.Context) error {
            return handleChangeEmail(c, srv.DB)
        })
    }
}

func handleFindUser(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    rawUserId := c.Param("id")

    userId, err := uuid.Parse(rawUserId)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
    }

    resp, err := services.FindOneUser(ctx, db, &protoapi.UserFindOneRequest{
        UserId: proto.UUIDToProtoUUID(userId),
    })

    if err != nil {
        if ent.IsNotFound(err) {
            return c.JSON(http.StatusNotFound, echo.Map{"message": "user not found"})
        }
        if err.Error() == "user deleted" {
            return c.JSON(http.StatusGone, echo.Map{"message": "user deleted"})
        }

        log.Errorf("error querying user: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
}

func handleCreateUser(c echo.Context, db *ent.Client) error {
    // get uuid from header provided by the middleware
    ctx := c.Request().Context()
    authUUID, _ := middleware.IDFromAccessContext(ctx)
    // check permissions
    if !services.DoesUserHavePermission(ctx, db, "create_user", authUUID) {
        return c.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    createReq := new(protoapi.UserCreateRequest)
    err := c.Bind(createReq)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    // check if all the fields are provided
    if strings.TrimSpace(createReq.Name) == "" || strings.TrimSpace(createReq.Password) == "" || strings.TrimSpace(createReq.Email) == "" || (strings.TrimSpace(createReq.RoleName) == "" && strings.TrimSpace(createReq.RoleId.Value) == "") {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    resp, err := services.CreateUser(ctx, db, createReq)
    // error checking
    if err != nil {
        if services.IsValidationError(err) {
            return c.JSON(http.StatusBadRequest, echo.Map{
                "message": err.Error(),
            })
        }

        if ent.IsNotFound(err) {
            return c.JSON(http.StatusNotFound, echo.Map{
                "message": "role not found",
            })
        }

        if ent.IsConstraintError(err) {
            return c.JSON(http.StatusConflict, echo.Map{
                "message": "username taken",
            })
        }

        // log any uncaught errors
        log.Errorf("uncaught error querying role: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
}

func handleUpdateUser(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    authUUID, _ := middleware.IDFromAccessContext(ctx)

    // check permissions
    if !services.DoesUserHavePermission(ctx, db, "update_user", authUUID) {
        return c.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    updateReq := new(protoapi.UserUpdateRequest)
    err := c.Bind(updateReq)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    // check if no fields are provided
    if (updateReq.Name == "" && updateReq.Email == "" && updateReq.Password == "" && updateReq.RoleName == "" && updateReq.RoleId.Value == "") || updateReq.UserId.Value == "" {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    updateReq.Password = hashing.HashPassword(updateReq.Password)
    resp, err := services.UpdateUser(ctx, db, updateReq)

    // error checking
    if err != nil {
        if ent.IsNotFound(err) {
            if strings.Contains(err.Error(), "role") {
                return c.JSON(http.StatusNotFound, echo.Map{
                    "message": "role not found",
                })
            }
            return c.JSON(http.StatusNotFound, echo.Map{
                "message": "user not found",
            })
        }
        if ent.IsValidationError(err) || ent.IsConstraintError(err) || services.IsValidationError(err) {
            return c.JSON(http.StatusBadRequest, echo.Map{
                "message": "validate error",
            })
        }
        if err.Error() == "user deleted" {
            return c.JSON(http.StatusGone, echo.Map{
                "message": "user deleted",
            })
        }

        log.Errorf("uncaught error updating user: %v", err)
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
}

func handleSetPfp(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    // get params from multipart
    userId := c.FormValue("uuid")
    file, err := c.FormFile("file")

    // validate request
    if err != nil || userId == "" {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
    }

    userUUID, err := uuid.Parse(userId)
    if !services.DoesUserWithUUIDExist(ctx, db, userUUID) {
        return c.JSON(http.StatusNotFound, echo.Map{"message": "user not found"})
    }

    // open file
    src, err := file.Open()
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid file format"})
    }
    defer src.Close()

    // create file
    dst, err := os.Create(fmt.Sprintf("%s/%s", config.Config.CDN.Directory, userId))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
    }

    // write to file
    if _, err = io.Copy(dst, src); err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
    }

    return c.NoContent(http.StatusOK)
}

func handleDeleteUser(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    authUUID, _ := middleware.IDFromAccessContext(ctx)

    // check permissions
    if !services.DoesUserHavePermission(ctx, db, "delete_user", authUUID) {
        return c.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    rawUserId := c.Param("id")

    userId, err := uuid.Parse(rawUserId)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
    }

    _, err = services.DeleteUser(ctx, db, &protoapi.UserDeleteRequest{
        UserId: proto.UUIDToProtoUUID(userId),
    })

    // error checking
    if err != nil {
        if ent.IsNotFound(err) {
            return c.JSON(http.StatusNotFound, echo.Map{
                "message": "user not found",
            })
        }
        if err.Error() == "already deleted" {
            return c.JSON(http.StatusGone, echo.Map{
                "message": "user already deleted",
            })
        }

        log.Errorf("uncaught error deleting user: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return c.NoContent(http.StatusOK)
}

func handleChangePassword(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    authUUID, _ := middleware.IDFromAccessContext(ctx)

    // TODO: add ability for an application API key to change any user's information

    bytes := make([]byte, c.Request().ContentLength)
    _, err := c.Request().Body.Read(bytes)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    req := new(protoapi.UserChangePasswordRequest)
    err = protojson.Unmarshal(bytes, req)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": err.Error(),
        })
    }
    req.UserId = proto.UUIDToProtoUUID(authUUID)

    _, err = services.ChangeUserPassword(ctx, db, req)
    if err != nil {
        if services.IsValidationError(err) {
            return c.JSON(http.StatusBadRequest, echo.Map{
                "message": err.Error(),
            })
        }
        if errors.Is(err, services.ErrUserNotFound) {
            return c.JSON(http.StatusNotFound, echo.Map{
                "message": err.Error(),
            })
        }

        log.Errorf("uncaught error changing password: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return c.NoContent(http.StatusOK)
}

func handleChangeEmail(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    authUUID, _ := middleware.IDFromAccessContext(ctx)

    // TODO: add ability for an application API key to change any user's information

    bytes := make([]byte, c.Request().ContentLength)
    _, err := c.Request().Body.Read(bytes)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    req := new(protoapi.UserChangeEmailRequest)
    err = protojson.Unmarshal(bytes, req)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": err.Error(),
        })
    }
    req.UserId = proto.UUIDToProtoUUID(authUUID)

    _, err = services.ChangeUserEmail(ctx, db, req)
    if err != nil {
        if services.IsValidationError(err) {
            return c.JSON(http.StatusBadRequest, echo.Map{
                "message": err.Error(),
            })
        }

        log.Errorf("uncaught error changing email: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return c.NoContent(http.StatusOK)
}

func handleChangeUsername(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    authUUID, _ := middleware.IDFromAccessContext(ctx)

    // TODO: add ability for an application API key to change any user's information

    bytes := make([]byte, c.Request().ContentLength)
    _, err := c.Request().Body.Read(bytes)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    req := new(protoapi.UserChangeUsernameRequest)
    err = protojson.Unmarshal(bytes, req)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": err.Error(),
        })
    }
    req.UserId = proto.UUIDToProtoUUID(authUUID)

    _, err = services.ChangeUsername(ctx, db, req)
    if err != nil {
        if services.IsValidationError(err) {
            return c.JSON(http.StatusBadRequest, echo.Map{
                "message": err.Error(),
            })
        }

        log.Errorf("uncaught error changing username: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return c.NoContent(http.StatusOK)
}

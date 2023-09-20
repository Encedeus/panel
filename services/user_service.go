package services

import (
    "context"
    "errors"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/ent/role"
    "github.com/Encedeus/panel/ent/user"
    "github.com/Encedeus/panel/hashing"
    "github.com/Encedeus/panel/proto"
    protoapi "github.com/Encedeus/panel/proto/go"
    "github.com/Encedeus/panel/validate"
    "github.com/google/uuid"
    "golang.org/x/exp/slices"
    "strings"
    "time"
)

func CreateUser(ctx context.Context, db *ent.Client, req *protoapi.UserCreateRequest) (*protoapi.UserCreateResponse, error) {
    if !validate.IsUsername(req.Name) {
        return nil, ErrInvalidUsername
    }
    if !validate.IsEmail(req.Email) {
        return nil, ErrInvalidEmail
    }
    if !validate.IsPassword(req.Password) {
        return nil, ErrInvalidPassword
    }

    roleId, err := uuid.Parse(req.RoleId.Value)
    if err != nil {
        r, err := db.Role.Query().Where(role.NameEQ(req.RoleName)).First(ctx)
        if err != nil {
            return nil, err
        }

        roleId = r.ID
    }

    userData, err := db.User.Create().
        SetName(req.Name).
        SetEmail(req.Email).
        SetPassword(hashing.HashPassword(req.Password)).
        SetRoleID(roleId).
        Save(ctx)

    resp := &protoapi.UserCreateResponse{
        User: proto.EntUserEntityToProtoUser(userData),
    }

    return resp, nil
}

// DoesUserHavePermission checks if user's role have a permission
func DoesUserHavePermission(ctx context.Context, db *ent.Client, permission string, userID uuid.UUID) bool {
    userData, err := db.User.Query().Where(user.IDEQ(userID)).Select("role_id").First(ctx)
    if err != nil {
        return false
    }

    if IsUserDeleted(userData) {
        return false
    }

    roleData, err := db.Role.Query().Where(role.ID(userData.RoleID)).Select("permissions").First(ctx)
    if err != nil {
        return false
    }

    return slices.Contains(roleData.Permissions, permission) || slices.Contains(roleData.Permissions, "*")
}

// UpdateUser updates the user given an updateInfo dto
func UpdateUser(ctx context.Context, db *ent.Client, req *protoapi.UserUpdateRequest) (*protoapi.UserUpdateResponse, error) {
    if !validate.IsEmail(req.Email) {
        return nil, ErrInvalidEmail
    }
    if !validate.IsPassword(req.Password) {
        return nil, ErrInvalidPassword
    }
    if !validate.IsUsername(req.Name) {
        return nil, ErrInvalidUserId
    }

    userData, err := db.User.Query().Where(user.IDEQ(proto.ProtoUUIDToUUID(req.UserId))).First(ctx)

    if err != nil {
        return nil, err
    }

    if IsUserDeleted(userData) {
        return nil, errors.New("user deleted")
    }

    if req.Name != "" {
        _, err = userData.Update().SetName(req.Name).Save(ctx)
    }

    if req.Password != "" {
        _, err = userData.Update().SetPassword(req.Password).Save(ctx)
    }

    if req.Email != "" {
        _, err = userData.Update().SetEmail(req.Email).Save(ctx)
    }

    if req.RoleName != "" {
        roleData, roleErr := db.Role.Query().Where(role.Name(req.RoleName)).First(ctx)
        if roleErr != nil {
            return nil, roleErr
        }
        _, err = userData.Update().SetRole(roleData).Save(ctx)

    }

    if s := req.RoleId.Value; s != "" {
        roleData, roleErr := db.Role.Query().Where(role.ID(proto.ProtoUUIDToUUID(req.RoleId))).First(ctx)
        if roleErr != nil {
            return nil, roleErr
        }
        _, err = userData.Update().SetRole(roleData).Save(ctx)
    }

    currUser, err := db.User.Get(ctx, proto.ProtoUUIDToUUID(req.UserId))
    if err != nil {
        return nil, err
    }

    resp := &protoapi.UserUpdateResponse{
        User: proto.EntUserEntityToProtoUser(currUser),
    }

    return resp, nil
}

func DeleteUser(ctx context.Context, db *ent.Client, req *protoapi.UserDeleteRequest) (*protoapi.UserDeleteResponse, error) {
    userData, err := db.User.Query().Where(user.IDEQ(uuid.MustParse(req.UserId.Value))).First(ctx)
    if err != nil {
        return nil, err
    }

    if IsUserDeleted(userData) {
        return nil, errors.New("already deleted")
    }

    userData, err = userData.Update().SetDeletedAt(time.Now()).Save(ctx)
    if err != nil {
        return nil, err
    }

    resp := &protoapi.UserDeleteResponse{}

    return resp, err
}

func FindOneUser(ctx context.Context, db *ent.Client, req *protoapi.UserFindOneRequest) (*protoapi.UserFindOneResponse, error) {
    userData, err := db.User.Query().
        Where(user.IDEQ(proto.ProtoUUIDToUUID(req.UserId))).
        Select("id", "name", "created_at", "updated_at", "deleted_at", "email", "role_id").
        First(ctx)
    if err != nil {
        return nil, err
    }

    if IsUserDeleted(userData) {
        return nil, errors.New("user deleted")
    }

    resp := &protoapi.UserFindOneResponse{
        User: proto.EntUserEntityToProtoUser(userData),
    }

    return resp, err
}

func DoesUserWithUUIDExist(ctx context.Context, db *ent.Client, userId uuid.UUID) bool {
    userData, err := db.User.Query().Where(user.IDEQ(userId)).First(ctx)

    if err != nil || IsUserDeleted(userData) {
        return false
    }

    return true
}

func GetLastUpdate(ctx context.Context, db *ent.Client, userId uuid.UUID) (int64, error) {
    userData, err := db.User.Query().
        Where(user.IDEQ(userId), user.DeletedAtIsNil()).
        Select("updated_at").
        First(ctx)

    if err != nil {
        return 0, nil
    }

    return userData.UpdatedAt.Unix(), nil
}
func IsUserDeleted(userData *ent.User) bool {
    return userData.DeletedAt.Unix() != -62135596800
}

func IsUserUpdated(ctx context.Context, db *ent.Client, userId uuid.UUID, issuedAt int64) (bool, error) {
    lastUpdate, err := GetLastUpdate(ctx, db, userId)
    if err != nil {
        return true, err
    }

    if lastUpdate > issuedAt {
        return true, nil
    }

    return false, nil
}

func ChangeUsername(ctx context.Context, db *ent.Client, req *protoapi.UserChangeUsernameRequest) (*protoapi.UserChangeUsernameResponse, error) {
    if !validate.IsUsername(req.NewUsername) || !validate.IsUsername(req.OldUsername) {
        return nil, ErrInvalidUsername
    }
    if !validate.IsUserId(ctx, db, req.UserId) {
        return nil, ErrInvalidUserId
    }

    userData, err := db.User.Query().Where(user.IDEQ(proto.ProtoUUIDToUUID(req.UserId))).Select(user.FieldName).First(ctx)
    if err != nil {
        if ent.IsNotFound(err) {
            return nil, ErrUserNotFound
        }

        return nil, err
    }
    if userData.Name != strings.TrimSpace(req.OldUsername) {
        return nil, ErrOldUsernameDoesNotMatch
    }
    if userData.Name == strings.TrimSpace(req.NewUsername) {
        return nil, ErrNewUsernameEqualsOld
    }

    _, err = userData.Update().SetName(req.NewUsername).Save(ctx)
    if err != nil {
        if ent.IsConstraintError(err) {
            return nil, ErrUsernameAlreadyTaken
        }

        return nil, err
    }

    resp := &protoapi.UserChangeUsernameResponse{}

    return resp, nil
}

func ChangeUserPassword(ctx context.Context, db *ent.Client, req *protoapi.UserChangePasswordRequest) (*protoapi.UserChangePasswordResponse, error) {
    if !validate.IsPassword(req.NewPassword) || !validate.IsPassword(req.OldPassword) {
        return nil, ErrInvalidPassword
    }
    if !validate.IsUserId(ctx, db, req.UserId) {
        return nil, ErrInvalidUserId
    }

    userData, err := db.User.Query().Where(user.IDEQ(proto.ProtoUUIDToUUID(req.UserId))).Select(user.FieldPassword).First(ctx)
    if err != nil {
        return nil, err
    }
    if !hashing.VerifyHash(req.OldPassword, userData.Password) {
        return nil, ErrOldPasswordDoesNotMatch
    }
    if userData.Password == hashing.HashPassword(req.NewPassword) {
        return nil, ErrNewPasswordEqualsOld
    }

    _, err = userData.Update().SetPassword(hashing.HashPassword(req.NewPassword)).Save(ctx)
    if err != nil {
        return nil, err
    }

    resp := &protoapi.UserChangePasswordResponse{}

    return resp, nil
}

func ChangeUserEmail(ctx context.Context, db *ent.Client, req *protoapi.UserChangeEmailRequest) (*protoapi.UserChangeEmailResponse, error) {
    if !validate.IsEmail(req.NewEmail) || !validate.IsEmail(req.OldEmail) {
        return nil, ErrInvalidEmail
    }
    if !validate.IsUserId(ctx, db, req.UserId) {
        return nil, ErrInvalidUserId
    }

    userData, err := db.User.Query().Where(user.IDEQ(proto.ProtoUUIDToUUID(req.UserId))).Select(user.FieldEmail).First(ctx)
    if err != nil {
        return nil, err
    }
    if userData.Email != strings.TrimSpace(req.OldEmail) {
        return nil, ErrOldEmailDoesNotMatch
    }
    if userData.Email == strings.TrimSpace(req.NewEmail) {
        return nil, ErrNewEmailEqualsOld
    }

    _, err = userData.Update().SetEmail(req.NewEmail).Save(ctx)
    if err != nil {
        if ent.IsConstraintError(err) {
            return nil, ErrEmailAlreadyTaken
        }
        
        return nil, err
    }

    resp := &protoapi.UserChangeEmailResponse{}

    return resp, nil
}

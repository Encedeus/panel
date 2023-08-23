package services

import (
    "context"
    "errors"
    "github.com/Encedeus/panel/dto"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/ent/role"
    "github.com/Encedeus/panel/ent/user"
    "github.com/google/uuid"
    "golang.org/x/exp/slices"
    "time"
)

func CreateUserRoleId(ctx context.Context, db *ent.Client, name string, email string, passwordHash string, roleId uuid.UUID) (*uuid.UUID, error) {

    roleData, err := db.Role.Get(ctx, roleId)

    if err != nil {
        return nil, err
    }

    return CreateUser(ctx, db, name, email, passwordHash, roleData)
}

func CreateUserRoleName(ctx context.Context, db *ent.Client, name string, email string, passwordHash string, roleName string) (*uuid.UUID, error) {
    roleData, err := db.Role.Query().Where(role.Name(roleName)).First(ctx)

    if err != nil {
        return nil, err
    }

    return CreateUser(ctx, db, name, email, passwordHash, roleData)
}

func CreateUser(ctx context.Context, db *ent.Client, name string, email string, passwordHash string, role *ent.Role) (*uuid.UUID, error) {
    userData, err := db.User.Create().
        SetName(name).
        SetEmail(email).
        SetPassword(passwordHash).
        SetRole(role).
        Save(ctx)

    if err != nil {
        return nil, err
    }

    return &userData.ID, err
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
func UpdateUser(ctx context.Context, db *ent.Client, updateInfo dto.UpdateUserDTO) error {
    userData, err := db.User.Query().Where(user.IDEQ(updateInfo.UserId)).First(ctx)

    if err != nil {
        return err
    }

    if IsUserDeleted(userData) {
        return errors.New("user deleted")
    }

    if updateInfo.Name != "" {
        _, err = userData.Update().SetName(updateInfo.Name).Save(ctx)
    }

    if updateInfo.Password != "" {
        _, err = userData.Update().SetPassword(updateInfo.Password).Save(ctx)
    }

    if updateInfo.Email != "" {
        _, err = userData.Update().SetEmail(updateInfo.Email).Save(ctx)
    }

    if updateInfo.RoleName != "" {
        roleData, roleErr := db.Role.Query().Where(role.Name(updateInfo.RoleName)).First(ctx)
        if roleErr != nil {
            return roleErr
        }
        _, err = userData.Update().SetRole(roleData).Save(ctx)

    }

    if s := updateInfo.RoleId.String(); s != "" {
        roleData, roleErr := db.Role.Query().Where(role.ID(updateInfo.RoleId)).First(ctx)
        if roleErr != nil {
            return roleErr
        }
        _, err = userData.Update().SetRole(roleData).Save(ctx)
    }

    return err
}

func DeleteUser(ctx context.Context, db *ent.Client, userId uuid.UUID) error {
    userData, err := db.User.Query().Where(user.IDEQ(userId)).First(ctx)
    if err != nil {
        return err
    }

    if IsUserDeleted(userData) {
        return errors.New("already deleted")
    }

    userData, err = userData.Update().SetDeletedAt(time.Now()).Save(ctx)
    if err != nil {
        return err
    }

    return err
}

func GetUser(ctx context.Context, db *ent.Client, userId uuid.UUID) (*ent.User, error) {
    userData, err := db.User.Query().
        Where(user.IDEQ(userId)).
        Select("uuid", "name", "created_at", "updated_at", "deleted_at", "email", "role_id").
        First(ctx)
    if err != nil {
        return nil, err
    }

    if IsUserDeleted(userData) {
        return nil, errors.New("user deleted")
    }

    return userData, err
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

package services

import (
    "context"
    "errors"
    "github.com/Encedeus/panel/dto"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/ent/apikey"
    "github.com/Encedeus/panel/util"
    "github.com/google/uuid"
)

func CreateAccountAPIKey(ctx context.Context, db *ent.Client, keyData dto.AccountAPIKeyDTO) (apiKey *ent.ApiKey, err error) {
    key, err := util.GenerateAPIKey(keyData)
    if err != nil {
        return nil, err
    }

    apiKey, err = db.ApiKey.Create().
        SetKey(key).
        SetDescription(keyData.Description).
        SetIPAddresses(keyData.IPAddresses).
        SetUserID(keyData.UserID).
        Save(ctx)
    if err != nil {
        return nil, err
    }

    return apiKey, nil
}

func DeleteAccountAPIKey(ctx context.Context, db *ent.Client, keyId uuid.UUID) (err error) {
    if s := keyId.String(); s == "" {
        return errors.New("missing API key")
    }

    err = db.ApiKey.DeleteOneID(keyId).Exec(ctx)

    return err
}

func FindAccountAPIKeysByUserID(ctx context.Context, db *ent.Client, userId uuid.UUID) (apiKeys []*ent.ApiKey, err error) {
    apiKeys, err = db.ApiKey.Query().Where(apikey.UserIDEQ(userId)).All(ctx)
    if err != nil {
        return nil, err
    }

    return apiKeys, nil
}

func FindAccountAPIKeyByID(ctx context.Context, db *ent.Client, keyId uuid.UUID) (apiKey *ent.ApiKey, err error) {
    apiKey, err = db.ApiKey.Get(ctx, keyId)
    if err != nil {
        return nil, err
    }

    return apiKey, nil
}

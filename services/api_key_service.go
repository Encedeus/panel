package services

import (
    "context"
    "entgo.io/ent/dialect/sql"
    "errors"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/ent/apikey"
    "github.com/Encedeus/panel/proto"
    protoapi "github.com/Encedeus/panel/proto/go"
    "github.com/Encedeus/panel/util"
    "github.com/google/uuid"
    "google.golang.org/protobuf/types/known/timestamppb"
    "strings"
)

func CreateAccountAPIKey(ctx context.Context, db *ent.Client, req *protoapi.AccountAPIKeyCreateRequest) (resp *protoapi.AccountAPIKeyCreateResponse, err error) {
    key, err := util.GenerateAPIKey(proto.ProtoAccountAPIKeyCreateRequestToToken(req))
    if err != nil {
        return nil, err
    }

    apiKey, err := db.ApiKey.Create().
        SetKey(key).
        SetDescription(req.Description).
        SetIPAddresses(req.IpAddresses).
        SetUserID(uuid.MustParse(req.UserId.Value)).
        Save(ctx)
    if err != nil {
        return nil, err
    }

    resp = &protoapi.AccountAPIKeyCreateResponse{
        AccountApiKey: &protoapi.AccountAPIKey{
            Id:          proto.UUIDToProtoUUID(apiKey.ID),
            CreatedAt:   timestamppb.New(apiKey.CreatedAt),
            UpdatedAt:   timestamppb.New(apiKey.UpdatedAt),
            Description: apiKey.Description,
            IpAddresses: apiKey.IPAddresses,
            UserId:      proto.UUIDToProtoUUID(apiKey.UserID),
            Key:         apiKey.Key,
        },
    }

    return resp, nil
}

func DeleteAccountAPIKey(ctx context.Context, db *ent.Client, req *protoapi.AccountAPIKeyDeleteRequest) (resp *protoapi.AccountAPIKeyDeleteResponse, err error) {
    if s := req.Id.Value; strings.TrimSpace(s) == "" {
        return nil, errors.New("missing API key")
    }

    err = db.ApiKey.DeleteOneID(proto.ProtoUUIDToUUID(req.Id)).Exec(ctx)

    if err != nil {
        return nil, err
    }

    resp = &protoapi.AccountAPIKeyDeleteResponse{}

    return resp, nil
}

func FindAccountAPIKeysByUserID(ctx context.Context, db *ent.Client, req *protoapi.AccountAPIkeyFindManyByUserRequest) (resp *protoapi.AccountAPIkeyFindManyResponse, err error) {
    apiKeys, err := db.ApiKey.Query().Where(apikey.UserIDEQ(proto.ProtoUUIDToUUID(req.UserId))).Order(apikey.ByCreatedAt(sql.OrderDesc())).All(ctx)
    if err != nil {
        return nil, err
    }

    resp = &protoapi.AccountAPIkeyFindManyResponse{}
    resp.AccountApiKeys = make([]*protoapi.AccountAPIKey, len(apiKeys))
    for i, apiKey := range apiKeys {
        resp.AccountApiKeys[i] = proto.EntAccountAPIKeyToProtoKey(apiKey)
    }

    return resp, nil
}

func FindAccountAPIKeyByID(ctx context.Context, db *ent.Client, req *protoapi.AccountAPIKeyFindOneRequest) (resp *protoapi.AccountAPIKeyFindOneResponse, err error) {
    apiKey, err := db.ApiKey.Get(ctx, proto.ProtoUUIDToUUID(req.Id))
    if err != nil {
        return nil, err
    }

    resp = &protoapi.AccountAPIKeyFindOneResponse{
        AccountApiKey: proto.EntAccountAPIKeyToProtoKey(apiKey),
    }

    return resp, nil
}

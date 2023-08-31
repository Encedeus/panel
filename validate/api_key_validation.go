package validate

import (
    "context"
    "github.com/Encedeus/panel/ent"
    protoapi "github.com/Encedeus/panel/proto/go"
    "github.com/google/uuid"
    "github.com/microcosm-cc/bluemonday"
    "net/netip"
    "strings"
)

func IsIPAddress(address string) bool {
    if len(strings.TrimSpace(address)) == 0 {
        return true
    }
    _, err := netip.ParseAddr(address)

    return err == nil
}

func IsIPAddressList(address []string) bool {
    for _, ip := range address {
        if !IsIPAddress(ip) {
            return false
        }
    }

    return true
}

func IsAPIKeyDescription(description string) bool {
    p := bluemonday.StrictPolicy()

    if s := p.Sanitize(description); s != description {
        return false
    }
    if len(description) > 28 {
        return false
    }

    return true
}

func IsUserId(ctx context.Context, db *ent.Client, userId *protoapi.UUID) bool {
    id, err := uuid.Parse(userId.Value)
    if err != nil {
        return false
    }

    _, err = db.User.Get(ctx, id)
    if err != nil {
        return false
    }

    return true
}

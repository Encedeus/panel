package validate

import (
    "context"
    "github.com/Encedeus/panel/ent"
    protoapi "github.com/Encedeus/panel/proto/go"
    "github.com/google/uuid"
    "github.com/microcosm-cc/bluemonday"
    "net/netip"
)

func IsIPAddress(address string) bool {
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
    _, err := db.User.Get(ctx, uuid.MustParse(userId.Value))

    return err == nil
}

package validate

import (
    "github.com/microcosm-cc/bluemonday"
    "net/mail"
)

func IsUsername(username string) bool {
    if len(username) > 24 || len(username) < 3 {
        return false
    }

    p := bluemonday.StrictPolicy()
    if s := p.Sanitize(username); s != username {
        return false
    }

    return true
}

func IsEmail(email string) bool {
    _, err := mail.ParseAddress(email)

    return err == nil
}

func IsPassword(password string) bool {
    if len(password) > 64 || len(password) < 8 {
        return false
    }

    p := bluemonday.StrictPolicy()
    if s := p.Sanitize(password); s != password {
        return false
    }

    return true
}

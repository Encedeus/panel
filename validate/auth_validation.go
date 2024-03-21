package validate

import (
    "github.com/microcosm-cc/bluemonday"
    "net/http"
    "net/mail"
    "strings"
    "time"
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
    if err != nil {
        return false
    }

    domain := strings.Split(email, "@")[1]
    cli := http.Client{
        Timeout: 5 * time.Second,
    }

    ch := make(chan error, 1)
    defer close(ch)
    go func() {
        _, err = cli.Get("http://" + domain)
        ch <- err
    }()

    if <-ch != nil {
        return false
    }

    return true
}

func IsPassword(password string) bool {
    // if len(password) > 64 || len(password) < 8 {
    //     return false
    // }

    p := bluemonday.StrictPolicy()
    if s := p.Sanitize(password); s != password {
        return false
    }

    return true
}

package security

import (
    "crypto/rand"
    "golang.org/x/crypto/bcrypt"
    _ "golang.org/x/crypto/bcrypt"
)

func RandomBytes(length uint32) []byte {
    bytes := make([]byte, length)
    _, err := rand.Read(bytes)
    if err != nil {
        return nil
    }

    return bytes
}

func HashPassword(password string) string {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }

    return string(hashedPassword)
}

func VerifyHash(unhashed string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(unhashed))

    return err == nil
}

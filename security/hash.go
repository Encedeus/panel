package security

import (
    "crypto/rand"
    "encoding/base64"
    "errors"
    "fmt"
    "golang.org/x/crypto/argon2"
    _ "golang.org/x/crypto/bcrypt"
    "strings"
)

var (
    ErrInvalidHash                = errors.New("hashed password format is invalid")
    ErrPasswordsDoNotMatch        = errors.New("password does not match stored password")
    ErrInvalidVersion             = errors.New("invalid argon2 version")
    ErrIncompatibleVersion        = errors.New("incompatible argon2 version")
    ErrInvalidSalt                = errors.New("invalid argon2 salt")
    ErrInvalidOperationParameters = errors.New("invalid argon2 operation params")
    ErrInvalidKey                 = errors.New("invalid argon2 key")
)

const PHCStringFormat = "$argon2id$v=%v$m=%v,t=%v,p=%v$%v$%v"

func RandomBytes(length uint32) []byte {
    bytes := make([]byte, length)
    _, err := rand.Read(bytes)
    if err != nil {
        return nil
    }

    return bytes
}

type Argon2Parameters struct {
    Salt      []byte
    Time      uint32
    Memory    uint32
    Version   uint8
    Threads   uint8
    KeyLength uint32
}

func DefaultArgon2Params() Argon2Parameters {
    // Param recommendations taken from RFC 9106
    params := Argon2Parameters{
        // 128-bit salt
        Salt: RandomBytes(16),
        // Iteration of 3
        Time: 3,
        // 64 MiB memory in KiB
        Memory: 64 * 1024,
        // 0x13 = 19
        Version: argon2.Version,
        // Number of parallel execution chains
        Threads: 4,
        // Length of derived key
        KeyLength: 32,
    }

    return params
}

/*func HashPassword(password string) string {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }

    return string(hashedPassword)
}*/

func HashPassword(params Argon2Parameters, password string) string {
    key := argon2.IDKey(
        []byte(password),
        params.Salt,
        params.Time,
        params.Memory,
        params.Threads,
        params.KeyLength,
    )

    enc := base64.StdEncoding.WithPadding(base64.NoPadding)
    encodedSalt := enc.EncodeToString(params.Salt)
    encodedKey := enc.EncodeToString(key)

    phcHashFormat := fmt.Sprintf(
        PHCStringFormat,
        params.Version,
        params.Memory,
        params.Time,
        params.Threads,
        encodedSalt,
        encodedKey,
    )

    return phcHashFormat
}

func VerifyPasswordHash(password, hashedPassword string) (bool, error) {
    params, _, err := decodeHash(hashedPassword)
    if err != nil {
        return false, err
    }

    attemptedHash := HashPassword(params, password)
    if attemptedHash != hashedPassword {
        return false, ErrPasswordsDoNotMatch
    }

    return true, nil
}

func decodeHash(hash string) (params Argon2Parameters, key string, err error) {
    vals := strings.Split(hash, "$")
    if len(vals) != 6 {
        return params, key, ErrInvalidHash
    }
    params = DefaultArgon2Params()

    var version int
    _, err = fmt.Sscanf(vals[2], "v=%d", &version)
    if err != nil {
        return params, key, fmt.Errorf("%e: %e", ErrInvalidVersion, err)
    }
    if version != argon2.Version {
        return params, key, ErrIncompatibleVersion
    }
    params.Version = uint8(version)

    _, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Time, &params.Threads)
    if err != nil {
        return params, key, ErrInvalidOperationParameters
    }

    enc := base64.StdEncoding.WithPadding(base64.NoPadding)

    salt, err := enc.DecodeString(vals[4])
    if err != nil {
        return params, key, ErrInvalidSalt
    }
    params.Salt = salt

    keyB, err := enc.DecodeString(vals[5])
    if err != nil {
        return params, key, ErrInvalidKey
    }
    key = string(keyB)
    params.KeyLength = uint32(len(key))

    return params, key, nil
}

/*func VerifyHash(unhashed string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(unhashed))

    return err == nil
}*/

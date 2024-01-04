package security

import "testing"

func TestHashPassword(t *testing.T) {
    cases := []struct {
        params       Argon2Parameters
        password     string
        expectedHash string
    }{
        {
            params:       defaultTestArgon2Params("iY6961)Wt{A9vsYG"),
            password:     ".v66HW0@q^j0",
            expectedHash: "$argon2id$v=19$m=100000,t=1,p=4$aVk2OTYxKVd0e0E5dnNZRw$Za8/q69sZhBkZfO0mcg3+LgWEmovTKGAbt2PVRcefS0",
        },
        {
            params:       defaultTestArgon2Params("9k0Q<(Nr~2%J\"P=z"),
            password:     "123456789",
            expectedHash: "$argon2id$v=19$m=100000,t=1,p=4$OWswUTwoTnJ+MiVKIlA9eg$DvNiVFOV8r6FIabmFF1yFMDR8VqFcRrB2uQ0v4OBIEE",
        },
    }

    for _, c := range cases {
        got := HashPassword(c.params, c.password)
        if got != c.expectedHash {
            t.Fatalf("hashing failed, got: %v, expected: %v", got, c.expectedHash)
        }
    }
}

func TestVerifyHashPassword(t *testing.T) {
    cases := []struct {
        password         string
        hashedPassword   string
        expectedValidity bool
    }{
        {
            password:         ".v66HW0@q^j0",
            hashedPassword:   "$argon2id$v=19$m=100000,t=1,p=4$aVk2OTYxKVd0e0E5dnNZRw$Za8/q69sZhBkZfO0mcg3+LgWEmovTKGAbt2PVRcefS0",
            expectedValidity: true,
        },
        {
            password:         "123456789",
            hashedPassword:   "$argon2id$v=19$m=100000,t=1,p=4$OWswUTwoTnJ+MiVKIlA9eg$DvNiVFOV8r6FIabmFF1yFMDR8VqFcRrB2uQ0v4OBIEE",
            expectedValidity: true,
        },
        {
            password:         ".v66HW0@q^Am",
            hashedPassword:   "$argon2id$v=19$m=100000,t=1,p=4$OWswUTwoTnJ+MiVKIlB6$IcbabJ8eJZrtOK7zIUodh0xubGWnbnI6/v0DSJaYklo",
            expectedValidity: false,
        },
    }

    for _, c := range cases {
        got, err := VerifyPasswordHash(c.password, c.hashedPassword)
        if got != c.expectedValidity {
            t.Fatalf("passwords do not match, got: %v, expected: %v, error: %e, hash: %v", got, c.expectedValidity, err, c.hashedPassword)
        }
    }
}

func defaultTestArgon2Params(salt string) Argon2Parameters {
    params := DefaultArgon2Params()
    params.Salt = []byte(salt)
    params.Memory = 100000

    return params
}

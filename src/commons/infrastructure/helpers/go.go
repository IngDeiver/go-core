package helpers

import "golang.org/x/crypto/bcrypt"

func CreateHash(value string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

// return true if is valid
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
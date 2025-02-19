package auth

import "golang.org/x/crypto/bcrypt"

// VerifyPassword compares a hashed password with its possible plaintext equivalent.
func VerifyPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
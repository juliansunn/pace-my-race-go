package util

import "golang.org/x/crypto/bcrypt"

func HashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "failed to hash password", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a plaintext password with a hashed password and returns an error if they do not match.
//
// Parameters:
// - password: a string representing the plaintext password to be checked.
// - hashedPassword: a string representing the hashed password to compare against.
//
// Returns:
// - error: an error indicating whether the passwords match or not.
func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

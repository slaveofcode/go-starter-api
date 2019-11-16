package password

import "golang.org/x/crypto/bcrypt"

// SaltRounds how much round for salt
const SaltRounds = 14

// Hash create password hash
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), SaltRounds)
	return string(bytes), err
}

// Compare will compare string password with the hash
func Compare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package password

import "golang.org/x/crypto/bcrypt"

// Hash hashes password
func Hash(pass string)  (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), 11)
	return string(hashed), err
}

// Compare compares hashed and password
func Compare (hashed, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass)) == nil
}

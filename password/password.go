package password

import (
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
	"io"
	"log"
	"unicode"
)

const (
	PW_SALT_BYTES = 512
	PW_HASH_BYTES = 1024
)

var mustHave = []func(rune) bool{
	unicode.IsUpper,
	unicode.IsLower,
	unicode.IsPunct,
	unicode.IsDigit,
}

// Encrypt a password string
func EncryptPassword(password string) ([]byte, []byte) {
	salt := make([]byte, PW_SALT_BYTES)

	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Fatal(err)
	}

	hash, err := getHash([]byte(password), salt)
	if err != nil {
		log.Fatal(err)
	}

	return salt, hash
}

// Check if password is valid
func IsPasswordValid(password string, salt string, hash string) (bool, error) {
	goodHash, err := getHash([]byte(password), []byte(salt))

	if err != nil {
		return false, err
	}

	isValid := (string(goodHash) == hash)

	return isValid, nil
}

// Hash password
func getHash(password []byte, salt []byte) ([]byte, error) {
	return scrypt.Key(password, salt, 1<<14, 8, 1, PW_HASH_BYTES)
}

func CheckPasswordRequirements(p string) bool {
	for _, testRune := range mustHave {
		found := false
		for _, r := range p {
			if testRune(r) {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

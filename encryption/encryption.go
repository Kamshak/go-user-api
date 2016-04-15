package encryption

import (
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
	"io"
	"log"
)

const (
	PW_SALT_BYTES = 32
	PW_HASH_BYTES = 64
)

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

func IsPasswordValid(password string, salt string, hash string) (bool, error) {
	goodHash, err := getHash([]byte(password), []byte(salt))

	if err != nil {
		return false, err
	}

	isValid := (string(goodHash) == hash)

	return isValid, nil
}

func getHash(password []byte, salt []byte) ([]byte, error) {
	return scrypt.Key(password, salt, 1<<14, 8, 1, PW_HASH_BYTES)
}

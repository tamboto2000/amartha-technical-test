package password

import (
	"bytes"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const DefaultCost = bcrypt.DefaultCost

func hashPasswordError(err error) error {
	return fmt.Errorf("error on hashing password: %v", err)
}

func HashPassword(pass string, saltLen int, cost int) ([]byte, []byte, error) {
	if cost <= 0 {
		cost = DefaultCost
	}

	salt, err := generateSalt(saltLen)
	if err != nil {
		return nil, nil, hashPasswordError(err)
	}

	passWithSalt := appendSaltToPass(pass, salt)

	bytes, err := bcrypt.GenerateFromPassword(passWithSalt, cost)
	if err != nil {
		return nil, nil, hashPasswordError(err)
	}

	return bytes, salt, nil
}

func CheckPasswordHash(pass string, salt, hash []byte) bool {
	passWithSalt := appendSaltToPass(pass, salt)
	err := bcrypt.CompareHashAndPassword(hash, passWithSalt)
	return err == nil
}

func generateSalt(n int) ([]byte, error) {
	salt := make([]byte, n)
	_, err := rand.Read(salt)
	return salt, err
}

func appendSaltToPass(pass string, salt []byte) []byte {
	return bytes.Join([][]byte{[]byte(pass), salt}, []byte(":"))
}

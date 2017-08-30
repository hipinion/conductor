package conductor

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func MakeSalt() string {
	n := 16
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func MakeSession() string {
	n := 64
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func ValidatePassword(password, testpassword string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(testpassword)); err != nil {
		return false, err
	}
	return true, nil
}

func Hash(password, salt string) (string, string, error) {

	if salt == "" {
		salt = MakeSalt()
	}

	inpass := password + salt

	hash, err := bcrypt.GenerateFromPassword([]byte(inpass), bcrypt.DefaultCost)

	if err != nil {
		return "", "", err
	}

	return string(hash), salt, nil
}

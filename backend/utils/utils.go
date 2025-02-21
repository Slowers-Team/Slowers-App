package utils

import (
	"encoding/json"
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"github.com/Slowers-team/Slowers-App/database"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func AreIDPtrSlicesEql(a []*database.ObjectID, b []*database.ObjectID) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range len(a) {
		if a[i].Hex() != b[i].Hex() {
			return false
		}
	}

	return true
}

func ToJSON(val any) []byte {
	asJSON, err := json.Marshal(val)
	if err != nil {
		log.Fatal(err)
	}
	return asJSON
}

package common

import (
	"github.com/alexandrevicenzi/unchained"
)

func HashPassword(password string) string {
	hash, err := unchained.MakePassword(password, unchained.GetRandomString(12), "default")
	if err != nil {
		return ""
	}
	return hash
}

func CheckPassword(password, storedPassword string) bool {
	valid, e := unchained.CheckPassword(password, storedPassword)
	if e != nil {
		return false
	}
	return valid
}

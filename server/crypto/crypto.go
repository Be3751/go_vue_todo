package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func EncPwd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash), err
}

func CompHashAndPwd(hash string, pwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err
}

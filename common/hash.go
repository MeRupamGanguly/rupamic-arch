package common

import "golang.org/x/crypto/bcrypt"

func Encrypt(s string) string {
	buf, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(buf)
}

func Decrypt(encPass string, pass string) (matched bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(encPass), []byte(pass))
	if err != nil {
		return false, err
	}
	return true, nil
}

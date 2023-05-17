package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Encryption 进行加密处理
func Encryption(s string) (str string, err error) {
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	str = string(hash)
	return
}

// EncryptionVerify 对比密码
func EncryptionVerify(pwd1 string, pwd2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwd2), []byte(pwd1))
	if err != nil {
		return true
	} else {
		return false
	}
}

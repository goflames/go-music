package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// 密码加密工具类
func HashPassword(password string) (string, error) {
	// GenerateFromPassword对密码进行加密，第一个参数是密码的byte数组
	//第二个参数是cost（值越高加密强度越高，默认使用bcrypt.DefaultCost）
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(hashedPassword, password string) bool {
	// CompareHashAndPassword对比密码和加密后的密码
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			log.Println("密码不匹配")
		} else {
			log.Println("密码验证错误:", err)
		}
		return false
	}
	return err == nil
}

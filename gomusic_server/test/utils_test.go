package test

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gomusic_server/utils"
	"log"
	"testing"
)

// TestHashPassword 测试 HashPassword 方法
func TestHashPassword(t *testing.T) {
	password := "testpassword"

	// 调用 HashPassword
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	log.Print(hashedPassword)

	// 检查返回的哈希密码是否符合 bcrypt 的格式
	cost, err := bcrypt.Cost([]byte(hashedPassword))
	if err != nil {
		t.Fatalf("hashed password does not have a valid bcrypt format: %v", err)
	}
	log.Print(cost)
}

// TestCheckPassword 成功的密码验证测试
func TestCheckPassword_Success(t *testing.T) {
	password := "testpassword"

	// 生成哈希密码
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	log.Print(hashedPassword)

	// 验证密码
	isValid := utils.CheckPassword(hashedPassword, password)
	assert.True(t, isValid, "password should be valid")
}

// TestCheckPassword_Failure 测试密码验证失败
func TestCheckPassword_Failure(t *testing.T) {
	password := "testpassword"
	wrongPassword := "wrongpassword"

	// 生成哈希密码
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// 验证错误的密码
	isValid := utils.CheckPassword(hashedPassword, wrongPassword)
	assert.False(t, isValid, "password should be invalid")
}

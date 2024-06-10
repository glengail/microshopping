package hash

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// 随机字符串
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// 使用当前时间戳创建seed
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// 创建salt
func CreateSalt() string {
	b := make([]byte, bcrypt.MaxCost)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// 使用bycrypt算法返回hash后密码
func HashPassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashpassword), err
}

//检查密码是否相等

func CheckPasswordHash(password, hashpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	return err == nil
}

package user

import (
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(30)"`
	Password  string `gorm:"type:varchar(100)"`
	Salt      string `gorm:"type:varchar(100)"`
	Token     string `gorm:"type:varchar(500)"`
	UserInfo  []byte `gorm:"type:bytes"`
	IsDeleted bool
	IsAdmin   bool
}

// 新建用户实例
func NewUser(username, password string, userInfo []byte) *User {

	return &User{
		Username:  username,
		Password:  password,
		UserInfo:  userInfo,
		IsDeleted: false,
		IsAdmin:   false,
	}
}

package user

import (
	hash "userservice/utils/hash"

	"gorm.io/gorm"
)

func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Salt == "" {
		salt := hash.CreateSalt()
		hashpassword, err := hash.HashPassword(u.Password + salt)
		if err != nil {
			return err
		}
		u.Password = hashpassword
		u.Salt = salt
	}
	return nil
}

package mysql

import (
	"errors"
	"github.com/Nero011/cloudStorage/server/shared/model/user"
	"gorm.io/gorm"
)

// check user exist or not, if exist return true
func CheckUser(u *user.User) bool {
	db := GetMysql()
	err := db.First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func CreateUser(u *user.User) error {
	if CheckUser(u) {
		return errors.New("user is exist")
	}
	db := GetMysql()
	err := db.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

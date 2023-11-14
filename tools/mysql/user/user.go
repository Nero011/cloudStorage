package user

import (
	"errors"
	"github.com/Nero011/cloudStorage/server/shared/model/user"
	"github.com/Nero011/cloudStorage/tools/mysql"
	"gorm.io/gorm"
)

// check user exist or not, if exist return true
func checkUser(u user.User) bool {
	db := mysql.GetMysql()
	err := db.First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

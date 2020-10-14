package users

import (
	"framework_seed/pkg/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID 			int64 		`gorm:"primaryKey"`
	UserName 	string      `gorm:"size:64"`
	Password 	string		`gorm:"size:64"`
	Email 		string		`gorm:"size:64"`
	Gender 		bool
	CreateTime 	time.Time
	UpdateTime 	time.Time
}

func ExistUserByUserName(username string) (bool2 bool, err error) {
	var user User
	result := mysql.DB.Debug().Find(&user, "username = ?", username)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}

	if result.RowsAffected > 0 {
		return true, nil
	}

	return false, nil
}
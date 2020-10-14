package mysql

import (
	"fmt"
	"framework_seed/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(cfg *settings.MySQLConfig) (err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	if DB.Error != nil {
		return DB.Error
	}

	return nil
}
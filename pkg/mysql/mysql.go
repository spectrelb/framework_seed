package mysql

import (
	"fmt"
	"framework_seed/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Init(cfg *settings.MySQLConfig) error {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	if db.Error != nil {
		return db.Error
	}
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key
	db.Delete(&product, 1)
	return nil
}

//type User struct {
//
//}
//
//func main()  {
//	Init(settings.Conf.MySQLConfig)
//	db.Select()
//}
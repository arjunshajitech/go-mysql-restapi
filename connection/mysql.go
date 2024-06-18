package connection

import "gorm.io/driver/mysql"
import "gorm.io/gorm"
import "fmt"
import "go/packages/models"

var DB *gorm.DB
var err error

const DNS = "root:root@tcp(127.0.0.1:5005)/db?charset=utf8&parseTime=True&loc=Local"

func InitlizeMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB.")
	}
	fmt.Println("Connected to mysql.")
	DB.AutoMigrate(&models.User{})
}

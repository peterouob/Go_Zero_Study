package init_gorm

import (
	"fmt"
	"go_zero/study_model/user_gorm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm(MysqlDatabase string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(MysqlDatabase))
	if err != nil {
		panic("connect to database error")
	}
	fmt.Println("success to connect")
	db.AutoMigrate(&models.UserModel{})
	return db
}

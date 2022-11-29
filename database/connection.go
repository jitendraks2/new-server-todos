package database

import (
	"github.com/jitendraks2/todo-crud-golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "Jitendra:jitendra18@tcp(127.0.0.1:3306)/todocrud?charset=utf8mb4&parseTime=True&loc=Local"// dsn := "dw4pbngraxqjfji2n1cv:pscale_pw_wSTtFQmC00SDpug2bVz5CH4vXOB6VHdmimbKm5I5uOm@tcp(ap-south.connect.psdb.cloud)/todolearn?tls=true"


	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to db")
	}

	DB = connection

	connection.AutoMigrate(&models.Todos{})
}

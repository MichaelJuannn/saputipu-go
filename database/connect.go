package database

import (
	"fiber-saputipu/config"
	models "fiber-saputipu/internals/model"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Println("db something wrong in port ", port)
	}

	// connection url
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
	// connect to db and create connection var
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("connection to db failed")
	}
	DB.AutoMigrate(&models.User{}, &models.Rekening{}, &models.PredictionText{}, &models.Laporan{})

	fmt.Println("connection to db established")

}

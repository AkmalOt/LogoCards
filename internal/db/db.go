package db

import (
	logging "LogoForCardsGin/logger"
	"LogoForCardsGin/models"
	"encoding/json"
	"fmt"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"os"
)

var DataB *gorm.DB

func GetDbConnection() *gorm.DB {
	logger := logging.GetLogger()

	file, err := os.Open("./internal/db/db.json")
	if err != nil {
		logger.Println(err)
		return nil
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		logger.Println(err)
		return nil
	}

	//var DB models.DbData
	var DB models.DbData

	err = json.Unmarshal(bytes, &DB)
	if err != nil {
		logger.Println(err)
		return nil
	}

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dushanbe",
		DB.DbConnection.Host, DB.DbConnection.User, DB.DbConnection.Password, DB.DbConnection.Dbname, DB.DbConnection.Port)

	conn, err := gorm.Open(postgresDriver.Open(connString))
	if err != nil {
		logger.Println(err, "Не удалось подключиться к базе данных")
		return nil
	}

	//err = conn.AutoMigrate(&UserCards{})
	logger.Println("Success connection to", DB.DbConnection.Host)
	DataB = conn

	return conn
	//return DataB, nil

}

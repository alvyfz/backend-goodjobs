package mysql

import (
	"fmt"
	"log"
)

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func (config *ConfigDB) InitialDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Database)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB.Logger = logger.Default.LogMode(logger.Info)
	if err != nil {
		log.Fatal(err)
	}
	return DB
}
package storage

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func connect(userPass, schema, hostAndPort string) (*gorm.DB, error) {
	connURL := fmt.Sprintf("%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10s", userPass, hostAndPort, schema)
	return gorm.Open(mysql.Open(connURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

// Connect to database
func ConnectMysql(userPass, schema, host string) (*gorm.DB, error) {
	dbConn, err := connect(userPass, schema, host)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

// Close -
func Close(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

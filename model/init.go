package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitSqlite(dsn string) error {
	var err error
	Db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	Db = Db.Debug()
	err = Db.AutoMigrate(&Word{})
	err = Db.AutoMigrate(&Pending{})
	return err
}

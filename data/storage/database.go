package storage

import (
	"citadel-api/data/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var once sync.Once
var db *gorm.DB

func GetDB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic("failed to connect database")
		}

		//Migrate the schema
		err = db.AutoMigrate(
			&model.Block{},
		)
	})
	return db
}

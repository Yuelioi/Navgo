package db

import (
	"backend/internal/global"
	"backend/internal/types"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	if _, err = os.Stat(global.ConfInst.Database.Resource); os.IsNotExist(err) {
		dir := filepath.Dir(global.ConfInst.Database.Resource)

		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("failed to create directory: %v", err)
		}

		file, err := os.Create(global.ConfInst.Database.Resource)
		if err != nil {
			panic("failed to create file")
		}
		defer file.Close()
	}

	DB, err = gorm.Open(sqlite.Open(global.ConfInst.Database.Resource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = DB.AutoMigrate(&types.DBCollection{}, &types.DBComment{})
}

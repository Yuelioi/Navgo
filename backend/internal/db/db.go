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

	err = DB.AutoMigrate(&types.Category{}, &types.Collection{}, &types.Comment{}, &types.Announce{})

	managers := make([]Manager, 0)
	managers = append(
		managers,
		NewCollectionManager(global.ConfInst.Resource.Collections),
		NewAnnounceManager(
			filepath.Join(global.ConfInst.Resource.Announces, global.ConfInst.Resource.MetaFile),
		),
		NewCommentManager(
			filepath.Join(global.ConfInst.Resource.Comments, global.ConfInst.Resource.MetaFile),
		),
	)

	for _, manager := range managers {
		manager.Read()
		manager.Build()
		manager.Reduce()
	}

}

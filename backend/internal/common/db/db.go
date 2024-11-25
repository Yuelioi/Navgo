package db

import (
	"backend/internal/common/constants"
	_ "backend/internal/common/startup"
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
	if _, err = os.Stat(constants.ConfInst.Database.Resource); os.IsNotExist(err) {
		dir := filepath.Dir(constants.ConfInst.Database.Resource)

		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("failed to create directory: %v", err)
		}

		file, err := os.Create(constants.ConfInst.Database.Resource)
		if err != nil {
			panic("failed to create file")
		}
		defer file.Close()
	}

	DB, err = gorm.Open(sqlite.Open(constants.ConfInst.Database.Resource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = DB.AutoMigrate(&types.Category{}, &types.Collection{}, &types.Comment{}, &types.Announce{}, &types.Statistics{}, &types.User{})

	managers := make([]Manager, 0)
	managers = append(
		managers,
		NewCollectionManager(constants.ConfInst.Resource.Collections),
		NewAnnounceManager(
			filepath.Join(constants.ConfInst.Resource.Announces, constants.ConfInst.Resource.MetaFile),
		),
		NewCommentManager(
			filepath.Join(constants.ConfInst.Resource.Comments, constants.ConfInst.Resource.MetaFile),
		),
	)

	for _, manager := range managers {
		manager.Read()
		manager.Build()
		manager.Reduce()
	}

}

package app

import (
	"github.com/pedroelsner/colly-crawler/internal/entity"
	"github.com/pedroelsner/colly-crawler/internal/provider/sqlite"
)

func migration() {
	db := sqlite.Init()

	db.AutoMigrate(&entity.Asset{})
}

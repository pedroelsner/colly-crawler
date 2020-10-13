package asset

import (
	"sync"

	"github.com/jinzhu/gorm"

	"github.com/pedroelsner/colly-crawler/internal/entity"
	"github.com/pedroelsner/colly-crawler/internal/provider/sqlite"
)

// SqliteAssetRepository
type SqliteAssetRepository struct {
	DB *gorm.DB
}

// FindByURL
func (r *SqliteAssetRepository) FindByURL(url string) (*entity.Asset, error) {
	record := &entity.Asset{}
	err := r.DB.
		Where("url = ?", url).
		First(record).
		Error

	return record, err
}

// Save
func (r *SqliteAssetRepository) Save(record *entity.Asset) error {

	// Create
	if record.ID == "" {
		return r.DB.Create(record).Error
	}

	// Update
	return r.DB.Save(record).Error
}

// Test interface
var _ AssetRepositoryIface = (*SqliteAssetRepository)(nil)

// Singleton
var (
	onceSqlite       sync.Once
	sqliteRepository *SqliteAssetRepository
)

func InitSqliteRepository() AssetRepositoryIface {
	onceSqlite.Do(func() {
		db := sqlite.Init()
		sqliteRepository = &SqliteAssetRepository{DB: db}
	})

	return sqliteRepository
}

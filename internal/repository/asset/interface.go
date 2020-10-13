package asset

import (
	"github.com/pedroelsner/colly-crawler/internal/entity"
)

// AssetRepositoryIface
type AssetRepositoryIface interface {
	FindByURL(string) (*entity.Asset, error)
	Save(*entity.Asset) error
}

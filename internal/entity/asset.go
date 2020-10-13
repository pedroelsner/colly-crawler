package entity

import (
	"gorm.io/datatypes"
)

type Asset struct {
	Base

	URL          string `gorm:"unique"`
	Source       string
	Type         string
	Description  string
	State        string
	City         string
	Neighborhood string
	Address      string
	Owner        string
	LastBid      float64
	IncrementBid float64
	FirstDate    string
	FirstPrice   float64
	SecondDate   string
	SecondPrice  float64
	Tags         datatypes.JSON
	Documents    datatypes.JSON
}

func (Asset) TableName() string {
	return "assets"
}

type AssetList []*Asset

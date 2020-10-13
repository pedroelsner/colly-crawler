package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        string `sql:"type:char(36)" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4().String()
	return scope.SetColumn("ID", uuid)
}

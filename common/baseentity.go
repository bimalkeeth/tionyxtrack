package common

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"time"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	guid := uuid.New()
	scope.SetColumn("ID", guid)
	return nil
}

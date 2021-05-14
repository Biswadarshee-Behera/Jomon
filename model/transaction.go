package model

import (
	"time"

	"github.com/google/uuid"
)

type TransactionRepository interface {
}

type Transaction struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time `gorm:"index"`
}
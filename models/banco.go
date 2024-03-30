package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Banco struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Codigo     string  `gorm:"unique"`
	Nombre       string 
	Cuenta     string
	

}

func (Banco) TableName() string {
	return "banco"
}

func (banco *Banco) BeforeCreate(tx *gorm.DB) (err error) {
	banco.ID = uuid.New()
	now := time.Now().UTC().Format(time.RFC3339)
	banco.CreatedAt, err = time.Parse(time.RFC3339, now)
	if err != nil {
		return err
	}
	banco.UpdatedAt = banco.CreatedAt
	return nil
}
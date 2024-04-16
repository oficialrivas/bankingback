package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Cedula     string  
	Rif       string `gorm:"unique"`
	Numero     string
	Banco      string
	Nombre     string
	Saldo     float32
	PIN       string

}

func (User) TableName() string {
	return "user"
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	now := time.Now().UTC().Format(time.RFC3339)
	user.CreatedAt, err = time.Parse(time.RFC3339, now)
	if err != nil {
		return err
	}
	user.UpdatedAt = user.CreatedAt
	return nil
}

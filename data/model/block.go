package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Block struct {
	Id         string         `gorm:"primarykey,type:uuid" json:"id"`
	Type       string         `json:"type"`
	Content    *string        `json:"content,omitempty"`
	Children   []Block        `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Properties map[string]any `json:"properties,omitempty" gorm:"serializer:json"`
	Parent     *Block         `json:"-"`
	ParentID   *string        `json:"-"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (block *Block) BeforeCreate(tx *gorm.DB) (err error) {
	block.Id = uuid.NewString()
	if block.Properties == nil {
		block.Properties = make(map[string]any)
	}
	return
}

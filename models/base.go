package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uint  `json:"-"`
	CreatedAt int64 `json:"created_at" description:"创建时间"`
	UpdatedAt int64 `json:"updated_at" description:"更新时间"`
}

type UUIDBase struct {
	Base
	UUID string `json:"uuid" gorm:"type:varchar(36);unique;not null" description:"uuid"`
}

type PageInput struct {
	PageID   int `query:"page_id" description:"查询起始页id" default:"0"`
	PageSize int `query:"page_size" description:"单页查询数量,默认十条" default:"10"`
	Order    string
}

func (b *UUIDBase) BeforeCreate(tx *gorm.DB) (err error) {
	b.UUID = uuid.New().String()
	return
}

func Count(tx *gorm.DB, query interface{}, args ...interface{}) int64 {
	var count int64
	if query != nil {
		tx.Where(query, args...).Count(&count)
	} else {
		tx.Count(&count)
	}
	return count
}

func Page(tx *gorm.DB, in PageInput, dest, query interface{}, args ...interface{}) error {
	tx = tx.Offset(in.PageID * in.PageSize).Limit(in.PageSize)
	if query != nil {
		tx = tx.Where(query, args...)
	}
	if in.Order != "" {
		tx = tx.Order(in.Order)
	}
	return tx.Find(dest).Error
}

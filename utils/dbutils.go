package utils

import "github.com/jinzhu/gorm"

// Conn is
type Conn interface {
	First(out interface{}, where ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Model(value interface{}) *gorm.DB
	Update(attrs ...interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
}

package database

import "gorm.io/gorm"

type SqlHandler interface {
	Create(object interface{})
	FindAll(object interface{})
	FindById(object interface{}, id int)
	DeleteById(object interface{}, id int) error
	Raw() *gorm.DB
}

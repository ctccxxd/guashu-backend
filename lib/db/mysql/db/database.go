package db

import (
	"gorm.io/gorm"
	"jihulab.com/guashu/gs-server/lib/db/mysql"
)

func Connect(name string) *gorm.DB {
	return mysql.Session(name)
}

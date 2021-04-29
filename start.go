package gorm_utilities

import "github.com/jinzhu/gorm"

type Utility struct {
	dbMain  *gorm.DB
	dbSlave *gorm.DB
}

func NewUtility(dbMain, dbSlave *gorm.DB) *Utility {
	return &Utility{dbMain: dbMain, dbSlave: dbSlave}
}

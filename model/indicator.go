package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Indicator struct {
	gorm.Model
	Light       int     `json:"lightintensity"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}

func IndicatorDBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Indicator{})
	return db
}

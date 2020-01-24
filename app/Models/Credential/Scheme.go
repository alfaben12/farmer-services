package FarmLandModel

import (
	"github.com/jinzhu/gorm"
)

type FarmLand struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(100);"`
	Image        string  `gorm:"type:varchar(255);"`
}

func (l *FarmLand) TableName() string {
	return "farm_land"
}



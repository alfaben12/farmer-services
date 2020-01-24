package FarmLandModel

import (
	"github.com/jinzhu/gorm"
)

type FarmLand struct {
	gorm.Model
	Accountid    string  `gorm:"type:int(11);"`
	Name         string  `gorm:"type:varchar(100);"`
	Image        string  `gorm:"type:varchar(255);"`
}

func (l *FarmLand) TableName() string {
	return "farm_land"
}



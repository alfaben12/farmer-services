package FarmLandModel

import (
	"../../Config"
	_ "github.com/go-sql-driver/mysql"
)

func Get(l *[]FarmLand) (err error) {
	if err = Config.DB.Find(l).Error; err != nil {
		return err
	}
	return nil
}

func Create(l *FarmLand) (err error) {
	if err = Config.DB.Create(l).Error; err != nil {
		return err
	}
	return nil
}

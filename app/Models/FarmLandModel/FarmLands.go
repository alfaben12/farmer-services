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
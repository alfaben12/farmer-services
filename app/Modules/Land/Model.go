package Land

import (
	"../../Config"
	_ "github.com/go-sql-driver/mysql"
)

func Get(l *[]Land) (err error) {
	if err = Config.DB.Find(l).Error; err != nil {
		return err
	}
	return nil
}

func CreateData(l *Land) (err error) {
	if err = Config.DB.Create(l).Error; err != nil {
		return err
	}
	return nil
}

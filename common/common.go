package common

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"webDemo/common/db"
)

var (
	DB    *gorm.DB
	Cache *db.Redis
)

func Init() error {
	var err error
	DB, err = db.NewMysql("config.Conf.Mysql")
	if err != nil {
		return errors.New(fmt.Sprintf("initialize db error: %s", err.Error()))
	}

	Cache, err = db.NewRedis("config.Conf.Redis")
	if err != nil {
		return errors.New(fmt.Sprintf("initialize cache error: %s", err.Error()))
	}

	return nil
}

package db

import (
	"github.com/isd-sgcu/rpkm67-gateway/config"
	"github.com/isd-sgcu/rpkm67-model/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func InitDatabase(conf *config.DbConfig, isDebug bool) (db *gorm.DB, err error) {
	gormConf := &gorm.Config{TranslateError: true}

	if !isDebug {
		gormConf.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
	}

	db, err = gorm.Open(postgres.Open(conf.Url), gormConf)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Group{}, &model.User{}, &model.Selection{}, &model.Stamp{}, &model.CheckIn{})
	if err != nil {
		return nil, err
	}

	return
}

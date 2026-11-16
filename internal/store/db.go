package store

import (
	"go-backend-template/internal/conf"
	"go-backend-template/internal/model"
	"net/url"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() error {
	var err error
	gormConfig := gorm.Config{Logger: logger.Discard}
	if conf.IsDebug() {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	db, err = initSQLite(conf.AppConfig.Database.Path, &gormConfig)
	if err != nil {
		return err
	}
	return db.AutoMigrate(new(model.User))
}

func Close() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func initSQLite(path string, config *gorm.Config) (*gorm.DB, error) {
	params := url.Values{}
	params.Add("_pragma", "journal_mode(WAL)")
	params.Add("_pragma", "synchronous(NORMAL)")
	params.Add("_pragma", "cache_size(10000)")
	params.Add("_pragma", "busy_timeout(5000)")
	params.Add("_pragma", "foreign_keys(ON)")
	params.Add("_pragma", "auto_vacuum(INCREMENTAL)")
	params.Add("_pragma", "mmap_size(268435456)")
	params.Add("_pragma", "locking_mode(NORMAL)")
	return gorm.Open(sqlite.Open(path+"?"+params.Encode()), config)
}

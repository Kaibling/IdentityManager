package database

import (
	"github.com/Kaibling/IdentityManager/lib/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDBConnection() (*gorm.DB, error) {
	if config.Configuration.Dialect == "SQLITE" {
		db, err := gorm.Open(sqlite.Open(config.Configuration.DBFilePath), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	}
	return nil, nil

}

func Migrate(db *gorm.DB, dbMigs []DBMigrator) error {
	for i := range dbMigs {
		err := dbMigs[i].Migrate()
		if err != nil {
			return err
		}
	}
	return nil
}

type DBMigrator interface {
	Migrate() error
}

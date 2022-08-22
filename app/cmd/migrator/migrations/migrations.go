package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var migrations = make([]*gormigrate.Migration, 0)

func Add(id string, migrate gormigrate.MigrateFunc, rollback gormigrate.RollbackFunc) {
	if migrate == nil {
		migrate = func(db *gorm.DB) error { return nil }
	}
	if rollback == nil {
		rollback = func(db *gorm.DB) error { return nil }
	}
	migrations = append(migrations, &gormigrate.Migration{
		ID:       id,
		Migrate:  migrate,
		Rollback: rollback,
	})
}

func Up(db *gorm.DB) (err error) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err = m.Migrate(); err != nil {
		log.Errorf("Could not migrate: %v", err)
		return
	}
	log.Infof("Run migrations successfully")
	return nil
}

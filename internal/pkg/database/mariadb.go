package database

import (
	"cpds/cpds-analyzer/internal/models/rules"

	"gorm.io/gorm"
)

type mariadb struct {
	db *gorm.DB
}

func New(db *gorm.DB) Database {
	return &mariadb{
		db: db,
	}
}

func (m *mariadb) Init() error {
	if err := m.db.AutoMigrate(&rules.Rule{}); err != nil {
		return err
	}

	return nil
}

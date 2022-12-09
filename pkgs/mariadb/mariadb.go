package mariadb

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mariadb struct {
	DatabaseAddress  string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
}

func (m *Mariadb) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/gorm?charset=utf8&parseTime=True&loc=Local",
		m.DatabaseUser,
		m.DatabasePassword,
		m.DatabaseAddress,
		m.DatabasePort,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

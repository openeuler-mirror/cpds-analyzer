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
		DSN:                       dsn,
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported  MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

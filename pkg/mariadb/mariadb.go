package mariadb

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MariaDB struct {
	Host        string
	Port        int
	Username    string
	Password    string
	MaxOpenConn int
	MaxIdleConn int
	MaxLifetime time.Duration
}

func (d *MariaDB) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%v)/cpds?charset=utf8mb4&parseTime=True&loc=Local`,
		d.Username,
		d.Password,
		d.Host,
		d.Port,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "cpds",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(d.MaxOpenConn)
	sqlDB.SetMaxIdleConns(d.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(d.MaxLifetime)

	// db.Use(&TracePlugin{})

	return db, nil
}

package utils

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Mariadb struct {
	conf *Config
	db   *gorm.DB
}

type Config struct {
	DatabaseAddress  string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

const (
	defaultDatabaseName = "cpds"
)

func NewDB(dbAddr string, dbPort int, dbUser string, dbPasswd string) *Mariadb {
	logrus.Debugf("Database Address: %s", dbAddr)
	logrus.Debugf("Database Port: %d", dbPort)
	logrus.Debugf("Database User: %s", dbUser)
	logrus.Debugf("Database password: %s", dbPasswd)
	logrus.Debugf("Database name: %s", defaultDatabaseName)
	c := &Config{
		DatabaseAddress:  dbAddr,
		DatabasePort:     strconv.Itoa(dbPort),
		DatabaseUser:     dbUser,
		DatabasePassword: dbPasswd,
		DatabaseName:     defaultDatabaseName,
	}
	return &Mariadb{
		conf: c,
	}
}

func (m *Mariadb) Connect() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		m.conf.DatabaseUser,
		m.conf.DatabasePassword,
		m.conf.DatabaseAddress,
		m.conf.DatabasePort,
		m.conf.DatabaseName,
	)
	logrus.Debugf("connect database with database user: %s", m.conf.DatabaseUser)
	logrus.Debugf("connect database with database password: %s", m.conf.DatabasePassword)
	logrus.Debugf("connect database with database address: %s", m.conf.DatabaseAddress)
	logrus.Debugf("connect database with database port: %s", m.conf.DatabasePort)
	logrus.Debugf("connect database with database name: %s", m.conf.DatabaseName)

	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported  MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}

	gormConfig := gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	}

	d, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig)
	if err != nil {
		return err
	}
	*m = Mariadb{
		db: d,
	}
	return nil
}

func (m *Mariadb) InitDBTables(db *gorm.DB, table interface{}) error {
	logrus.Infof("initialize database tables")
	if err := db.AutoMigrate(&table); err != nil {
		return err
	}
	return nil
}

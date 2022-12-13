package db

import (
	"database/sql"

	entSql "entgo.io/ent/dialect/sql"

	"github.com/sirupsen/logrus"

	// Need to work with migration files.
	_ "github.com/lib/pq"
)

func GetDriver(conf *Config, log *logrus.Logger) (*entSql.Driver, error) {
	db, err := GetConnect(conf, log)
	if err != nil {
		return nil, err
	}
	return entSql.OpenDB(conf.DriverName, db), nil
}

func Migrations(conf *Config, log *logrus.Logger) error {
	db, err := sql.Open(conf.DriverName, conf.DSN())
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}

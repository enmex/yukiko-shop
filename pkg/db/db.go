package db

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	DefaultHost               = "127.0.0.1"
	DefaultPort               = "5432"
	DefaultDatabaseName       = "versioning_pg"
	DefaultUser               = "postgres"
	DefaultSSLMode            = "disable"
	DefaultMigrationDirectory = "migrations"
	DefaultDriverName         = "postgres"
)

type Config struct {
	Host               string
	Port               string
	DatabaseName       string
	User               string
	Password           string
	Schema             string
	MigrationDirectory string
	SslMode            string
	DriverName         string
}

func (db *Config) String() string {
	return fmt.Sprintf("Connnecting to DB on %s:%s/%s as '%s' ...", db.Host, db.Port, db.DatabaseName, db.User)
}

// DSN - Data Source Name or connection string
func (db *Config) DSN() string {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s sslmode=%s user=%s password=%s",
		db.Host, db.Port, db.DatabaseName, db.SslMode, db.User, db.Password,
	)
	if db.Schema != "" {
		dsn += fmt.Sprintf(" search_path=%s", db.Schema)
	}
	return dsn
}

// DefaultDSN - Data Source Name or connection string with dbname="postgres"
func (db *Config) DefaultDSN() string {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s sslmode=%s user=%s password=%s",
		db.Host, db.Port, "sprint_pulse_test", db.SslMode, db.User, db.Password,
	)
	if db.Schema != "" {
		dsn += fmt.Sprintf(" search_path=%s", db.Schema)
	}
	return dsn
}

func GetConnect(cfg *Config, log *logrus.Logger) (*sql.DB, error) {
	log.Infof("Connecting to DB on %s:%s/%s as '%s' ... ", cfg.Host, cfg.Port, cfg.DatabaseName, cfg.User)
	db, err := sql.Open(cfg.DriverName, cfg.DSN())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return db, nil
}

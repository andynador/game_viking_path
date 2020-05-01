package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	DBDriverName            = "postgres"
	ErrDBConnAttemptsFailed = errors.New("All attempts failed")
)

type Config struct {
	MaxConnLifetimeSec int
	MaxIdleConns       int
	MaxOpenConns       int
}

type Database struct {
	db *sql.DB
}

func New(dsn string, cfg Config) (*Database, error) {
	db, err := sql.Open(DBDriverName, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "Can't open database")
	}
	db.SetConnMaxLifetime(time.Second * time.Duration(cfg.MaxConnLifetimeSec))
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	return &Database{
		db: db,
	}, nil
}

func (d *Database) GetConnection() *sql.DB {
	return d.db
}

func (d *Database) Connect() error {
	var dbError error

	maxAttempts := 5
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		dbError = d.db.Ping()
		if dbError == nil {
			break
		}
		nextAttemptWait := time.Duration(attempt) * time.Second

		time.Sleep(nextAttemptWait)
	}

	if dbError != nil {
		return ErrDBConnAttemptsFailed
	}
	return nil
}

func (d *Database) Close() error {
	err := d.db.Close()
	if err != nil {
		return errors.Wrap(err, "Can't close database")
	}
	return nil
}

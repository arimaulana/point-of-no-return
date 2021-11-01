package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/arimaulana/point-of-no-return/internal/common/pkg/log"
	"github.com/arimaulana/point-of-no-return/internal/sample/configs"
	migrater "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

// Migrate to run database migration up or down
func MigrateSql(cfg configs.Config, logger log.Logger, command string) error {
	dsn := generateDSN(cfg)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Error(err)
		return err
	}
	defer db.Close()

	path, err := os.Getwd()
	if err != nil {
		logger.Error(err)
		return err
	}

	migrationPath := fmt.Sprintf("file://%s/migrations", path)
	logger.Infof("migrationPath : %s", migrationPath)

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		logger.Error(err)
		return err
	}
	m, err := migrater.NewWithDatabaseInstance(
		migrationPath,
		"mysql",
		driver,
	)
	if err != nil {
		logger.Error(err)
		return err
	}
	if command == "up" {
		logger.Info("Migrate up")
		if err := m.Up(); err != nil && err != migrater.ErrNoChange {
			logger.Errorf("An error occurred while syncing the database.. %v", err)
			return err
		}
	}

	if command == "down" {
		logger.Info("Migrate down")
		if err := m.Down(); err != nil && err != migrater.ErrNoChange {
			logger.Errorf("An error occurred while syncing the database.. %v", err)
			return err
		}
	}

	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Migrate complete")
	return nil
}

// Seed to populate database with seed data
func SeedSql(cfg configs.Config, logger log.Logger, sqlFilename string) error {
	dsn := generateDSN(cfg)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	// load from SQL file
	bytes, err := ioutil.ReadFile(sqlFilename)
	if err != nil {
		return err
	}
	sql := string(bytes)
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(sql); err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Info("Seed data complete")
	return nil
}

func generateDSN(cfg configs.Config) string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&multiStatements=true",
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)
}

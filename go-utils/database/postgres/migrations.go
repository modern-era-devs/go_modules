package postgres

// TODO: Use go-utils/logger instead of logger
import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"log"
	"os"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	appMigrationFilesPath string
	appMigrate            *migrate.Migrate
)

const (
	migrationFilePath = "./migrations"
)

// create new migration files
func CreateMigrationFiles(filename string, cfg PostgresConfig) error {

	err := InitPostgres(cfg, migrationFilePath)
	if err != nil {
		return err
	}
	return Create(filename)
}

// run migrations present in ./migrations
func RunDatabaseMigrations(cfg PostgresConfig) error {
	err := InitPostgres(cfg, migrationFilePath)
	if err != nil {
		return err
	}
	return Run()
}

// rollback last migration present in ./migrations
func RollbackLatestMigration(cfg PostgresConfig) error {
	err := InitPostgres(cfg, migrationFilePath)
	if err != nil {
		return err
	}
	return RollbackLatest()
}

// initiating migrate
func InitPostgres(dbConfig PostgresConfig, migrationFilesPath string) error {
	appMigrationFilesPath = migrationFilesPath

	var err error
	appMigrate, err = migrate.New("file://"+migrationFilesPath, dbConfig.GetConnectionURL())
	if err != nil {
		return errors.Wrap(err, "failed to init migration")
	}
	return nil
}

func Create(filename string) error {
	if len(filename) == 0 {
		return errors.Wrap(nil, "Migration filename is not provided")
	}

	timeStamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", appMigrationFilesPath, timeStamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", appMigrationFilesPath, timeStamp, filename)

	if err := createFile(upMigrationFilePath); err != nil {
		return errors.Wrap(err, "up migration file creation failed")
	}
	log.Println("up migration file created successfully")

	if err := createFile(downMigrationFilePath); err != nil {
		os.Remove(upMigrationFilePath)
		return errors.Wrap(err, "down migration file creation failed")
	}

	log.Println("down migration file created successfully")

	return nil
}

func Run() error {
	err := appMigrate.Up()
	if err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "applying migrations failed")
	}

	log.Println("migrations applied successfully")
	return nil
}

func RollbackLatest() error {
	err := appMigrate.Steps(-1)
	if err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "rollback failed")
	}

	log.Println("rollback successful")
	return nil
}

func createFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

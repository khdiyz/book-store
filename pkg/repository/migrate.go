package repository

import (
	"database/sql"
	"strconv"

	"github.com/khdiyz/book-store/config"
	"github.com/khdiyz/book-store/utils/logger"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateSQLite(cfg *config.Config, log *logger.Logger, up bool) error {
	// Read migrations from a folder
	migrations := &migrate.FileMigrationSource{
		Dir: "./schema",
	}
	db, err := sql.Open("sqlite3", cfg.SQLiteFilePath)
	if err != nil {
		log.Error("error in creating migrations: " + err.Error())
		return err
	}
	defer db.Close()

	migrateState := migrate.Up
	if !up {
		migrateState = migrate.Down
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrateState)
	if err != nil {
		log.Error("error in creating migrations: " + err.Error())
	}

	if n > 0 {
		log.Info("migrations applied: " + strconv.Itoa(n))
	}

	return nil
}

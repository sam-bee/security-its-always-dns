package persistence

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	dbHandle *sql.DB
}

func GetDb(pathToSqlite string) *Database {
	dbHandle, err := open(pathToSqlite)
	if err != nil {
		panic(err)
	}

	db := Database{dbHandle: dbHandle}

	return &db
}

func open(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	return db, err
}

func (db *Database) Initialise() error {
	_, err := db.dbHandle.Exec(
		"CREATE TABLE IF NOT EXISTS `dns_lookups_received` " +
			" (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `fqdn` TEXT, `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP)",
	)
	return err
}

func (db *Database) Store(fqdn string) error {
	_, err := db.dbHandle.Exec("INSERT INTO `dns_lookups_received` (`fqdn`) VALUES (?)", fqdn)
	return err
}

func (db *Database) Close() {
	db.dbHandle.Close()
}

func (db *Database) GetAllFqdns() ([]string, error) {
	rows, err := db.dbHandle.Query("SELECT `fqdn` FROM `dns_lookups_received`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fqdns []string
	for rows.Next() {
		var fqdn string
		err = rows.Scan(&fqdn)
		if err != nil {
			return nil, err
		}
		fqdns = append(fqdns, fqdn)
	}

	return fqdns, nil
}

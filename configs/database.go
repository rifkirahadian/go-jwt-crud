package configs

import (
	"database/sql"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func Migrate(db *sql.DB) {
	sql := `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100),
			email VARCHAR(100) UNIQUE,
			password VARCHAR(100)
		);
	`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
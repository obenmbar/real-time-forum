package forum

import (
	"database/sql"
	"fmt"
	"os"

 _  "modernc.org/sqlite"
)

func InitialeDb() (*sql.DB, error) {
	err := os.MkdirAll("db", 0o777)
	if err != nil {
		return nil, fmt.Errorf("errror en creation db dir")
	}

	db, err := sql.Open("sqlite", "db/forum.db")
	if err != nil {
		return nil, fmt.Errorf("error en create forum.db")

	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, fmt.Errorf("error enabling foreign keys: %v", err)
	}

	data, err := os.ReadFile("db/forum.sql")
	if err != nil {
		return nil, fmt.Errorf("error en read formsql")
	}

	commandsql := string(data)
	_, err = db.Exec(commandsql)
	if err != nil {
		return nil, fmt.Errorf("error en execution db")
	}

	return db, nil

}

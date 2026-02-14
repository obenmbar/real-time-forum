package forum

import (
	"database/sql"
	"fmt"
	"os"
)

func InitialeDb() (*sql.DB, error) {
	err := os.MkdirAll("db", 0o777)
	if err != nil {
		return nil, fmt.Errorf("errror en creation db dir")
	}

	db, err := sql.Open("sqlite3", "db/forumm.db")
	if err != nil {
		return nil,fmt.Errorf("error en create forum.db")
		
	}

	if err = db.Ping(); err != nil {
		return nil,  fmt.Errorf("errror en connecting avec database ")
	}
   
	data, err := os.ReadFile("db/forum.sql")
	if err != nil {
		return nil,  fmt.Errorf("error en read formsql")
	}

	commandsql :=  string(data)
    _, err = db.Exec(commandsql)
	if err != nil {
		return nil, fmt.Errorf("error en execution db")
	}

 return db ,nil
	
}

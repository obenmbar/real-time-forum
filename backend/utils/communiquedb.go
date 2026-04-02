package forum

import (
	"database/sql"
	"fmt"
	"time"
)

func InsertUser(db *sql.DB, data Users, Id string, hashpassword string) error {
	query := "INSERT INTO users (id,nickname,age,gender,first_name,last_name,email,password)VALUES(?,?,?,?,?,?,?,?)"

	_, err := db.Exec(query, Id, data.Nickname, data.Age, data.Gender, data.FirstName, data.LastName, data.Email, data.Password)
	if err != nil {
		return err
	}

	return nil
}
func CreateSession(db *sql.DB, userid string) (string, error) { 
    queryDelete := "DELETE FROM sessions WHERE user_id = ?"
    queryInsert := "INSERT INTO sessions (id, user_id, expires_at) VALUES (?, ?, ?)"
    
    tx, err := db.Begin()
    if err != nil {
        return "", fmt.Errorf("failed to begin transaction: %v", err)
    }

    
    _, err = tx.Exec(queryDelete, userid)
    if err != nil {
        tx.Rollback()
        return "", fmt.Errorf("failed to delete old session: %v", err)
    }

    id, err := GenerateUUID()
    if err != nil {
        tx.Rollback()
        return "", err
    }

  
    expireDate := time.Now().Add(24 * time.Hour).Unix()

    
    _, err = tx.Exec(queryInsert, id, userid, expireDate)
    if err != nil {
        tx.Rollback()
        return "", fmt.Errorf("failed to insert new session: %v", err)
    }

  
    if err = tx.Commit(); err != nil {
        return "", fmt.Errorf("failed to commit transaction: %v", err)
    }

    return id, nil
}
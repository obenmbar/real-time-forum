package forum

import (
	"database/sql"
	
)

func InsertUser(db *sql.DB, data Users, Id string, hashpassword string) error {
	query := "INSERT INTO users (id,nickname,age,gender,first_name_last_name,email,password)VALUES(?,?,?,?,?,?,?)"

	
	 _ , err:= db.Exec(query,Id,data.Nickname,data.Age,data.Gender,data.FirstName,data.LastName,data.Email,data.Password)
	 if err != nil {
		return err
	 }
	 
     return nil
}

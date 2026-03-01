package forum

import (
	"database/sql"
	"encoding/json"
	"fmt"
	forum "forum/utils"
	"strings"

	"net/http"
)

type DB struct {
	Db *sql.DB
}

func (db *DB) Registerhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "../frontend/index.html")
		return
	} else if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad reaquest"})
		return
	}

	var users forum.Users

	err := json.NewDecoder(r.Body).Decode(&users)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad reaquest"})
		return
	}

	err = forum.ValidUserdata(users)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"messege": err.Error()})
		return
	}

	hashpassword, err := forum.GeneratePassword(users.Password)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"messege": err.Error()})
		return
	}

	Id, err := forum.GenerateUUID()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"messege": err.Error()})
		return
	}

	err = forum.InsertUser(db.Db, users, Id, hashpassword)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if strings.Contains(err.Error(), "email") {
			json.NewEncoder(w).Encode(map[string]string{"messege": "bad reaquest  email deja exist"})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"messege": "bad reaquest  nickname deja exist"})
		}
		return
	}

	fmt.Println("passwordhash", hashpassword)
	fmt.Println("uuid", Id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"messege": "User created successfully!", "status": "success"})
}

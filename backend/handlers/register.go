package forum

import (
	"database/sql"
	"encoding/json"

	forum "forum/utils"
	"time"

	"strings"

	"net/http"
)

type DB struct {
	Db *sql.DB
}

func (db *DB) Registerhandler(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPost {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(map[string]string{"messege": "Method not allowed"})
        return
    }

    var users forum.Users
    err := json.NewDecoder(r.Body).Decode(&users)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"messege": "bad request: invalid json"})
        return
    }


    if err = forum.ValidUserdata(users); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"messege": err.Error()})
        return
    }


    hashpassword, err := forum.GeneratePassword(users.Password)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"messege": "internal error: hashing failed"})
        return
    }

   
    userId, _ := forum.GenerateUUID()
    if err = forum.InsertUser(db.Db, users, userId, hashpassword); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusConflict) 
        msg := "nickname already exists"
        if strings.Contains(err.Error(), "email") {
            msg = "email already exists"
        }
        json.NewEncoder(w).Encode(map[string]string{"messege": msg})
        return
    }

    sessionId, err := forum.CreateSession(db.Db, userId)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"messege": "failed to create session"})
        return
    }

   
    http.SetCookie(w, &http.Cookie{
        Name:     "session_tocken",
        Value:    sessionId,
        Path:     "/",
        Expires:  time.Now().Add(24 * time.Hour),
        HttpOnly: true,
        SameSite: http.SameSiteLaxMode,
    })

 
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "status":  "success",
        "messege": "User created successfully!",
    })
   
}
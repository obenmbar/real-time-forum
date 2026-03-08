package forum

import (
	"encoding/json"
	forum "forum/utils"
	"net/http"
	"time"
)

func (db *DB) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"messege": "method not allowed"})
		return
	}

	var loginData forum.Login
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"messege": "Invalid Json"})
		return
	}

	userid, err := forum.ValidLoginData(db.Db, loginData)
	if err != nil {
		if userid == "walo" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"messege": err.Error()})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"messege": "Invalid Password"})
		return
	}

	sessionId, err := forum.CreateSession(db.Db, userid)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"messege": "failed to create session"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_tocken",
		Value:    sessionId,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteLaxMode,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"messege": "login successfully!",
	})

}

package forum

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func (db *DB) AuthMiddleware(nextfunchandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cockie, err := r.Cookie("session_tocken")
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"messege": "Login required"})
			return
		}
		var expireddate int64
		var userid string
		query := "SELECT user_id,expires_at FROM sessions WHERE id = ?"

		err = db.Db.QueryRow(query, cockie.Value).Scan(&userid, &expireddate)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"messege": "invalid session"})
			return
		}

		if time.Now().Unix() > expireddate {
			db.Db.Exec("DELETE FROM sessions WHERE id = ?",cockie.Value)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"messege": "Session expired"})
			return
		}
		ctx :=  context.WithValue(r.Context(),"user_id",userid) 
         nextfunchandler(w,r.WithContext(ctx))
	}
}

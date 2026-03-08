package forum

import (
	"encoding/json"
	"net/http"
)

func (db *DB) LgoutHandler(w http.ResponseWriter, r *http.Request){

	if (r.Method != http.MethodPost){
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"messege":"method not allowed"})
		return
	}

	cockie, err := r.Cookie("session_tocken")
	if err == nil {
	_, err = db.Db.Exec("DELETE FROM sessions WHERE id = ?",cockie.Value)
	 if err != nil {
         w.Header().Set("Content-type","application/json")
		 w.WriteHeader(http.StatusInternalServerError)
		 json.NewEncoder(w).Encode(map[string]string{"messege":"internal server error, try later"})
	 }
	}

  http.SetCookie(w,&http.Cookie{
	Name: "session_tocken",
	Value: "",
	Path: "/",
	HttpOnly: true,
	MaxAge: -1,
	SameSite: http.SameSiteLaxMode,
  })

w.Header().Set("Content-Type","application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(map[string]string{"messege":"Logged out successfully",
 "status":"success"})
}
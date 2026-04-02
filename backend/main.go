package main

import (
	"fmt"
	forum "forum/utils"
	nki "forum/handlers"
	"net/http"
)

func main() {
	db, err := forum.InitialeDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	serverDB := nki.DB{Db: db}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/index.html")
	})

	mux.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
                   serverDB.Registerhandler(w,r)
	})
       
	mux.HandleFunc("/api/login",func(w http.ResponseWriter, r *http.Request) {
		serverDB.LoginHandler(w,r)
	})

    mux.Handle("/static/",http.StripPrefix("/static/",nki.SafeFileServer()))

	fmt.Println("server started: http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("error in starting server")
		return
	}

}

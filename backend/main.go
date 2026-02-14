package main

import (

	"fmt"
	forum "forum/utils"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := forum.InitialeDb()
  if err != nil {
	fmt.Println(err)
	return 
  }
  defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
	})
	fmt.Println("server started: http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("error in starting server")
		return
	}

}

package main

import (
	"fmt"
	"net/http"

	"github.com/Manuelda10/restaurant-backend/config"
	router "github.com/Manuelda10/restaurant-backend/internal/router"
)

func main(){
	db, err := config.InitDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := router.Router(db)
	port := 3000
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	errServer := http.ListenAndServe(addr, router)

	if errServer != nil {
		panic(errServer)
	}
}


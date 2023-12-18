package main

import (
	"fmt"
	"net/http"

	router "github.com/Manuelda10/restaurant-backend/internal/router"
)

func main(){
	router := router.NewRouter()
	port := 3000
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)

	if err != nil {
		panic(err)
	}
}


package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/restaurant-backend/controllers"
	"github.com/restaurant-backend/models"
);

type plate struct {
	ID string `json:"id"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}

func main(){
	godotenv.Load()
	fmt.Println("Hola mundo")
	models.ConnectDatabase()
	fmt.Println("Conexión establecida")

	router := controllers.InitializeRoutes()
	router.Run("localhost:9090")

	//handler := controllers.New()
	//handler.ServeHTTP()

	/*server := &http.Server{
		Addr: "0.0.0.0:8008",
		Handler: handler,
	}

	
	server.ListenAndServe()*/
}
package config

import (
	"database/sql"
	"fmt"
	"log"
)

func InitDatabase() (*sql.DB, error){
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("cannot ping database:", err)
	}

	log.Println("Database connected")	
	return db, nil
}
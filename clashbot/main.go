package main

import (
	"clashbot/config"
	"clashbot/cr"
	"clashbot/service"
	"clashbot/tg"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found but who cares lol")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load: %v", err)
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DatabaseUser, cfg.DatabasePassword, "db", cfg.DatabasePort, cfg.DatabaseName)
	time.Sleep(5 * time.Second)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("error connection:", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("DB ping failed:", err)
	}

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("error connection:", err)
	}
	defer db.Close()	
	crClient := cr.NewClient(cfg.ClashApiKey)
	svc := service.NewService(crClient)

	bot := tg.NewBot(cfg.BotToken, svc, db)

	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}
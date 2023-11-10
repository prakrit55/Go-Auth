package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ChatDB struct {
	db *sql.DB
}

func NewChatDB() (*ChatDB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	// .env file contains the POSTGRESS url
	POSTGRESS := os.Getenv("POSTGRESS_URL")

	db, err := sql.Open("postgres", POSTGRESS)
	if err != nil {
		return nil, err
	}
	return &ChatDB{db: db}, nil
}

func (d *ChatDB) Close() {
	d.db.Close()
}

func (d *ChatDB) GetDB() *sql.DB {
	return d.db
}

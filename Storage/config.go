package storage

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Login    string
	Password string
	Adress   string
	Port     string
	DbName   string
}

func loadEnv() {
    err := godotenv.Load("Storage/config.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

func NewConfig() *Config {
	loadEnv()
	login := os.Getenv("LOGIN")
	password := os.Getenv("PASSWORD")
	adress := os.Getenv("ADRESS")
	port := os.Getenv("PORT")
	dbName := os.Getenv("DBNAME")
	return &Config{Login: login, Password: password, Adress: adress, Port: port, DbName: dbName}
	//прочитать файл .env
	//запихнуть в структуру и вернуть указатель

}

package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Database struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

type App struct {
	APP_NAME    string
	APP_DOMAIN  string
	APP_PORT    string
	API_VERSION string
	API_PATTERN string
	SIGN_SECRET string
	ALLOWED_IP  []string
}

type Email struct {
	SMTP_HOST     string
	SMTP_PORT     int
	SMTP_USER     string
	SMTP_PASSWORD string
	FROM_EMAIL    string
	WEB_NAME      string
	WEB_DOMAIN    string
	APP_NAME      string
}

var err = godotenv.Load(".env")

func DBConfig() Database {
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	conf := Database{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}
	return conf
}

func AppConfig() App {
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	allowdIps := strings.Split(os.Getenv("ALLOWED_IP"), ",")
	conf := App{
		APP_NAME:    os.Getenv("APP_NAME"),
		APP_PORT:    os.Getenv("APP_PORT"),
		APP_DOMAIN:  os.Getenv("APP_DOMAIN"),
		API_VERSION: os.Getenv("API_VERSION"),
		API_PATTERN: os.Getenv("API_PATTERN"),
		SIGN_SECRET: os.Getenv("SIGN_SECRET"),
		ALLOWED_IP:  allowdIps,
	}
	return conf
}

func EmailConfig() Email {
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatal("Error converting port")
	}
	conf := Email{
		SMTP_HOST:     os.Getenv("SMTP_HOST"),
		SMTP_PORT:     port,
		SMTP_USER:     os.Getenv("STMP_USER"),
		SMTP_PASSWORD: os.Getenv("SMTP_PASSWORD"),
		FROM_EMAIL:    os.Getenv("FROM_EMAIL"),
		WEB_NAME:      os.Getenv("WEB_NAME"),
		WEB_DOMAIN:    os.Getenv("WEB_DOMAIN"),
	}
	return conf
}

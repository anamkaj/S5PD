package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AuthToken struct {
	AccessToken string
	DirectTable string
	StatusTable string
	Password    string
	Clients     string
	UrlApi      string
}

func GetToken() (AuthToken, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
		return AuthToken{}, err
	}

	access_token := os.Getenv("ACCESS_TOKEN")
	client_table := os.Getenv("DIRECT_TABLE")
	status_table := os.Getenv("STATUS_TABLE")
	password := os.Getenv("PASSWORD")
	client := os.Getenv("CLIENT")
	url_api := os.Getenv("URLAPI")

	token := AuthToken{
		AccessToken: access_token,
		DirectTable: client_table,
		StatusTable: status_table,
		Password:    password,
		Clients:     client,
		UrlApi:      url_api,
	}

	return token, nil
}

package utils

import (
	"fmt"
	"os"
)

func ConnectionUrlBuilder(serverName string) (string, error) {

	var url string
	switch serverName {
	case "postgres":
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		sslmode := os.Getenv("DB_SSL_MODE")
		url = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)
	case "fiber":
		host := os.Getenv("SERVER_HOST")
		port := os.Getenv("SERVER_PORT")
		url = fmt.Sprintf("%s:%s", host, port)
	default:
		return "", fmt.Errorf("connection name '%v' is not supported", serverName)
	}
	return url, nil
}

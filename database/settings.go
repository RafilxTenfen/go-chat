package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rhizomplatform/log"
)

// Settings basic setup connection for the database
type Settings struct {
	Host         string
	Port         uint16
	DatabaseName string
	User         string
	Password     string
}

// SettingsFromEnv return a new settings based on a .env file
func SettingsFromEnv() (Settings, error) {
	var s Settings

	if err := godotenv.Load(".env"); err != nil {
		log.Error(err) // just print the error if a .env does not exists
	}

	host, ok := os.LookupEnv("DATABASE_HOST")
	if !ok {
		return s, fmt.Errorf("Error on read HOST enviroment value")
	}
	s.Host = host

	portStr, ok := os.LookupEnv("DATABASE_PORT")
	if !ok {
		return s, fmt.Errorf("Error on read PORT enviroment value")
	}

	port, err := strconv.ParseUint(portStr, 10, 64)
	if err != nil {
		return s, err
	}
	s.Port = uint16(port)

	user, ok := os.LookupEnv("DATABASE_USER")
	if !ok {
		return s, fmt.Errorf("Error on read User enviroment value")
	}
	s.User = user

	password, ok := os.LookupEnv("DATABASE_PASSWORD")
	if !ok {
		return s, fmt.Errorf("Error on read Password enviroment value")
	}
	s.Password = password

	databaseName, ok := os.LookupEnv("DATABASE_NAME")
	if !ok {
		return s, fmt.Errorf("Error on read Database Name enviroment value")
	}
	s.DatabaseName = databaseName

	log.With(log.F{
		"Host":         s.Host,
		"Port":         s.Port,
		"User":         s.User,
		"DatabaseName": s.DatabaseName,
	}).Debug("load database env variables")
	return s, nil
}

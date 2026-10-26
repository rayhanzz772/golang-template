package config

import (
	"os"
	"strconv"
)

type Config struct {
	JWT struct {
		Secret []byte // kunci rahasia untuk JWT signing & validation
		Expiry int    // masa berlaku token (dalam jam)
	}
	Server struct {
		Port string // port server (misal "8000")
		Host string // host server (misal "localhost" atau "0.0.0.0")
	}
}

var App Config

func LoadConfig() {
	App.JWT.Secret = []byte(os.Getenv("JWT_SECRET"))

	expiry, err := strconv.Atoi(os.Getenv("JWT_EXPIRY"))
	if err != nil || expiry <= 0 {
		expiry = 24
	}
	App.JWT.Expiry = expiry

	App.Server.Port = os.Getenv("PORT")
	if App.Server.Port == "" {
		App.Server.Port = "8000"
	}

	App.Server.Host = os.Getenv("SERVER_HOST")
	if App.Server.Host == "" {
		App.Server.Host = "0.0.0.0"
	}
}

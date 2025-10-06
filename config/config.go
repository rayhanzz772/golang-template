package config

import (
	"os"
	"strconv"
)

type Config struct {
	JWT struct {
		Secret []byte // kunci rahasia untuk JWT signing & validation
		Expiry int    // masa berlaku token (biasanya dalam menit atau jam)
	}
	Server struct {
		Port string // port server (misal ":8080")
		Host string // host server (misal "localhost" atau "0.0.0.0")
	}
}

var App Config

func LoadConfig() {
	App.JWT.Secret = []byte(os.Getenv("JWT_SECRET"))

	expiry, _ := strconv.Atoi(os.Getenv("JWT_EXPIRY"))
	App.JWT.Expiry = expiry

	App.Server.Port = os.Getenv("SERVER_PORT")
	App.Server.Host = os.Getenv("SERVER_HOST")
}

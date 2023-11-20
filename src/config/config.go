package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectString = ""
	Porta         = 0
	SecretKey     []byte
)

func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Porta = 5000
	}
	ConnectString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_NOME"))
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}

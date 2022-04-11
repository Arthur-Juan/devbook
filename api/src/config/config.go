package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	//StrConn é a string de conexão com o banco
	StrConn = ""

	//Port é a porta que a API vai rodar
	Port = 0

	//CHAVE PARA ASSINAR O JWT
	SecretKey []byte
)

//Load carrega as variaveis de ambiente
func Load() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		Port = 9000
	}

	StrConn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}

package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Port = 0
	StringCon = ""
)

//carrega as variaveis de ambiente
func Loading(){

	var err error

	//carregando os arquivos .env
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
		
	}


	//lendo a porta http
	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 9000
	}


	//string de conexão com o bd
	StringCon = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
	os.Getenv("DB_USERS"),
	os.Getenv("DB_SENHA"),
	os.Getenv("DB_NOME"),)


}
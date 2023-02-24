package banco

import (
	"database/sql"  

_ "github.com/go-sql-driver/mysql"
	"crud-api/src/config"
)

func ConnectionDB() (*sql.DB, error){

	//abrindo o banco de dados
	db, err:= sql.Open("mysql", Config.StringCon)
	if err != nil {
		return nil, err
	}

	//vendo se a conexão ainda esta ativa
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, err

}
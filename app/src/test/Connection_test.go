package test

import (
	//Config "crud-api/src/config"
	"database/sql"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

//teste para ver se a conexao com o banco de dados esta funcionando
func TestDBconnection(T *testing.T){

	db, err := sql.Open("mysql", "root:davi2003@/crud?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			T.Fatalf("erro ao se conectar com o banco: %s", err.Error())
		}

		defer db.Close()

		err = db.Ping()
		if err != nil {
			T.Fatalf("erro ao testar a conexao com o banco: %s", err.Error())
		}
	

}
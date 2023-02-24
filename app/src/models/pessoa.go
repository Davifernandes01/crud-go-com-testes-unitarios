package models


type cadastro struct{
	id int64 `json: "id, omitempyt"`
	nome string `json: "nome, omitempyt"`
	idade string `json: "idade, omitempyt"`
	cpf string `json: "cpf, omitempyt"`
	telefone string `json: "telefone, omitempyt,"`
}
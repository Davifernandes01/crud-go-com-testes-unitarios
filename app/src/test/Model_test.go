package test

import (
	"crud-api/src/models"
	"testing"
)

func TestModel(T *testing.T) {

	

	//vendo se a validacao dos valores cadastros estao funcionando certinho
	_,err:= models.NewCadastro(1,"davi", "76", "16227523682", "(34) 99873-3098")
	if err != nil {
		T.Fatalf("houve um erro, erro: %s", err)
	}

	_ ,err1:= models.NewCadastro(2,"rada", "56", "78173298343", "(65) 2258-0608")
	if err1 != nil {
		T.Fatalf("houve um erro, erro: %s", err)
	}

	_,err2:= models.NewCadastro(3,"minos", "34", "66624414308", "(86) 2261-7525")
	if err2 != nil {
		T.Fatalf("houve um erro, erro: %s", err)
	}

	_,err3:= models.NewCadastro(4,"aiacos", "23", "82567491562", "(35) 2781-4703")
	if err3 != nil {
		T.Fatalf("houve um erro, erro: %s", err)
	}

	//vendo se os erros estao funcionando certinho

	user,_:= models.NewCadastro(1,"","76", "82567491562","(34) 99873-3098")
	if user != nil {
		T.Error("erro esperando: 'falta de nome no cadastro'" )
	}
	

	user1,_:= models.NewCadastro(1,"rada","760", "82567491562","(34) 99873-3098")
	if user1 != nil {
		T.Error("erro esperando: 'idade invalida, ou a falta dela'" )
	}

	user2,_:= models.NewCadastro(1,"shion","76", "","(34) 99873-3098")
	if user2 != nil {
		T.Error("erro esperando: 'falta o cpf, ou cpg invalido'" )
	}

	user3,_:= models.NewCadastro(1,"dohko","76", "82567491562"," ")
	if user3 != nil {
		T.Error("erro esperando: 'telefone invalido, ou a falta dele'" )
	}

	
}

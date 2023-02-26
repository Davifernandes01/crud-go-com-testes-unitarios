package models

import (
	"errors"
	"regexp"
	"strings"

	"github.com/klassmann/cpfcnpj"
)

// struct principal de cadastro

type Cadastro struct {
	Id int64 `json: "id, omitempyt"`
	Nome string `json: "nome, omitempyt"`
	Idade string `json: "idade, omitempyt"`
	Cpf string `json: "cpf, omitempyt"`
	Telefone string `json: "telefone, omitempyt,"`
}

//metados

// essa funcao ira funcionar como se fosse um "construtor"
func NewCadastro(id int64, nome string, idade string, cpf string, telefone string)(*Cadastro, error){

	C:= &Cadastro{
		Id:       id,
		Nome:     nome,
		Idade:    idade,
		Cpf:      cpf,
		Telefone: telefone,
	}

	if err:= C.prepare(); err != nil {
		return nil, err
	}

	return C, nil
}

func (c *Cadastro) prepare() error{
	
	

	if err:= c.valid(); err != nil {
		return err
	}

	c.format()

	return nil
}

//validando os campos
func (c *Cadastro) valid() error{
	if c.Nome == ""{
		return errors.New("por favor digite um nome")
	}

	//validando a idade
	if err:= regexIdade(c.Idade); err != nil {
		return err
	}

	//validando o cpf regex: xxxxxxxxxxx (sem pontuacao)
	 err:= cpfcnpj.ValidateCPF(c.Cpf)
	 if !err {
		return errors.New("cpf invalido")
	 }
		
	 if err:= regexTelefone(c.Telefone); err != nil {
		return err
	 }
	 
	return nil
}


//tirando os espaços em branco do começo e do fim
func (c *Cadastro) format(){

	c.Nome = strings.TrimSpace(c.Nome)
	c.Idade = strings.TrimSpace(c.Idade)
	c.Cpf = strings.TrimSpace(c.Cpf)
	c.Telefone = strings.TrimSpace(c.Telefone)

}

//validando o numero de telefone regex: xx xxxxx-xxxx
func regexTelefone(telefone string) error{

	regex := regexp.MustCompile(`^\([1-9]{2}\) (?:[2-8]|9[1-9])[0-9]{3}\-[0-9]{4}$`)

	if !regex.MatchString(telefone){
		return errors.New("telefone invalido!")
	}

		return nil

}

//validando a idade (de 1 ate 99)
func regexIdade(idade string) error {
	regex:= regexp.MustCompile(`^[1-9][0-9]?$|^100$`)

	if !regex.MatchString(idade){
		return errors.New("idade invalida!!")
	}

	return nil
}


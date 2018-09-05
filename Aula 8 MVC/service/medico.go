package service

import (
	"time"

	"../model"
)

//AlgumaCoisaMedico ...
func AlgumaCoisaMedico() model.Pessoa {
	var pessoa model.Pessoa
	pessoa.Nome = "teste1"
	pessoa.Idade = 32
	pessoa.DataNascimento = time.Now()
	return pessoa
}

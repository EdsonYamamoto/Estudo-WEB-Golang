package service

import (
	"time"

	"../model"
)

//SomarESubtrair ...
func SomarESubtrair(var1 int, var2 int) (int, int) {
	soma := var1 + var2
	subtrair := var1 - var2
	return soma, subtrair
}

//CriaPessoa ...
func CriaPessoa() model.Pessoa {
	var pessoa model.Pessoa
	pessoa.Nome = "teste1"
	pessoa.Idade = 32
	pessoa.DataNascimento = time.Now()
	return pessoa
}

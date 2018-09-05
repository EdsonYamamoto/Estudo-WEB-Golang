package controller

import (
	"log"
	"time"

	"../service"
)

func FazAlgumaCoisaComPessoa() {

	pessoa := service.CriaPessoa()

	diaDeHoje := time.Now().Day()

	soma, subtrair := service.SomarESubtrair(pessoa.Idade, diaDeHoje)
	log.Println("soma da idade da pessoa mais o dia atual: ", soma)
	log.Println("subtracao da idade da pessoa mais o dia atual: ", subtrair)
	return
}

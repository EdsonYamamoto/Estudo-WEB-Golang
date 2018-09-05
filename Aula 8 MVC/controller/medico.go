package controller

import (
	"log"

	"../service"
)

func Medico() {

	pessoa := service.AlgumaCoisaMedico()
	log.Println(pessoa)
	return
}

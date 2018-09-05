package main

import (
	"log"

	ctr "./controller"
)

func main() {
	log.Println("Executa o primeiro metodo")
	ctr.FazAlgumaCoisaComPessoa()

	log.Println("Executa outra coisa")
	ctr.Medico()

	log.Println("finaliza programa")
}

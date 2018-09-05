package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ModelTeste struct {
	Var1 string `json="var1"`
	Var2 int    `json="var2"`
	Var3 bool   `json="var3"`
}

func main() {
	r := gin.Default()
	//carrego todos os templates dentro da pasta template
	r.LoadHTMLGlob("templates/*")
	//index
	r.GET("/", func(c *gin.Context) {
		// eu devo executar o metodo Index e passo
		Index(c)
	})
	r.Run(":8080")

}

func Index(c *gin.Context) {
	var Vetor []ModelTeste
	var objeto ModelTeste
	objeto.Var1 = "valor 1"
	objeto.Var2 = 54
	objeto.Var3 = false

	var objeto2 ModelTeste
	objeto2.Var1 = "nao tenho nada"
	objeto2.Var2 = 10
	objeto2.Var3 = true

	for index := 0; index < 5; index++ {
		//funcao que adiciona uma estrutura dentro de um vetor
		Vetor = append(Vetor, objeto)
		Vetor = append(Vetor, objeto2)
	}

	//estrutura que uso para passar as informações para o template
	data := struct {
		VariavelString string
		VetorTeste     []ModelTeste
	}{
		VariavelString: "oi",
		VetorTeste:     Vetor,
	}
	//metodo usado pra carregar a pagina trazerValores da pasta template
	c.HTML(http.StatusOK, "trazerValores.tmpl.html", data)
}

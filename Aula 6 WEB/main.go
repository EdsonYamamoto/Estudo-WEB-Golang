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
	r.GET("/paginaRuim", func(c *gin.Context) {
		// eu devo executar o metodo Index e passo
		OutroMetodo(c)
	})

	r.Run(":8080")

}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", nil)
}
func OutroMetodo(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"erro": "O quem requisita recebe uma mensagem que deu erro"})
}

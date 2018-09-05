package main

import (
	"log"
	//biblioteca GIN para rodar o servidor

	"github.com/gin-gonic/gin"
)

type Teste struct {
	Var1 string `json="var1"`
	Var2 int    `json="var2"`
}

func main() {
	r := gin.Default()
	//index
	r.GET("/", func(c *gin.Context) { Index(c) })
	//exemplo de params
	r.GET("/paramExample/:id", func(c *gin.Context) { ParamGet(c) })
	//este cara daqui vc precisa usar o POSTMAN é um programa que vai precisar ser instalado
	//o JSON enviado
	//EX:
	//	{
	//		"var1":"jansdkjanskjd"
	//		"var2":50
	//	}
	r.POST("/post", func(c *gin.Context) { ExemploPost(c) })

	r.Run(":8090")

}

func Index(c *gin.Context) {
	log.Println("entrei aqui")
}

func ParamGet(c *gin.Context) {
	log.Println("entrei acola")
	id := c.Param("id")
	log.Println(id)
}

func ExemploPost(c *gin.Context) {
	var t Teste
	err := c.BindJSON(&t)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(t)
}

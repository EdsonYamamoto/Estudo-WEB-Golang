//pacote principal
package main

//import das bibliotecas utilizadas de
import (
	//paginas html
	"html/template"
	//sistema para printar
	"log"
	//biblioteca pra criar o servidor
	"net/http"
)

func main() {
	log.Println("Server is up and listening on port 8090.")
	//toda pagina web como default entra aqui primeiro
	http.HandleFunc("/", index)
	//a porta que esta sendo lida é o 8090
	http.ListenAndServe(":8090", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	//carrego a pagina index.html
	tpl, _ := template.ParseFiles("index.html")
	//devo retornar uma mensagem de OK
	w.WriteHeader(http.StatusOK)
	//devo executar mostrar a pagina pro usuario
	tpl.Execute(w, nil)
}

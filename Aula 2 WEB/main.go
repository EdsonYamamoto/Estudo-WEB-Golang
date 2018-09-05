//pacote principal
package main

//import das bibliotecas utilizadas
import (
	"html/template"
	"log"
	"net/http"
)

//Estrutura
type DataTeste struct {
	Titulo string
	var1   string
}

func main() {
	log.Println("Server is up and listening on port 8090.")
	http.HandleFunc("/", index)
	http.HandleFunc("/pagina1", pag1)
	http.ListenAndServe(":8090", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	tpl, _ := template.ParseFiles("index.html")

	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, nil)
}

func pag1(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("pagina1.html")

	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, nil)
}

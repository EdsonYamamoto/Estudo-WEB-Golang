/*
Programa que percorre um JSON e retorna o Modelo cujo nome foi passado
*/
package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

//estrutura que deve receber todos os itens que estão salvos no appDoc.json
//Api ...
type Api struct {
	ApiName        string `json: "apiTeste"`
	ApiDescription string `json: "apiDescription"`
	ApiServer      string `json: "apiServer"`
	ApiUser        string `json: "apiUser"`
	ApiToken       string `json: "apiToken"`
}

func main() {
	//metodo que faz a busca no arquivo de JSON pra trazer a informação
	api, err := GetAPI("apiTeste")
	if err != nil {
		log.Println(err)
	}
	log.Println(api)
}

//GetAPI ...
func GetAPI(apiName string) (Api, error) {
	var allAPI []Api
	var API Api
	raw, err := ioutil.ReadFile("./appDoc.json")
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(raw, &allAPI)
	for i := 0; i < len(allAPI); i++ {
		if allAPI[i].ApiName == apiName {
			API = allAPI[i]
			return API, nil
		}
	}
	return API, errors.New("API nao encontrada")
}

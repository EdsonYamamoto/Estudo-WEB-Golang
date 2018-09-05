package main

import (
	"log"

	"time"

	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
	mapstructure "github.com/mitchellh/mapstructure"
)

type Pessoa struct {
	Nome       string    `json:"nome" mapstructure:"Nome"`
	Var1       string    `json:"var1" mapstructure:"Nome"`
	Idade      int       `json:"idade" mapstructure:"Idade"`
	Faculdade  string    `json:"faculdade" mapstructure:"Faculdade"`
	Nascimento time.Time `json:"nascimento" mapstructure:"Nascimento"`
}

func main() {
	//Selecionar um item por documento de uma collection
	/*
			m, err := GetOneDataByDoc("Medicos", "teste3")
		//deu algum devo printar algo na tela
			if err != nil {
				log.Println(err)
			}
			log.Println(m)
			var p1 Pessoa
			err = mapstructure.Decode(m, &p1)
		//deu algum devo printar algo na tela
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(p1)
	*/
	//Retornar todos registros por collection
	/*
			results, err := GetAllDataFromCollection("Medicos")
		//deu algum devo printar algo na tela
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(results)
	*/
	//Salva um unico dado na collection
	var a1 Pessoa
	a1.Faculdade = "Esamc"
	a1.Idade = 50
	a1.Var1 = "batata sei la"
	a1.Nascimento = time.Now()
	a1.Nome = "Eduarda"
	err := SaveUniqueInterfaceDataIntoCollection("Alunos", a1)
	//deu algum devo printar algo na tela
	if err != nil {
		log.Println(err)
	}
}

//DataBaseAccess ...
func DataBaseAccess() (*firestore.Client, error) {
	//aonde o arquivo estiver o json do firebase deve estar junto
	opt := option.WithCredentialsFile("firebase.json")

	//aqui eu abro a conexão pra qualquer coisa no banco
	app, err := firebase.NewApp(context.Background(), nil, opt)
	//deu algum erro eu devo sair
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//aqui eu pego a conexão e seto dentro de um client pra podermos usar no programa
	client, err := app.Firestore(context.Background())
	//deu algum erro eu devo sair
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return client, err
}

//GetOneDataByDoc ... Pega um unico dado com o ID do documento e a collection passada
func GetOneDataByDoc(collection string, document string) (map[string]interface{}, error) {
	//retorna as informações do banco
	client, err := DataBaseAccess()
	//deu algum erro eu devo sair
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result, err := client.Collection(collection).Doc(document).Get(context.Background())
	//deu algum erro eu devo sair
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//armazeno todos os dados em "m"
	m := result.Data()
	//fecho a conexão com o banco
	defer client.Close()

	return m, err
}

//GetAllDataFromCollection ... Retorna todos os dados numa unica coleção
func GetAllDataFromCollection(collection string) ([]map[string]interface{}, error) {
	//retorna as informações do banco
	var results []map[string]interface{}

	//executo o acesso ao client do banco
	client, err := DataBaseAccess()
	//deu algum erro eu devo sair
	if err != nil {
		log.Println(err)
	}
	//comando where é o mesmo do SQL SERVER
	//result := client.Collection(collection).Where("ativo", "==", true).Documents(context.Background())
	result := client.Collection(collection).Documents(context.Background())
	for {
		doc, err := result.Next()
		if err == iterator.Done {
			break
		}
		//deu algum erro eu devo sair
		if err != nil {
			return nil, err
		}
		//aqui eu concateno todos os dados numa string e retorno o resultado
		results = append(results, doc.Data())
	}
	return results, err
}

//SaveUniqueInterfaceDataIntoCollection ... salvar uma estrutura model na collection
func SaveUniqueInterfaceDataIntoCollection(collection string, model ...interface{}) error {
	var mapString map[string]interface{}
	mapstructure.WeakDecode(model, &mapString)
	client, err := DataBaseAccess()
	//deu algum erro eu devo sair
	if err != nil {
		return err
	}
	//aqui devo salvar todos os items dentro da collection indicada
	_, _, err = client.Collection(collection).Add(context.Background(), mapString)
	//deu algum erro eu devo sair
	if err != nil {
		return err
	}
	return err
}

//SaveUniqueMapStringDataIntoCollection ... salvar uma map string model na collection
func SaveUniqueMapStringDataIntoCollection(collection string, mapString map[string]interface{}) error {
	client, err := DataBaseAccess()
	if err != nil {
		return err
	}
	//aqui devo salvar todos os items dentro da collection indicada
	_, _, err = client.Collection(collection).Add(context.Background(), mapString)
	if err != nil {
		return err
	}
	return err
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type CEP struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
	Ibge       string `json:"ibge"`
	Ddd        string `json:"ddd"`
	Siafi      string `json:"siafi"`
}

func main() {
	cepFlag := flag.String("cep", "", "CEP para pesquisar")
	flag.Parse()

	if *cepFlag == "" {
		fmt.Println("Por favor, forneça um CEP usando -cep <cep>")
		os.Exit(1)
	}

	cep, err := SearchCEP(*cepFlag)
	if err != nil {
		fmt.Printf("Erro ao buscar CEP: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Endereço correspondente ao CEP %s:\n", cep.Cep)
	fmt.Printf("Logradouro: %s\n", cep.Logradouro)
	fmt.Printf("Bairro: %s\n", cep.Bairro)
	fmt.Printf("Localidade: %s\n", cep.Localidade)
	fmt.Printf("UF: %s\n", cep.Uf)
	fmt.Printf("IBGE: %s\n", cep.Ibge)
	fmt.Printf("DDD: %s\n", cep.Ddd)
	fmt.Printf("SIAFI: %s\n", cep.Siafi)
}

func SearchCEP(cep string) (*CEP, error) {
	response, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Erro ao fechar o corpo da resposta: %v\n", err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var c CEP
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

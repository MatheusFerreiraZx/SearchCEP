package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", SearchHandlerCEP)
	http.ListenAndServe(":8080", mux)
}

func SearchHandlerCEP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, error := SearchCEP(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}

func SearchCEP(cep string) (*CEP, error) {
	response, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}

	var c CEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}

	return &c, nil
}

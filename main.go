package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"mux"
)

func main() {
	rotas := mux.NewRouter().StrictSlash(true)

	rotas.HandleFunc("/", getAll).Methods("GET")
	rotas.HandleFunc("/cadastro", create).Methods("POST")
	var port = ":8000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, rotas))

}

type Personagem struct {
	Name  string
	Movie bool
	Serie bool
}

var cadastro = []Personagem{

	Personagem{Name: "Ragnar Lothbrok", Movie: false, Serie: true},
	Personagem{Name: "Lagertha", Movie: false, Serie: true},
}

func getAll(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(cadastro)
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var p Personagem

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &p); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	json.Unmarshal(body, &p)

	cadastro = append(cadastro, p)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}
}

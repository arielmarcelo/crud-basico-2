package main

import (
	"crud-basico/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//CRUD -> Create, Read, Update, Delete

	//CREATE -> post
	//READ   -> get
	//UPDATE -> put
	//DELETE -> delete

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuaio).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)

	fmt.Println("Executando na ROTA 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}

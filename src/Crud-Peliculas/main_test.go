package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()

}
func TestCreatePeliculas(t *testing.T) {
	//Inicializar el test
	test := httptest.NewServer(router)
	defer test.Close()
	//Crear la pelicula
	addPelicula := Pelicula{
		Isbn:   "1234",
		Titulo: "Pelicula 1",
		Director: &Director{
			PrimerNombre:  "Director 1",
			SegundoNombre: "Director 2",
		},
	}
	//Convertir la pelicula a json
	jsonPelicula, _ := json.Marshal(addPelicula)
	res, err := http.Post(test.URL+"/peliculas", "application/json", bytes.NewBuffer(jsonPelicula))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", res.Status)
	}
}

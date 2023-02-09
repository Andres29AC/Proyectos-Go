package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Pelicula struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Titulo   string    `json:"titulo"`
	Director *Director `json:"director"`
}
type Director struct {
	PrimerNombre  string `json:"PrimerNombre"`
	SegundoNombre string `json:"SegundoNombre"`
}

var peliculas []Pelicula

func getPeliculas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peliculas)
}
func deletePelicula(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range peliculas {
		if item.ID == params["id"] {
			peliculas = append(peliculas[:index], peliculas[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(peliculas)
}
func getPelicula(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range peliculas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Pelicula{})
}
func createPelicula(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pelicula Pelicula
	_ = json.NewDecoder(r.Body).Decode(&pelicula)
	pelicula.ID = strconv.Itoa(rand.Intn(10000000))
	peliculas = append(peliculas, pelicula)
	json.NewEncoder(w).Encode(pelicula)
}
func updatePelicula(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range peliculas {
		if item.ID == params["id"] {
			peliculas = append(peliculas[:index], peliculas[index+1:]...)
			var pelicula Pelicula
			_ = json.NewDecoder(r.Body).Decode(&pelicula)
			pelicula.ID = params["id"]
			peliculas = append(peliculas, pelicula)
			json.NewEncoder(w).Encode(pelicula)
			return
		}
	}
}

//func handleIndex(w http.ResponseWriter, r *http.Request) {
//	http.ServeFile(w, r, "./static/index.html")
//}

func main() {
	//fileServer := http.FileServer(http.Dir("./static/index.html"))
	//http.HandleFunc("/index", handleIndex)
	//http.Handle("/", fileServer)
	//http.HandleFunc("/formulario", formularioHandler)
	//http.HandleFunc("/hola", holaHandler)

	router := mux.NewRouter()
	peliculas = append(peliculas, Pelicula{ID: "1", Isbn: "448743", Titulo: "El señor de los anillos", Director: &Director{PrimerNombre: "Peter", SegundoNombre: "Jackson"}})
	peliculas = append(peliculas, Pelicula{ID: "2", Isbn: "458844", Titulo: "Freedy Kruger", Director: &Director{PrimerNombre: "Wes", SegundoNombre: "Cronenberg"}})

	router.HandleFunc("/peliculas", getPeliculas).Methods("GET")
	router.HandleFunc("/peliculas/{id}", getPelicula).Methods("GET")
	router.HandleFunc("/peliculas", createPelicula).Methods("POST")
	router.HandleFunc("/peliculas/{id}", updatePelicula).Methods("PUT")
	router.HandleFunc("/peliculas/{id}", deletePelicula).Methods("DELETE")
	fmt.Printf("Iniciando servidor en el puerto 8081\n")
	log.Fatal(http.ListenAndServe(":8081", router))
	//http.ListenAndServe(":7000", router)
}

//Para Documentacion
//godoc -http=:6060

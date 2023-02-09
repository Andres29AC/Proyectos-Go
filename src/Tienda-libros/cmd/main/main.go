package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"Tienda-libros/pkg/routes"
)
func main(){
	r := mux.NewRouter()
	routes.RegistrarRutas(r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	http.Handle("/", r)
}

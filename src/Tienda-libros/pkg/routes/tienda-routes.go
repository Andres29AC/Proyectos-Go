package routes

import(
	"Tienda-libros/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegistrarRutas =func(router *mux.Router){
	router.HandleFunc("/libro/", controllers.CrearLibro).Methods("POST")
	router.HandleFunc("/libro/", controllers.GetLibro).Methods("GET")
	router.HandleFunc("/libro/{libroId}", controllers.GetLibroId).Methods("GET")
	router.HandleFunc("/libro/{libroId}", controllers.UpdateLibro).Methods("PUT")
	router.HandleFunc("/libro/{libroId}", controllers.DeleteLibro).Methods("DELETE")
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func formularioHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Hubo un error %v", err)
		return
	}
	fmt.Fprintf(w, "POST request establecido correctamente\n")
	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")
	edad := r.FormValue("edad")
	telefono := r.FormValue("telefono")
	carrera := r.FormValue("carrera")
	correo := r.FormValue("correo")
	contraseña := r.FormValue("contraseña")
	fmt.Fprintf(w, "Nombre: %s\n", nombre)
	fmt.Fprintf(w, "Apellido: %s\n", apellido)
	fmt.Fprintf(w, "Edad: %s\n", edad)
	fmt.Fprintf(w, "Telefono: %s\n", telefono)
	fmt.Fprintf(w, "Carrera: %s\n", carrera)
	fmt.Fprintf(w, "Correo: %s\n", correo)
	fmt.Fprintf(w, "Contraseña: %s\n", contraseña)
}
func holaHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hola" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Metodo no soportado!!.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hola Mundo!!")
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/formulario", formularioHandler)
	http.HandleFunc("/hola", holaHandler)
	fmt.Printf("Servidor iniciado en el puerto 8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		//panic(err)
		log.Fatal(err)
	}
}

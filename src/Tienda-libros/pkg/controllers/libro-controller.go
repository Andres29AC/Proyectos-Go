package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"Tienda-libros/pkg/utils"
	"Tienda-libros/pkg/models"
)
var NuevoLibro models.Libro

func GetLibro(w http.ResponseWriter, r *http.Request){
	nuevoLibro :=models.GetAllLibros()
	res, _ :=json.Marshal(nuevoLibro)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetLibroId(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	libroId := vars["libroId"]
	ID, err := strconv.ParseInt(libroId,0,0)
	if err != nil{
		fmt.Println("Error al parsear")
	}
	libroDetalles, _:= models.GetLibroId(ID)
	res, _:=json.Marshal(libroDetalles)
	w.Header().Set("Content Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CrearLibro(w http.ResponseWriter, r *http.Request){
	CrearLibro:= &models.Libro{}
	utils.ParseBody(r, CrearLibro)
	b:= CrearLibro.CrearLibro()
	res, _:=json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteLibro(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	libroId := vars["libroId"]
	ID, err := strconv.ParseInt(libroId,0,0)
	if err != nil{
		fmt.Println("Error while parsing")
	}
	libro := models.DeleteLibro(ID)
	res, _:= json.Marshal(libro)
	w.Header().Set("Content Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
//TODO: UpdateLibro 

func UpdateLibro(w http.ResponseWriter,r *http.Request){
	var updateLibro = &models.Libro{}
	utils.ParseBody(r, updateLibro)
	vars := mux.Vars(r)
	libroId := vars["libroId"]
	ID, err := strconv.ParseInt(libroId,0,0)
	if err != nil{
		fmt.Println("Error while Parsear")
	}
	libroDetalles, db := models.GetLibroId(ID)
	if  updateLibro.Nombre != ""{
		libroDetalles.Nombre = updateLibro.Nombre
	}
	if  updateLibro.Autor != ""{
		libroDetalles.Autor = updateLibro.Autor
	}
	if  updateLibro.Publicacion != ""{
		libroDetalles.Publicacion = updateLibro.Publicacion
	}
	db.Save(&libroDetalles)
	res, _ := json.Marshal(libroDetalles)
	w.Header().Set("Content Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//NOTE: .Marshal() convierte un objeto en un JSON 




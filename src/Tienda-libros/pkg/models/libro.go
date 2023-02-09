package models

import(
  "Tienda-libros/pkg/config"
	"github.com/jinzhu/gorm"
)
var db *gorm.DB
type Libro struct{
	gorm.Model
	Nombre string `gorm:""json:"nombre"`
	Autor string  `json:"autor"`
	Publicacion string `json:"publicacion"`
}
func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Libro{})
}

func(b *Libro) CrearLibro() *Libro{
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllLibros()[]Libro{
	var Libros []Libro
	db.Find(&Libros)
	return Libros
}
func GetLibroId(Id int64)(*Libro, *gorm.DB){
	var getLibro Libro
	db :=db.Where("ID=?", Id).Find(&getLibro)
	return &getLibro, db
}
func DeleteLibro(ID int64) Libro{
	var libro Libro
	db.Where("ID=?", ID).Delete(libro)
	return libro
}

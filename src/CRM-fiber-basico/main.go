package main

import (
	"fmt"
	"github.com/Andres29AC/Proyectos-Go/src/CRM-fiber-basico/lead"
	"github.com/Andres29AC/Proyectos-Go/src/CRM-fiber-basico/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//NOTE: GetLeads sirve para obtener todos los leads
//NOTE: GetLead sirve para obtener un lead en especifico
//NOTE: NewLead sirve para crear un nuevo lead
//NOTE: DeleteLead sirve para eliminar un lead en especifico

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Get("/api/v1/lead/:id",lead.GetLead)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("/api/v1/lead/:id",lead.DeleteLead)
}
func initDatabase(){
	var err error 
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	//NOTE: Traduccion de: Failed to connect to database
	//NOTE: Fallo al conectarse a la base de datos
	if err != nil { 
		panic("Fallo al conectarse a la base de datos")
	}
	fmt.Println("Se conecto a la base de datos")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("La base de datos esta migrada")
}


func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")
	defer database.DBConn.Close()
}

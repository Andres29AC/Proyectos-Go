package lead

import (
	"github.com/Andres29AC/Proyectos-Go/src/CRM-fiber-basico/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model 
	Nombre 		string  `json:"nombre"`
	Empresa 	string  `json:"empresa"`
	Correo 		string  `json:"correo"`
	Telefono 	int 	`json:"telefono"`
}
func GetLeads(c *fiber.Ctx){
	db := database.DBConn 
	var leads []Lead 
	db.Find(&leads)
	c.JSON(leads)
}
func GetLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn 
	var lead Lead 
	db.Find(&lead, id)
	c.JSON(lead)

}
func NewLead(c *fiber.Ctx){
	db := database.DBConn 
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil{
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}
func DeleteLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn 
	var lead Lead 
	db.First(&lead, id)
	if lead.Nombre == ""{
		c.Status(500).Send("No se encontro el lead con el ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead eliminado")
}






//NOTE: Sinonimos de Compañia tenemos:
//NOTE:  - Empresa 
//NOTE:  - Sociedad 
//NOTE:  - Negocio 
//NOTE:  - Firma 
//NOTE:  - Casa 
//NOTE:  - Industria 
//NOTE:  - Organizacion
//NOTE:  - Corporacion 
//NOTE:  - Consorcio 
//NOTE:  - Grupo 
//NOTE:  - Union 
//NOTE:  - Agencia 
//NOTE:  - Establecimiento 
//NOTE:  - Entidad 
//NOTE:  - Asociacion 
//NOTE:  - Institucion
//NOTE:  - Organismo 

package config

import(
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

// El caracter _ significa que se esta importando el paquete solo para sus
// efectos secundarios, en lugar de para su uso directo. En este caso, se
// esta importando el controlador "mysql" para Gorm, pero no se usara
// directamente en el codigo.

var(
	db * gorm.DB
)

func Connect(){
	dsn := "root:andres54AC@tcp(localhost:3306)/librosrest?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil{
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB{
	return db
}

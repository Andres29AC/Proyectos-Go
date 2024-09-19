package main 

import(
	"fmt"
	"os"
	"github.com/Andres29AC/Proyectos-Go/src/api-restaurant/middleware"
	"github.com/Andres29AC/Proyectos-Go/src/api-restaurant/database"
	"github.com/Andres29AC/Proyectos-Go/src/api-restaurant/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)

//NOTE: variables para la base de datos
var mealCollection *mongo.Collection = database.OpenCollection(database.Client, "meal")
func main() {	
	port := os.Getenv("PORT")
	if port == ""{
		port = "3001"
	}
	fmt.Println("Server running on port", port)

    //NOTE: Usando la base de datos
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.MealRoutes(router)
	routes.DishRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	//routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}

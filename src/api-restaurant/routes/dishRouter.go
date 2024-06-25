package routes 

import( 
	controller "Andres29AC/Proyectos-Go/src/api-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func DishRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/dishes", controller.GetDishes())
	incomingRoutes.GET("/dishes/:dish_id", controller.GetDish())
	incomingRoutes.POST("/dishes", controller.CreateDish())
	incomingRoutes.PATCH("/dishes/:dish_id", controller.UpdateDish())
}

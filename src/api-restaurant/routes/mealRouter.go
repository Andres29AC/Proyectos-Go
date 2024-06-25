package routes 

import(
	controller "Andres29AC/Proyectos-Go/src/api-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func MealRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/meals", controllers.GetMeals())
	incomingRoutes.GET("/meals/:meal_id", controllers.GetMeal())
	incomingRoutes.POST("/meals", controllers.CreateMeal())
	incomingRoutes.PATCH("/meals/:meal_id", controllers.UpdateMeal())
}

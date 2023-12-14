package routes

import(
	"github.com/Andres29AC/Proyectos-Go/src/Ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

//NOTE: incomingRoutes -> rutas entrantes


func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST ("/user/signup", controllers.Signup)
	incomingRoutes.POST ("/user/login", controllers.Login)
	incomingRoutes.POST ("/admin/addProduct", controllers.AddProduct)
	incomingRoutes.POST ("/admin/productView", controllers.ProductView)
	incomingRoutes.POST ("/admin/searchProduct", controllers.SearchProduct)

}

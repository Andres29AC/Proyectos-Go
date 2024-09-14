package middleware

import (
	"fmt"
	"net/http"
	helper "github.com/Andres29AC/Proyectos-Go/src/api-restaurant/helpers"
	"github.com/gin-gonic/gin"
)
func Authentication() gin.HandlerFunc{
	return func(c *gin.Context){
		clientToken := c.Request.Header.Get("token")
		if clientToken == ""{
			c.JSON(http.StatusInternalServerError,gin.H{"error": fmt.Sprintf("Se necesita un token para acceder a este recurso")})
			c.Abort()
			return
        }
		clains , err := helper.ValidateToken(clientToken)
		if err != ""{
			c.JSON(http.StatusInternalServerError,gin.H{"error": err})
			c.Abort()
			return
        }
		c.Set("email", clains.Email)
		c.Set("first_name", clains.First_name)
		c.Set("last_name", clains.Last_name)
		c.Set("uuid", clains.Uid)
		c.Next()
	}
}

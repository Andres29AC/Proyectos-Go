package controllers

import "github.com/gin-gonic/gin"

//NOTE: gin.HandlerFunc sirve para manejar las peticiones HTTP

//FIXME: Gin sirve para manejar las peticiones HTTP

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetInvoice() gin.HandlerFunc {
   return func(c *gin.Context) {}
   }
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

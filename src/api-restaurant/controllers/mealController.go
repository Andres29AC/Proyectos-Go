package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Andres29AC/Proyectos-Go/src/api-restaurant/database"
	"github.com/Andres29AC/Proyectos-Go/src/api-restaurant/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var mealCollection *mongo.Collection = database.OpenCollection(database.Client, "meal")


func GetMeals() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx,cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result,err := mealCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al listar los ingredientes del menu"})
        }
		var allMeals []bson.M
		if err = result.All(ctx, &allMeals); err != nil{
			log.Fatal(err)
        }
		c.JSON(http.StatusOK, allMeals)

	}
}
func GetMeal() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx,cancel = context.WithTimeout(context.Background(), 100*time.Second)
		mealId := c.Param("meal_id")
		var meal models.Meal
		err := mealCollection.FindOne(ctx, bson.M{"meal_id" : mealId}).Decode(&meal)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al encontrar la comida"})
        }
		c.JSON(http.StatusOK, meal)
	}
}

func CreateMeal() gin.HandlerFunc {
	return func(c *gin.Context) {
		var meal models.Meal
		var ctx,cancel = context.WithTimeout(context.Background(), 100*time.Second)
		if err := c.BindJSON(&meal); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
        }
		validationErr := validate.Struct(meal)
		if validationErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
        }
		meal.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		meal.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		meal.ID = primitive.NewObjectID()
		meal.Meal_id = meal.ID.Hex()
		result, insertErr := mealCollection.InsertOne(ctx, meal)
		if insertErr != nil{
			msg := fmt.Sprintf("Ingrediente no creado")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
        }
		defer cancel()
		c.JSON(http.StatusOK, result)
		defer cancel()
	}
}
func inTimeSpan(start, end, check time.Time) bool{
	return start.After(time.Now()) && end.After(start)
}

func UpdateMeal() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx,cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var meal models.Meal
		if err := c.BindJSON(&meal); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
        }
		mealId := c.Param("meal_id")
		filter := bson.M{"meal_id": mealId}
		var updateObj primitive.D
		if meal.Start_date != nil && meal.End_date != nil{
			if !inTimeSpan(*meal.Start_date, *meal.End_date, time.Now()){
				msg := "Por favor, vuelva a escribir la hora."
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				defer cancel()
				return
           }
		    updateObj = append(updateObj, bson.E{"start_date", meal.Start_date})
		    updateObj = append(updateObj, bson.E{"end_date", meal.End_date})
			if meal.Name != ""{
				updateObj = append(updateObj, bson.E{"name", meal.Name})
            }
			if meal.Category != ""{
				updateObj = append(updateObj, bson.E{"category", meal.Category})
            }
			meal.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			updateObj = append(updateObj, bson.E{"updated_at", meal.Updated_at})
			upsert := true
			opt := options.UpdateOptions{
				Upsert: &upsert,
            }
			result, err := mealCollection.UpdateOne(
				ctx,
				filter,
				bson.D{
					{"$set", updateObj},
                },
				&opt,
            )
			if err != nil{
				msg := "Comida no actualizada"
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            }
			defer cancel()
			c.JSON(http.StatusOK, result)
        }
	}
}


//NOTE: la funcion to Fixed sirve para redondear los numeros decimales a un numero especifico de decimales

//NOTE primitive.D es un tipo de dato que se utiliza para poder actualizar solo los campos que se envien en el body

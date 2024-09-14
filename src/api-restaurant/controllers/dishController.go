package controllers

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/Andres29AC/Proyectos-Go/src/api-restaurant/database"
	"github.com/Andres29AC/Proyectos-Go/src/api-restaurant/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var dishCollection *mongo.Collection = database.OpenCollection(database.Client, "dish")
var mealCollection *mongo.Collection = database.OpenCollection(database.Client, "meal")
var validate = validator.New()

func GetDishes() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		recordPerPage ,err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
        }

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
        }
		startIndex := (page - 1) * recordPerPage
		startIndex ,err = strconv.Atoi(c.Query("startIndex"))

		mathStage := bson.D{{"$math", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{{"_id", "null"}, {"total_count", bson.D{{"$sum", 1}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}}
		projectStage := bson.D{
			{
				"$project", bson.D{
					{"_id", 0},
					{"total_count", 1},
					{"dish_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
                },
            },
        }
		result, err := dishCollection.Aggregate(ctx, mongo.Pipeline{mathStage, groupStage, projectStage})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error ocurred while listing dishes items"})
        }
		var allDishes []bson.M
		if err = result.All(ctx, &allDishes); err != nil{
			log.Fatal(err)
        }
		c.JSON(http.StatusOK, allDishes[0])
	}
}

func GetDish() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx ,cancel = context.WithTimeout(context.Background(), 100*time.Second)
		dishId := c.Param("dish_id")
		var dish models.Dish
		err := dishCollection.FindOne(ctx, bson.M{"dish_id": dishId}).Decode(&dish)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error ocurrido mientras se buscaba el item del plato"})
       }
	   c.JSON(http.StatusOK, dish)
	}
}

func CreateDish() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	    var meal models.Meal
		var dish models.Dish
		if err := c.BindJSON(&dish); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
        }
		validationErr := validate.Struct(dish)
		if validationErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
        }
		err := mealCollection.FindOne(ctx, bson.M{"meal_id": dish.Meal_id}).Decode(&meal)
		defer cancel()
		if err != nil{
			msg := fmt.Sprintf("Meal was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
        }
		dish.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		dish.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		dish.ID = primitive.NewObjectID()
		dish.Dish_id = dish.ID.Hex()
		var num = toFixed(*dish.Price, 2)
		dish.Price = &num
		result ,insertErr := dishCollection.InsertOne(ctx, dish)
		if insertErr != nil{
			msg := fmt.Sprintf("Dish item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
        }
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
func round(num float64)int{
	return int(num + math.Copysign(0.5, num))
}
func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
func UpdateDish() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var meal models.Meal
		var dish models.Dish
		dishId := c.Param("dish_id")
		if err := c.BindJSON(&dish); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
        }
		var updateObj primitive.D
		if dish.Name != nil{
			updateObj = append(updateObj, bson.E{"name", dish.Name})
        }
		if dish.Price != nil{
			updateObj = append(updateObj, bson.E{"price", dish.Price})
        }
		if dish.Dish_image != nil{
			updateObj = append(updateObj, bson.E{"dish_image", dish.Dish_image})
        }
		if dish.Meal_id != nil{
			err := mealCollection.FindOne(ctx, bson.M{"meal_id": dish.Meal_id}).Decode(&meal)
			defer cancel()
			if err != nil{
				msg := fmt.Sprintf("Message: Meal was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            }
        }
		dish.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", dish.Updated_at})
		upsert := true
		filter := bson.M{"dish_id": dishId}
		opt := options.UpdateOptions{
			Upsert: &upsert,
        }
		result, err := dishCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
            },
			&opt,
        )
		if err != nil{
			msg := fmt.Sprint("Dish item was not updated")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
        }
		c.JSON(http.StatusOK, result)
	}
}

package controllers

import (
	"context"
	"fmt"
	"api-restaurant/database"
	"api-restaurant/models"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v10"
	"go.mongodb.org/mongo-driver/options"
	"go.mongodb.org/mongo-driver/mongo"
)
var dishCollection *mongo.Collection = database.OpenCollection(database.Client, "dish")
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
		startIndex ,err := strconv.Atoi(c.Query("startIndex"))

		mathStage := bson.D{{"$math", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"total_count", bson.D{{"$sum", 1}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}
		projectStage := bson.D{
			{
				"$project", bson.D{
					{"_id", 0},
					{"total_count", 1},
					{"dish_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
                }
            }
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

	}
}

func CreateDish() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}
func UpdateDish() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}

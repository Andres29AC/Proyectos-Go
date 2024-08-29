package models 

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dish struct {
	ID	    primitive.ObjectID `bson:"_id"` 
	Name    *string  `json:"name" validate:"required,min=2,max=100"`
	Price   *float64 `json:"price" validate:"required"`
	Meal_image *string `json:"meal_image" validate:"required"` 
	Created_at time.Time `json:"created_at"` 
	Updated_at time.Time `json:"updated_at"`
	Dish_id string `json:"dish_id"`
	Meal_id *string `json:"meal_id" validate:"required"`
}



//NOTE: omitempty sirve para que no se muestre el campo si no tiene valor


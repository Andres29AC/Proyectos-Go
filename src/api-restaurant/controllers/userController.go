package controllers


import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"github.com/Andres29AC/Proyectos-Go/src/api-restaurant/database"
	"github.com/Andres29AC/Proyectos-Go/src/api-restaurant/models"
	helper "github.com/Andres29AC/Proyectos-Go/src/api-restaurant/helpers"

)
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// func GetUsers() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
// 		if err != nil || recordPerPage < 1{
// 			recordPerPage = 10
// 		}
// 		page, err1 := strconv.Atoi(c.Query("page"))
// 		if err1 != nil || page < 1{
// 			page = 1
//         }
// 		startIndex := (page - 1) * recordPerPage
// 		startIndex, err = strconv.Atoi(c.Query("startIndex"))
// 		mathStage := bson.D{{"$math", bson.D{{}}}}
//         projectStage := bson.D{
// 			{"$project", bson.D{
// 				{"_id", 0},
// 				{"total_count", 1},
// 				{"user_items", bson.D{
// 					{"$slice", []interface{}{"$data", startIndex, recordPerPage}},
// 				}},
// 			}},
// 		}

// 		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{mathStage, projectStage})
// 		defer cancel()
// 		if err != nil{
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Se produjo un error al listar los elementos del usuario"})
//         }
// 		var allUsers []bson.M
// 		if err = result.All(ctx, &allUsers); err != nil{
// 			log.Fatal(err)
//         }
// 		c.JSON(http.StatusOK, allUsers[0])
// 	}
// }
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage

		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_items", bson.D{
					{"$slice", []interface{}{"$data", startIndex, recordPerPage}},
				}},
			}},
		}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{projectStage})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Se produjo un error al listar los elementos del usuario"})
			return
		}

		var allUsers []bson.M
		if err = result.All(ctx, &allUsers); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer los resultados"})
			return
		}

		if len(allUsers) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "No se encontraron usuarios"})
			return
		}

		c.JSON(http.StatusOK, allUsers[0])
	}
}


// func GetUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		userId := c.Param("user_id")
// 		var user models.User

// 		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)

// 		defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Se produjo un error al listar los elementos del usuario"})
// 		}
// 		c.JSON(http.StatusOK, user)
// 	}
// }
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		userId := c.Param("user_id")
		var user models.User

		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Se produjo un error al listar los elementos del usuario"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		//NOTE: Convierte los datos JSON que vienen de Postman en algo que Golang entiende.
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//NOTE: Validar los datos en funcion de la estructura del usuario.
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		//NOTE: Comprobaras si el correo electronico ya ha sido utilizado por otro usuario
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Se produjo un error al verificar el correo electronico."})
			return
		}
		//NOTE: Contraseña hash
		password := HashPassword(*user.Password)
		user.Password = &password

		//NOTE: Tambien comprobaras si el numero de telefono ya ha sido utilizado por otro usuario.
		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Se produjo un error al verificar el numero de telefono"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Este correo electronico o numero de telefono ya existe"})
			return
		}

		//NOTE: Crea algunos detalles adicionales para el objeto de usuario: created_at, updated_at, ID
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()

		//NOTE: Generar token y referir token (funcion de generacion de todos los tokens desde el asistente)
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		//NOTE: Si todo esta bien, inserte este nuevo usuario en la coleccion de usuarios.
		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("No se creo el usuario")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		//NOTE: Devolver el estado OK y enviar el resultado de vuelta
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}


func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found, login seems to be incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, foundUser.User_id)

		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		c.JSON(http.StatusOK, foundUser)
	}
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string,providedPassword string)(bool,string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("El nombre de usuario o la contraseña son incorrectos")
		check = false
	}
	return check, msg
}

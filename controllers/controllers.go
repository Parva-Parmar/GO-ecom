package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/Parva-Parmar/GO-ecom/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
) 

func HashPassword(password string) string{

}

func VerifyPassword(userPassword string , givenPassword string) (bool,string){

}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx , cancel = context.WithTimeout(context.Background(),100*time.Second)
		defer cancel()

		var user models.User 
		if err := c.BindJSON(&user); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error": validationErr})
		}

		count,err := UserCollection.CountDocuments(ctx,bson.M{"email": user.Email})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest,gin.H{"error":"user already exists"})
		}

		count,err = UserCollection.CountDocuments(ctx,bson.M{"phone":user.Phone})

		defer cancel()
		if err != nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest , gin.H{"error": "this phone no. is already in use"})
			return
		}
	}
}

func Login() gin.HandlerFunc{

}

func ProductViewerAdmin() gin.HandlerFunc{

}

func SearchProduct() gin.HandlerFunc{

}

func SearchProductByQuery() gin.HandlerFunc{

}


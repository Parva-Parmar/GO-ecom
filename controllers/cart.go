package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Parva-Parmar/GO-ecom/database"
	"github.com/Parva-Parmar/GO-ecom/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.Handler {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeOut(context.Background(), 5*time.Second)

		defer cancel()

		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError,err)
		}
		c.IndentedJSON(200, "Successfully added to the cart")
	}
}

func  (app *Application) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context){
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeOut(context.Background(), 5*time.Second)

		defer cancel()

		err = database.RemoveCartItem(ctx,app.prodCollection,app.userCollection,productID,userQueryID)

		if err != nil{
			c.IndentedJSON(http.StatusInternalServerError,err)
			return
		}
		c.IndentedJSON(200,"Successfully removed item from cart")
	}
}

func GetItemFromCart() gin.HandlerFunc {
	 return func(c *gin.Context){
		user_id := c.Query("id")

		if user_id == ""{
			c.Header("Context-Type","application/json")
			c.JSON(http.StatusNotFound,gin.H{"error":"invalid id"})
			c.Abort()
			return
		}

		usert_id,_ := primitive.ObjectIDFromHex(user_id)

		var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)
		defer cancel()

		var filledcart models.User
		UserCollection.FindOne(ctx,bson.D{primitive.E{Key:"_id",Value:usert_id}}).Decode(&filledcart)

	 }
}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context){
		userQueryID := c.Query("id")

		if userQueryID == ""{
			log.Panic("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest,errors.New("UserID is empty"))
		}

		var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)

		defer cancel()

		err := database.BuyItemFromCart(ctx,app.userCollection,userQueryID)
		if err != nil{
			c.IndentedJSON(http.StatusInternalServerError,err)
		}
		c.IndentedJSON("successfully placed the order")
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context){
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeOut(context.Background(), 5*time.Second)

		defer cancel()

		err = database.InstantBuyer(ctx,app.prodCollection,app.userCollection,productID,userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError,err)
		}

		c.IndentedJSON(200,"successfully placed the order")
	}
}

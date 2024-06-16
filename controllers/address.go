package controllers

import (
	"github.com/Parva-Parmar/GO-ecom/models"
	"github.com/gin-gonic/gin"
) 


func AddAddress() gin.HandlerFunc{

}

func EditHomeAddress() gin.HandlerFunc{

}

func EditWorkAddress() gin.HandlerFunc{

}

func DeleteAddress() gin.HandlerFunc{
	return func(c *gin.Context){
		user_id := c.Query("id")

		if user_id == ""{
			c.Header("Content-Type","application/json")
			c.JSON(http.StatusNotFound,gin.H{"Error":"Invalid Search Index"})
			c.Abort()
			return 
		}

		addresses := make([]models.Address,0)
		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil{
			c.IndentedJSON(500,"Internal Server Error")
		} 
	}
}
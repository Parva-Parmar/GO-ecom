package token

import (
	"go/token"
	"log"
	"os"
	"time"

	"github.com/Parva-Parmar/GO-ecom/database"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
)


type SignedDetails struct{
	Email string
	First_Name string
	Last_Name string
	Uid string
	jwt.StandardClaims
}

var UserData *mongo.Collection = database.userData(database.Client,"Users")

var SECRET_KEY = os.Getenv("SECRET_KEY")

func TokenGenerator(email string,firstname string,lastname string,uid string)(signedtoken string,signedrefreshtoken string,err error){
	claims := &SignedDetails{
		Email: email,
		First_Name: firstname,
		Last_Name: lastname,
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreahclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token,err := jwt.NewWithClaims(jwt.SigningMethodHS256,claims).SignedString([]byte(SECRET_KEY))
	if err != nil{
		return "", "", err
	}

	refreshtoken,err := jwt.NewWithClaims(jwt.SigningMethodHS256,refreahclaims).SignedString([]byte(SECRET_KEY))
	if err != nil{
		log.Panic(err)
		return
	}
	return token,refreshtoken,err

}

func ValidateToken(){

}

func UpdateAllTokens(){

}
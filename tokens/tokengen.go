// package token

// import (
// 	"context"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/Parva-Parmar/GO-ecom/database"
// 	jwt "github.com/golang-jwt/jwt/v4"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )


// type SignedDetails struct{
// 	Email string
// 	First_Name string
// 	Last_Name string
// 	Uid string
// 	jwt.RegisteredClaims 
// }

// var UserData *mongo.Collection = database.UserData(database.Client,"Users")

// var SECRET_KEY = os.Getenv("SECRET_KEY")

// func TokenGenerator(email string, firstname string, lastname string, uid string) (signedtoken string, signedrefreshtoken string, err error) {
//     claims := &SignedDetails{
//         Email:      email,
//         First_Name: firstname,
//         Last_Name:  lastname,
//         Uid:        uid,
//         RegisteredClaims: jwt.RegisteredClaims{
//             ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
//         },
//     }

//     refreshClaims := &SignedDetails{
//         RegisteredClaims: jwt.RegisteredClaims{
//             ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)), // 7 days
//         },
//     }

//     token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
//     if err != nil {
//         return "", "", err
//     }

//     refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
//     if err != nil {
//         log.Panic(err)
//         return "", "", err
//     }

//     return token, refreshToken, nil
// }

// // func TokenGenerator(email string,firstname string,lastname string,uid string)(signedtoken string,signedrefreshtoken string,err error){
// // 	claims := &SignedDetails{
// // 		Email: email,
// // 		First_Name: firstname,
// // 		Last_Name: lastname,
// // 		Uid: uid,
// // 		StandardClaims: jwt.StandardClaims{
// // 			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
// // 		},
// // 	}

// // 	refreahclaims := &SignedDetails{
// // 		StandardClaims: jwt.StandardClaims{
// // 			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
// // 		},
// // 	}

// // 	token,err := jwt.NewWithClaims(jwt.SigningMethodHS256,claims).SignedString([]byte(SECRET_KEY))
// // 	if err != nil{
// // 		return "", "", err
// // 	}

// // 	refreshtoken,err := jwt.NewWithClaims(jwt.SigningMethodHS256,refreahclaims).SignedString([]byte(SECRET_KEY))
// // 	if err != nil{
// // 		log.Panic(err)
// // 		return
// // 	}
// // 	return token,refreshtoken,err

// // }

// func ValidateToken(signedtoken string)(claims *SignedDetails,msg string){
// 	token,err := jwt.ParseWithClaims(signedtoken,&SignedDetails{},func(token *jwt.Token)(interface{},error){
// 		return []byte(SECRET_KEY),nil
// 	})

// 	if err != nil {
// 		msg = err.Error()
// 		return
// 	}

// 	claims,ok := token.Claims.(*SignedDetails)
// 	if !ok {
// 		msg = "the token is invalid"
// 	}

// 	if claims.ExpiresAt.Time.Before(time.Now()){
// 		msg = "token is already expired"
// 		return
// 	}
// 	return claims,msg
// }

// func UpdateAllTokens(signedtoken string,signedrefreshtoken string, userid string){
// 	var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)

// 	var updateobj primitive.D

// 	updateobj = append(updateobj,bson.E{Key:"token",Value:signedtoken})
// 	updateobj = append(updateobj,bson.E{Key:"refresh_token",Value:signedrefreshtoken})
// 	updated_at,_ := time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
// 	updateobj = append(updateobj,bson.E{Key:"updatedat",Value:updated_at})

// 	upsert := true
// 	filter := bson.M{"user_id":userid}
// 	opt := options.UpdateOptions{
// 		Upsert: &upsert,

// 	}
// 	_,err := UserData.UpdateOne(ctx,filter,bson.D{
// 		{Key:"$set",Value:updateobj},
// 		},
// 	&opt)

// 	defer cancel()
// 	if err != nil{
// 		log.Panic(err)
// 		return
// 	}
// }


package token

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Parva-Parmar/GO-ecom/database"
	jwt "github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SignedDetails contains custom claims as well as registered claims.
type SignedDetails struct {
	Email      string
	First_Name string
	Last_Name  string
	Uid        string
	jwt.RegisteredClaims
}

// UserData is the MongoDB collection.
var UserData *mongo.Collection = database.UserData(database.Client, "Users")

// SECRET_KEY is the secret key used for signing tokens.
var SECRET_KEY = os.Getenv("SECRET_KEY")

// TokenGenerator generates and returns a signed token and a signed refresh token.
func TokenGenerator(email string, firstname string, lastname string, uid string) (signedtoken string, signedrefreshtoken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_Name: firstname,
		Last_Name:  lastname,
		Uid:        uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)), // 7 days
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}

	return token, refreshToken, nil
}

// ValidateToken validates the provided token and returns the claims.
func ValidateToken(signedtoken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "the token is invalid"
		return
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		msg = "token is already expired"
		return
	}
	return claims, ""
}

// UpdateAllTokens updates the tokens in the database for the specified user.
func UpdateAllTokens(signedtoken string, signedrefreshtoken string, userid string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{Key: "token", Value: signedtoken})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedrefreshtoken})
	updatedAt := time.Now()
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: updatedAt})

	upsert := true
	filter := bson.M{"user_id": userid}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := UserData.UpdateOne(ctx, filter, bson.D{
		{Key: "$set", Value: updateObj},
	}, &opt)

	if err != nil {
		log.Panic(err)
	}
}

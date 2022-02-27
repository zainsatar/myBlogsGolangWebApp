package controllers

import (
	"context"
	"fmt"
	"log"
	"main/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//write your database username and password use set in mongodb atlas
const connectionstring = "mongodb+srv://username:password@cluster0.cbczv.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

const dbName = "BlogSpot"
const collectionName = "Users"

//a pointer that will point at the required table of given dbname
var collection *mongo.Collection

//creating the connection with db
func init() {
	//client options
	clientOptions := options.Client().ApplyURI(connectionstring)
	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		collection = client.Database(dbName).Collection(collectionName)
		fmt.Println("collection instance is created and ready")
	}
}
func SignupController(s models.Signup) (bool, string) {
	email := s.Email
	filter := bson.M{"email": email}
	resultSet := collection.FindOne(context.Background(), filter)
	if resultSet.Err() == nil {
		return false, "This email is already in use"
	}
	_, err := collection.InsertOne(context.Background(), s)
	if err != nil {
		log.Fatal(err)
	}
	return true, "Signup Successful."
}
func LoginController(l models.Login) (bool, string, models.Person) {
	email := l.Email
	filter := bson.M{"email": email, "password": l.Password}
	resultSet := collection.FindOne(context.Background(), filter)
	var userInfo models.Person
	if resultSet.Err() != nil {
		return false, "Incorrect Email or Password.", userInfo
	}

	err := resultSet.Decode(&userInfo)
	if err != nil {
		return false, "An error has occured.", userInfo
	}
	return true, "LoginSuccessful", userInfo

}

func UpdateProfileController(email string, password string, profile models.Signup) (bool, string) {
	if email != profile.Email {
		filter1 := bson.M{"email": profile.Email}
		resultset1 := collection.FindOne(context.Background(), filter1)
		if resultset1.Err() == nil {
			return false, "This email is already in use"
		}

	}
	filter2 := bson.M{"email": email, "password": password}
	// resultset := collection.FindOneAndUpdate(context.Background(), filter2, update)
	_, err := collection.ReplaceOne(context.Background(), filter2, profile)
	if err != nil {
		log.Fatal(err)
		return false, "Password is incorrent"
	}
	return true, "Profile Updated Successfully."
}

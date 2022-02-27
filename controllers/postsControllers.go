package controllers

import (
	"context"
	"fmt"
	"log"
	"main/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
//write your database username and password use set in mongodb atlas
const constring = "mongodb+srv://username:password@cluster0.cbczv.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

const dbname = "BlogSpot"
const collectionname = "Posts"

//a pointer that will point at the required table of given dbname
var collec *mongo.Collection

//creating the connection with db
func init() {
	//client options
	clientOptions := options.Client().ApplyURI(constring)
	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		collec = client.Database(dbname).Collection(collectionname)

		fmt.Println("collection instance is created and ready")
	}
}
func AddNewPostController(post models.Post) (bool, string) {
	_, err := collec.InsertOne(context.Background(), post)
	if err != nil {
		return false, "An error has occured."
	}
	return true, "Posted successfully."
}
func GetAllPosts() []primitive.M {
	resultSet, err := collec.Find(context.Background(), bson.D{{}}) //returns a result set of type []bson.M
	var posts []primitive.M
	if err != nil {
		log.Fatal(err)
	} else {
		for resultSet.Next(context.Background()) {
			var post primitive.M
			err := resultSet.Decode(&post)
			if err != nil {
				log.Fatal(err)
			} else {
				posts = append(posts, post)
			}
		}

	}
	defer resultSet.Close(context.Background())
	return posts

}

func GetMyPosts(id string) []primitive.M {
	filter := bson.M{"authorid": id}
	resultSet, err := collec.Find(context.Background(), filter) //returns a result set of type []bson.M
	var posts []primitive.M
	if err != nil {
		log.Fatal(err)
	} else {
		for resultSet.Next(context.Background()) {
			var post primitive.M
			err := resultSet.Decode(&post)
			if err != nil {
				// fmt.Println("error has occcures inside GetMyPosts()")
				log.Fatal(err)
			} else {
				posts = append(posts, post)
			}
		}

	}
	defer resultSet.Close(context.Background())
	return posts

}

package controllers

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const db = "BlogSpot"
const fileCollectionName = "fs.files"

//a pointer that will point at the required table of given dbname
var fileCollection *mongo.Collection

//creating the connection with db
func init() {
	//client options
	clientOptions := options.Client().ApplyURI(constring)
	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		fileCollection = client.Database(db).Collection(fileCollectionName)
		fmt.Println("collection instance is created and ready")
	}
}

func RetrieveImage(fileName string) string {
	filter := bson.M{"filename": fileName}
	result := fileCollection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return ""
	}
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(fileName, &buf)
	if err != nil {
		fmt.Println("error occurs here in download stream")
		log.Fatal(err)
	}
	fmt.Printf("File size to download: %v \n", dStream)
	base64Data := base64.StdEncoding.EncodeToString([]byte(buf.Bytes()))
	return base64Data
}

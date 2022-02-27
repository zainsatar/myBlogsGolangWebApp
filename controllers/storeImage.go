package controllers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var bucket *gridfs.Bucket
var collectionHandle *mongo.Collection
var cname = "fs.files"

func init() {
	//client options
	clientOptions := options.Client().ApplyURI(constring)
	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	collectionHandle = client.Database(dbname).Collection(cname)
	if err != nil {
		fmt.Println("client not found")
		log.Fatal(err)
	} else {
		// return client
		bucket, err = gridfs.NewBucket(client.Database("BlogSpot"))
		if err != nil {
			fmt.Println("bucket not found")
			log.Fatal(err)
		}
	}
}
func UploadFile(data []byte, filename string) {
	fmt.Println("data and filename received")
	fmt.Println("filename:", filename)
	filter := bson.M{"filename": filename}
	result, err := collectionHandle.Find(context.Background(), filter)
	fmt.Println("result is:", result)
	if err == nil { //image found
		fmt.Println("images found")
		var images []bson.M
		err := result.All(context.Background(), &images)
		if err != nil {
			fmt.Println("error occured while decoding image metadata")
			log.Fatal(err)
		}
		for _, val := range images {
			image := val
			id := image["_id"]
			fmt.Println("file id is:", id)
			eror := bucket.Delete(id)
			if eror != nil {
				fmt.Println("an error has occured while deleting the image")
			}
			fmt.Println("image deleted")
		}
	} else {
		fmt.Println(err)
	}
	uploadStream, err := bucket.OpenUploadStream(filename) // this is the name of the file which will be saved in the database
	if err != nil {
		fmt.Println("can not open the upload stream")
		log.Fatal(err)
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(data)
	if err != nil {
		fmt.Println("can not write data to the upload stream")
		log.Fatal(err)
	}
	fmt.Printf("Write file to DB was successful. File size: %d \n", fileSize)
}

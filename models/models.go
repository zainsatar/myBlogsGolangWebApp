package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Signup struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	Bio       string `json:"bio" bson:"bio"`
}
type UserDetails struct {
	Picture   string
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	Bio       string `json:"bio" bson:"bio"`
}
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname" bson:"firstname"`
	LastName  string             `json:"lastname" bson:"lastname"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password,omitempty" bson:"password,,omitempty"`
	Bio       string             `json:"bio" bson:"bio"`
}

type Login struct {
	Email    string `json:"email,,omitempty" bson:"email,,omitempty"`
	Password string `json:"password,,omitempty" bson:"password,,omitempty"`
}

type Post struct {
	Author   string `json:"author" bson:"author"`
	Authorid string `json:"authorid" bson:"authorid"`
	Title    string `json:"title" bson:"title"`
	Content  string `json:"content" bson:"content"`
	Date     string `json:"date" bson:"date"`
}

type Home struct {
	Picture   string
	FirstName string
	LastName  string
	Bio       string
	POSTS     []primitive.M
}

// type Image struct {
// 	ID         string    `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Length     int64     `json:"length" bson:"length"`
// 	ChunkSize  int32     `json:"chunkSize" bson:"chunkSize"`
// 	UploadDate time.Time `json:"uploadDate" bson:"uploadDate"`
// 	Filename   string    `json:"filename" bson:"filename"`
// }

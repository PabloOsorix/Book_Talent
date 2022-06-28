// Package engine provides functions to conect with a
// service of mongodb database, also  provides a CRUD
// functions.
package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx = context.TODO()

// User type is a struct that provides an architecture
// that allow us cast from bson(format of Mongodb) to json
// and vice versa.
type User struct {
	ObjectID   primitive.ObjectID `bson:"_id" json:"_id"`
	Name       string             `json: "name" bson: "name"`
	Profession string             `json: "profession" bson: "professsion"`
	Education  []string           `json: "education" bson: "education`
	Experience []string           `json: "experience" bson: "experience"`
	Years_exp  int                `json:  "years_exp" bson: "years_exp"`
	Languajes  string             `json: "languajes" bson: "languajes"`
	Residence  string             `json: "residence" bson: "residence"`
	Image      string             `json: "image" bson: "image"`
	Link       string             `json: "link" bson: "link"`
}

type Userer interface {
	Init()
}

func (user *User) Init() {
	user.ObjectID = primitive.NewObjectID()
	user.Name = ""
	user.Profession = ""
	user.Education = append(user.Education, "")
	user.Experience = append(user.Experience, "")
	user.Years_exp = 0
	user.Languajes = ""
	user.Residence = ""
	user.Image = ""
	user.Link = ""
}

// Create - function that creates a new connection with the database
/*
 Return: return a pointer to a Client session with mongodb.
*/
func Create() *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	//defer func() {
	//	if err := client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	return client
}

// New - function that creaate a new register in the database
/*
 (*mongo-Collection) coll = Pointer to the user collection in the
 database.
 (type User)user = new object of type User with all infomation of the
 new user to add iin the database.
 return: Return the object in format Byte
*/
func New(coll *mongo.Collection, user User) []byte {
	if _, err := coll.InsertOne(ctx, user); err != nil {
		panic(err)
	}
	response, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		panic(err)
	}
	return response
}

// GetAll - funtion to return all documents in a database.
/*
 (*mongo.Collection) coll = pointer to a collection in a mongo database
 (error) err = in success == nil otherwise is error.
 return: a slice of type User []User and error
*/
func GetAll(coll *mongo.Collection) []User {
	var result []User
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}

	for cursor.Next(ctx) {
		var element User
		if err := cursor.Decode(&element); err != nil {
			log.Fatal(err)
		}
		result = append(result, element)
	}
	if err != nil {
		panic(err)
	}
	return result
}

func getOne(moviesColl *mongo.Collection, link string) User {
	var result User
	var err error
	err = moviesColl.FindOne(ctx, bson.D{{"link", link}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the link %s\n", link)
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	//jsonData, err := json.MarshalIndent(result, "", "	")
	//if err != nil {
	//	panic(err)
	//}
	return result
}

// Update - function to update a document inside of database, return
// document updated.
/*
 (*mongo.Collection)coll = pointer to collection in database.
 (string) link = link of the user to edit.
 (User) user_updates = updates to be performed on the given user
 return: document updated type User
*/
func Update(coll *mongo.Collection, link string, user_updates User) string {
	var result User
	result = getOne(coll, link)
	user_updates.ObjectID = result.ObjectID
	if link == "" {
		log.Fatal("link is missing")
	}
	filter := bson.M{"link": link}
	if _, err := coll.ReplaceOne(ctx, filter, user_updates); err != nil {
		panic(err)
	}
	return "Update successfull"
}

// Delete - function to delete a document of the mongo database by its _id
/*
 (*mongo.Collection) coll = pointer to a collection of the database.
 (string) id = id of element to delete
 return: number of documents delete otherwise 0.
*/
func Delete(coll *mongo.Collection, link string) string {
	filter := bson.M{"link": link}
	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		panic(err)
	}
	return "{}"
}

// Disconnect - function to close connection with mongo database.
/*
 (*mongo.Client) coll = pointer to variable that contains currently session.
 return: In success error otherwise nil.
*/
func Disconnect(client *mongo.Client) error {
	if err := client.Disconnect(context.Background()); err != nil {
		panic(err)
		return err
	}

	return nil
}

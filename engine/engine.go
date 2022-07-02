// Package engine provides functions to conect with a
// service of mongodb database, also  provides a CRUD
// functions.
package engine

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/PabloOsorix/Book_Talent/user_model"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx = context.TODO()
var USER = goEnvVariable("USER_DB")
var PWD = goEnvVariable("USER_PWD")
var DATABASE = goEnvVariable("DATABASE")
type User = user_model.User



// Create - function that creates a new connection with the database
/*
 Return: return a pointer to a Client session with mongodb.
*/
func Create() (*mongo.Client, error) {
	if USER == "" || PWD == "" {
		log.Fatal("Missing database User or Password")
	}
	url := fmt.Sprintf(
		"mongodb+srv://%v:%v@%v.catis.mongodb.net/?retryWrites=true&w=majority", USER, PWD, DATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, errors.New(err.Error())
	}
	return client, nil
}

// Collection - function to connect with collection users in database
/*
 (*mongo.Client) client - pointer to the sopen session with mongodb
 return: pointer to a collection in db in success or error in case of
 fail
*/
func Collection(client *mongo.Client) (*mongo.Collection, error) {
	usersColl := client.Database("booktalent").Collection("users")
	if usersColl == nil {
		return nil, errors.New("Collection not found")
	}
	return usersColl, nil
}

// New - function that creaate a new register in the database
/*
 (*mongo-Collection) coll = Pointer to the user collection in the
 database.
 (type User)user = new object of type User with all infomation of the
 new user to add iin the database.
 return: Return the object in format Byte
*/
func New(coll *mongo.Collection, user User) (string, error) {
	if _, err := coll.InsertOne(ctx, user); err != nil {
		return "", errors.New(err.Error())
	}
	//response, err := json.MarshalIndent(user, "", "  ")
	//if err != nil {
	//	panic(err)
	//}
	return user.Name, nil
}

// GetAll - funtion to return all documents in a database.
/*
 (*mongo.Collection) coll = pointer to a collection in a mongo database
 (error) err = in success == nil otherwise is error.
 return: a slice of type User []User and error
*/
func GetAll(coll *mongo.Collection) ([]User, error) {
	var result []User
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}

	for cursor.Next(ctx) {
		var element User
		if err := cursor.Decode(&element); err != nil {
			return nil, errors.New(err.Error())
		}
		result = append(result, element)
	}
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return result, nil
}

// Update - function to update a document inside of database, return
// document updated.
/*
 (*mongo.Collection)coll = pointer to collection in database.
 (string) link = link of the user to edit.
 (User) user_updates = updates to be performed on the given user
 return: document updated type User
*/
func Update(coll *mongo.Collection, link string, user_updates User) (string, error) {
	
	if link == "" {
		log.Fatal("link is missing")
	}
	filter := bson.M{"link": link}

	result, err := getID(coll, link)
	if err != nil {
		return "", errors.New(err.Error())
	}

	user_updates.ObjectID = result
	_, err = coll.DeleteOne(ctx, filter)
	if err != nil {
		return "Record not Found!", errors.New(err.Error())
	}

	if _, err := coll.InsertOne(ctx, user_updates); err != nil {
		return "Record could'n update", errors.New(err.Error())
	}
	return "Update successfull", nil
}

// Delete - function to delete a document of the mongo database by its _id
/*
 (*mongo.Collection) coll = pointer to a collection of the database.
 (string) id = id of element to delete
 return: number of documents delete otherwise 0.
*/
func Delete(coll *mongo.Collection, link string) (string, error) {
	filter := bson.M{"link": link}
	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return "", err
	}
	return "record was deleted {}", nil
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



/*--------LOCAL PACKAGE FUNCTIONS-----------*/


// getID - Local package function to obtain the id of an existing user in
// the database. (is used by Update function)
/*
 (*mongo.Collection) coll = pointer to the collection users of the database
 (string) link = link to find the user to update
 return: primitive.ObjectID = id of existing user, in case of fail return
 an error
*/
func getID(coll *mongo.Collection, link string) (primitive.ObjectID, error) {
	var result User
	var err error
	err = coll.FindOne(ctx, bson.D{{"link", link}}).Decode(&result)
	if err != nil {
		log.Fatal(errors.New(err.Error()))
	}
	//jsonData, err := json.MarshalIndent(result, "", "	")
	//if err != nil {
	//	panic(err)
	//}
	return result.ObjectID, nil
}

// goEnvVariable - Local package function to obtain 
// enviroment variables
/*
 (string) key - variable to find in the file .env
 return: key in success.
*/
func goEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

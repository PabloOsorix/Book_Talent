package user_model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User type is a struct that provides an architecture
// that allow us cast from bson(format of Mongodb) to json
// and vice versa.
type User struct {
	ObjectID   primitive.ObjectID `bson:"_id" json:"_id"`
	Name       string             `json: "name" bson: "name"`
	Profession string             `json: "profession" bson: "professsion"`
	Education  []Education        `json: "education" bson: "education"`
	Experience []Experience       `json: "experience" bson: "experience"`
	Years_exp  int                `json: "years_exp" bson: "years_exp"`
	Languajes  string             `json: "languajes" bson: "languajes"`
	Residence  string             `json: "residence" bson: "residence"`
	Image      string             `json: "image" bson: "image"`
	Link       string             `json: "link" bson: "link"`
}

type Userer interface {
	Init()
}

type Education struct {
	Collague string `json: "collague" bson: "collague"`
	Title    string `json: "title" bson: "title"`
	Period   string `json: "period" bson: "period"`
}
type Experience struct {
	Title   string `json: "title" bson: "title"`
	Company string `json: "company" bson: "company"`
	Time    string `json: "time" bson: "time"`
}

// Init - function to initialize structure of type User
func (user *User) Init() {
	user.ObjectID = primitive.NewObjectID()
	user.Name = ""
	user.Profession = ""
	user.Years_exp = 0
	user.Languajes = ""
	user.Residence = ""
	user.Image = ""
	user.Link = ""
}

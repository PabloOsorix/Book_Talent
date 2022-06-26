package main

import (
	"fmt"
	"engine"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	
	type User = engine.User
	var new_user User
	//var result User
	var users []User

	client, err := engine.Create()
	if  err == nil {
		fmt.Println("successful connection!\n")
	}
	var usersColl = client.Database("myNewDataBase").Collection("users")


	//result = engine.Update(usersColl, "Thor: Ragnarok", "title", "Thor: Ragnarok1")
	//fmt.Println(result)


	//sucessful := engine.Delete(usersColl, "")
	//fmt.Println(sucessful)
	//users, err = engine.GetAll(usersColl)
	//if err != nil {
	//	panic(err)
	//}

	//new_user.Init()
	new_user.Name = "John Freddy Vega"
	new_user.Profession  = "CEO and Founder at Platzi "
	new_user.Education  = append(new_user.Education, 
		`"Logotipo de Stanford University Graduate School of Business
		Stanford University Graduate School of BusinessStanford University Graduate School of Business
		Endeavor Innovation & Growth, Business Administration and Management, GeneralEndeavor Innovation & Growth,
		Business Administration and Management, General 2018 - 2018"`)
	new_user.Experience = append(new_user.Experience,
		`CEO Platzi ene. 2014 - actualidad · 8 años 6 meses United States`)
	new_user.Years_exp = 8
	new_user.Languajes = "English, Spanish"
	new_user.Residence = "San Francisco, California, Estados Unidos"
	new_user.Image = `https://media-exp2.licdn.com/dms/image/C5603AQG0Z3gvXe1LFw/profile-displayphoto-shrink_800_800/0/1645721408804?e=1661385600&v=beta&t=SZUpmt5S0Kfrvc_C0oZVpSxRYdee1-bMoN1pA3pvAQk`
	new_user.Link = "https://www.linkedin.com/in/johnfreddyvega/2.0"
	
	users = append(users, new_user)
	//result := engine.Add(usersColl, users)
	result := engine.Update(usersColl, "https://www.linkedin.com/in/johnfreddyvega/2.0", new_user)
	//result := engine.Delete(usersColl, "https://www.linkedin.com/in/johnfreddyvega/")
	fmt.Println(string(result))
	

	//users = engine.GetAll(usersColl)
	//fmt.Println(users)

/*
	users, err = engine.GetAll(usersColl)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
	fmt.Println(users)
	for _, value := range users{
		fmt.Printf("Title: %v\tID: %v\n", value.Title, value.ID)
	}
*/	
}

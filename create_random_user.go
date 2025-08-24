package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"go-mysql-crud/database"
	"go-mysql-crud/models"
)

var firstNames = []string{"Manish", "Amit", "Priya", "Suresh", "Neha", "Rajesh", "Anita"}
var lastNames = []string{"Patel", "Shah", "Kumar", "Singh", "Reddy", "Verma", "Gupta"}
var domains = []string{"example.com", "testmail.com", "myapp.io", "golang.dev"}

func randomUser() models.User {
	// pick random name
	first := firstNames[rand.Intn(len(firstNames))]
	last := lastNames[rand.Intn(len(lastNames))]
	name := fmt.Sprintf("%s %s", first, last)

	// generate email
	email := fmt.Sprintf("%s.%s%d@%s",
		first, last, rand.Intn(1000), domains[rand.Intn(len(domains))])

	return models.User{
		Name:  name,
		Email: email,
	}
}

func main() {
	// seed randomness
	rand.Seed(time.Now().UnixNano())

	// connect DB
	database.Connect()

	// run loop - create 10 users
	for i := 0; i < 10; i++ {
		user := randomUser()
		if err := database.DB.Create(&user).Error; err != nil {
			log.Printf("❌ Failed to insert user: %v\n", err)
		} else {
			log.Printf("✅ Inserted user: %s (%s)\n", user.Name, user.Email)
		}
		time.Sleep(1 * time.Second) // small delay
	}
}

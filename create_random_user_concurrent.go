package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"go-mysql-crud/database"
	"go-mysql-crud/models"
)

// random string generator
func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// random email generator
func randomEmail() string {
	domains := []string{"gmail.com", "yahoo.com", "example.com"}
	return fmt.Sprintf("%s@%s", randomString(6), domains[rand.Intn(len(domains))])
}

func createUser(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	user := models.User{
		Name:  fmt.Sprintf("User_%s", randomString(5)),
		Email: randomEmail(),
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		log.Printf("[Worker %d] Failed to insert user: %v\n", id, result.Error)
		return
	}

	log.Printf("[Worker %d] Created user: %s (%s)\n", id, user.Name, user.Email)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	database.Connect()

	var wg sync.WaitGroup
	numUsers := 10 // create 10 users concurrently

	wg.Add(numUsers)
	for i := 1; i <= numUsers; i++ {
		go createUser(&wg, i)
	}

	wg.Wait()
	log.Println("✅ All users created concurrently.")
}

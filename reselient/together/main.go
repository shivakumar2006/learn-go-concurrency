package main

import (
	"context"
	"log"
)

func main() {
	client, err := NewMongoClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("appdb")
	cb := NewMongoBreaker()

	repo := NewUserRepository(db, cb)

	ctx := context.Background()

	user, err := repo.FindByID(ctx, "user_123")
	if err != nil {
		log.Println("Failed", err)
		return
	}

	log.Println("User found : ", user)
}

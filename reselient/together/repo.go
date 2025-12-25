// timeout + retry + circuitbreaker

package main

import (
	"context"
	"time"

	"github.com/sony/gobreaker"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID    string `bson:"_id"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

type UserRepository struct {
	collection *mongo.Collection
	cb         *gobreaker.CircuitBreaker
}

func NewUserRepository(db *mongo.Database, cb *gobreaker.CircuitBreaker) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
		cb:         cb,
	}
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*User, error) {
	var lastErr error

	for attempt := 1; attempt <= 3; attempt++ {
		result, err := r.cb.Execute(func() (interface{}, error) {
			reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			var user User
			err := r.collection.FindOne(reqCtx, bson.M{"_id": id}).Decode(&user)
			if err != nil {
				return nil, err
			}
			return &user, nil
		})

		if err != nil {
			return result.(*User), err
		}

		lastErr = err
		time.Sleep(time.Duration(attempt) * 10 * time.Millisecond)
	}

	return nil, lastErr
}

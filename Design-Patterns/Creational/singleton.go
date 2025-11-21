// package main

package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now")
			singleInstance = &single{}
		} else {
			fmt.Println("single instance already created")
		}
	} else {
		fmt.Println("single instance already created")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	fmt.Scanln()
}

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"sync"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type MongoDB struct {
// 	Client *mongo.Client
// }

// var instance *MongoDB
// var once sync.Once

// func GetMongoInstance() *MongoDB {
// 	once.Do(func() {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()

// 		clientsOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 		client, err := mongo.Connect(ctx, clientsOptions)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		// ping the database
// 		if err := client.Ping(ctx, nil); err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Println("MongoDB connection established!...")
// 		instance = &MongoDB{
// 			Client: client,
// 		}
// 	})
// 	return instance
// }

// func main() {
// 	// this first call creates a connection
// 	db1 := GetMongoInstance()

// 	// this second call calls the same connection
// 	db2 := GetMongoInstance()

// 	if db1 == db2 {
// 		fmt.Println("Both MongoDB instances are the same (Singleton works)")
// 	}

// 	// example usage: get a collection
// 	collection := db1.Client.Database("testdb").Collection("users")
// 	fmt.Println("MongoDB collections : ", collection.Name())
// }

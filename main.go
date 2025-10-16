package main

import (
    "context"
    "log"
    "net/http"
    "time"

    "github.com/Daniel4000-dev/mongo-golang/controllers"
    "github.com/julienschmidt/httprouter"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    // "go.mongodb.org/mongo-driver/mongo/readpref"
)


func main() {
	// Create MongoDB client and connect
    client := connectMongo()
    defer func() {
        if err := client.Disconnect(context.Background()); err != nil {
            log.Println("Error disconnecting MongoDB:", err)
        }
    }()

	r := httprouter.New()
	uc := controllers.NewUserController(client)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

// connectMongo sets up a MongoDB client and pings the server to verify connection.
func connectMongo() *mongo.Client {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // // Ping to check connection
    // if err := client.Ping(ctx, readpref.Primary()); err != nil {
    //     log.Fatalf("Failed to ping MongoDB: %v", err)
    // }

    return client
}

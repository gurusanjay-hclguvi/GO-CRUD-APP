package database
import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)
var Client *mongo.Client
var TodoCollection *mongo.Collection
func ConnectToMongoDB(uri string) error {
	ctx , cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	client , err := mongo.Connect(ctx,options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	Client = client
	TodoCollection = client.Database("todo_db").Collection("todos")
	return nil 
}

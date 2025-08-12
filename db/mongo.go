package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"LRProject3/config"
)

var Client *mongo.Client
var SessionCollection *mongo.Collection

func ConnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := config.GetEnv("MONGO_URI")
	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	SessionCollection = Client.Database("LRProject3").Collection("sessions")
	log.Println("Connected to MongoDB and session collection initialized")
}

func SaveSession(sessionID, token string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := SessionCollection.InsertOne(ctx, map[string]any{
		"session_id":   sessionID,
		"access_token": token,
		"expires_at":   time.Now().Add(24 * time.Hour),
	})
	if err != nil {
		log.Println("Error saving session:", err)
	}
}

func GetSessionToken(sessionID string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result map[string]any
	err := SessionCollection.FindOne(ctx, map[string]any{"session_id": sessionID}).Decode(&result)
	if err != nil {
		log.Println("Error fetching session token:", err)
		return ""
	}
	return result["access_token"].(string)
}

func DeleteSession(sessionID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := SessionCollection.DeleteOne(ctx, map[string]any{"session_id": sessionID})
	if err != nil {
		log.Println("Error deleting session:", err)
	}
}

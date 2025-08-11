package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var sessionCollection *mongo.Collection

func connectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	sessionCollection = client.Database("LRProject3").Collection("sessions")
}

func saveSession(sessionID, token string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := sessionCollection.InsertOne(ctx, map[string]any{
		"session_id":   sessionID,
		"access_token": token,
		"expires_at":   time.Now().Add(24 * time.Hour),
	})
	if err != nil {
		log.Println("Error saving session:", err)
	}
}

func getSessionToken(sessionID string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result map[string]any
	err := sessionCollection.FindOne(ctx, map[string]any{"session_id": sessionID}).Decode(&result)
	if err != nil {
		return ""
	}
	return result["access_token"].(string)
}

func deleteSession(sessionID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, _ = sessionCollection.DeleteOne(ctx, map[string]any{"session_id": sessionID})
}

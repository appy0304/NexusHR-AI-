package dao

import (
	"context"
	"time"

	"simple-go-api/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const chatHistoryCollection = "chat_history"

type ChatHistory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"userId"`
	Query     string             `bson:"query"`
	Answer    string             `bson:"answer"`
	Sources   []string           `bson:"sources"`
	CreatedAt time.Time          `bson:"createdAt"`
}

// SaveChat inserts a chat record into MongoDB
func SaveChat(chat *ChatHistory) error {
	col := config.MongoClient.Database(config.DB_NAME).Collection(chatHistoryCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	chat.CreatedAt = time.Now()
	_, err := col.InsertOne(ctx, chat)
	return err
}

// GetChatsByUser returns last N chats for a user (most recent first)
func GetChatsByUser(userID string, limit int) ([]ChatHistory, error) {
	col := config.MongoClient.Database(config.DB_NAME).Collection(chatHistoryCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}).SetLimit(int64(limit))
	cursor, err := col.Find(ctx, bson.M{"userId": userID}, opts)
	if err != nil {
		return nil, err
	}

	var chats []ChatHistory
	if err := cursor.All(ctx, &chats); err != nil {
		return nil, err
	}
	return chats, nil
}

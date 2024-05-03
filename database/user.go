package database

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	Password     string             `bson:"password"`
	Email        string             `bson:"email"`
	ProfileImage string             `json:"profile_image"`
}

func GetUserByID(userID primitive.ObjectID) (*User, error) {
	var user User
	collection := Db.Collection("users")
	filter := bson.D{{Key: "_id", Value: userID}}

	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		log.Println("Failed to fetch user:", err)
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	collection := Db.Collection("users")
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Println("Error retrieving user:", err)
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *User) error {
	collection := Db.Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println("Error creating user:", err)
	}
	return err
}

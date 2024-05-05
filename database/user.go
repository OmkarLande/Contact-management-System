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
	existingUser, err := GetUserByUsername(user.Username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("user with this username already exists")
	}

	existingUser, err = GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	collection := Db.Collection("users")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println("Error creating user:", err)
	}
	return err
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	collection := Db.Collection("users")
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Println("Error retrieving user by email:", err)
		return nil, err
	}
	return &user, nil
}

func UserExists(userID primitive.ObjectID) (bool, error) {
	collection := Db.Collection("users")
	filter := bson.D{{Key: "_id", Value: userID}}
	err := collection.FindOne(context.Background(), filter).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("User not found for ID:", userID)
			return false, nil
		}
		log.Println("Error checking user existence:", err)
		return false, err
	}

	return true, nil
}

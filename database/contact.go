package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id"`
	Name         string             `json:"name"`
	PhoneNumber  string             `json:"phoneNumber"`
	Email        string             `json:"email"`
	ProfileImage string             `json:"profile_image"`
}

func AddContact(contact Contact, userID primitive.ObjectID) error {
	contact.UserID = userID
	collection := Db.Collection("contacts")
	_, err := collection.InsertOne(context.Background(), contact)
	if err != nil {
		log.Println("Failed to insert contact:", err)
		return err
	}

	log.Println("Contact added successfully:")
	return nil
}

func GetContact(contactID, userID primitive.ObjectID) (Contact, error) {
	var contact Contact
	collection := Db.Collection("contacts")
	filter := bson.D{{Key: "_id", Value: contactID}, {Key: "user_id", Value: userID}}
	err := collection.FindOne(context.Background(), filter).Decode(&contact)
	if err != nil {
		return Contact{}, err
	}
	return contact, nil
}

func GetContacts(userID primitive.ObjectID) ([]Contact, error) {
	collection := Db.Collection("contacts")
	filter := bson.D{{Key: "user_id", Value: userID}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var contacts []Contact
	for cursor.Next(context.Background()) {
		var contact Contact
		if err := cursor.Decode(&contact); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return contacts, nil
}

func DeleteContact(contactID, userID primitive.ObjectID) error {
	collection := Db.Collection("contacts")
	filter := bson.D{{Key: "_id", Value: contactID}, {Key: "user_id", Value: userID}}
	_, err := collection.DeleteOne(context.Background(), filter)
	return err
}

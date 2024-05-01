package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type Contact struct {
	Name         string `json:"name"`
	PhoneNumber  string `json:"phoneNumber"`
	Email        string `json:"email"`
	ProfileImage string `json:"profile_image"`
}

func AddContact(contact Contact) error {
	collection := Db.Collection("contacts")
	_, err := collection.InsertOne(context.Background(), contact)
	if err != nil {
		log.Println("Failed to insert contact:", err)
		return err
	}

	log.Println("Contact added successfully:")
	return nil
}

func GetContacts() ([]Contact, error) {
	collection := Db.Collection("contacts")
	cursor, err := collection.Find(context.Background(), bson.M{})
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

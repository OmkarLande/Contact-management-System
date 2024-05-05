package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"contact-management-system/database"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EditContactHandler(w http.ResponseWriter, r *http.Request) {
	var updatedContact database.Contact
	err := json.NewDecoder(r.Body).Decode(&updatedContact)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if updatedContact.Name == "" && updatedContact.PhoneNumber == "" && updatedContact.Email == "" {
		log.Println("No fields to update")
		http.Error(w, "No fields to update", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	contactID, ok := params["contact_id"]
	if !ok {
		log.Println("Contact ID not found in request")
		http.Error(w, "Contact ID not found in request", http.StatusBadRequest)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(contactID)
	if err != nil {
		log.Println("Invalid contact ID:", err)
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	userID, ok := params["user_id"]
	if !ok {
		log.Println("User ID not found in request")
		http.Error(w, "User ID not found in request", http.StatusBadRequest)
		return
	}

	// Convert userID to primitive.ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println("Invalid user ID:", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	updatedContact.UserID = userObjectID

	avatarURL, err := GenerateRandomAvatar(updatedContact.Name)
	if err != nil {
		log.Println("Failed to generate avatar:", err)
		http.Error(w, "Failed to generate avatar", http.StatusInternalServerError)
		return
	}
	updatedContact.ProfileImage = avatarURL

	collection := database.Db.Collection("contacts")
	filter := bson.D{{Key: "_id", Value: objectID}}
	update := bson.D{{Key: "$set", Value: updatedContact}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Failed to update contact:", err)
		http.Error(w, "Failed to update contact", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Contact updated successfully"})
}

package handlers

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

	userID, ok := r.Context().Value("userID").(primitive.ObjectID)
	if !ok {
		log.Println("Failed to get user ID from request context")
		http.Error(w, "Failed to get user ID from request context", http.StatusInternalServerError)
		return
	}

	existingContact, err := database.GetContact(objectID, userID)
	if err != nil {
		log.Println("Failed to fetch existing contact:", err)
		http.Error(w, "Failed to fetch existing contact", http.StatusInternalServerError)
		return
	}

	if existingContact.UserID != userID {
		log.Println("Contact does not belong to the logged-in user")
		http.Error(w, "Contact does not belong to the logged-in user", http.StatusForbidden)
		return
	}

	if updatedContact.Name != "" {
		existingContact.Name = updatedContact.Name
	}
	if updatedContact.PhoneNumber != "" {
		existingContact.PhoneNumber = updatedContact.PhoneNumber
	}
	if updatedContact.Email != "" {
		existingContact.Email = updatedContact.Email
	}

	collection := database.Db.Collection("contacts")
	filter := bson.D{{Key: "_id", Value: objectID}}
	update := bson.D{{Key: "$set", Value: existingContact}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Failed to update contact:", err)
		http.Error(w, "Failed to update contact", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Contact updated successfully"})
}

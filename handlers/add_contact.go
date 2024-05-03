package handlers

import (
	"contact-management-system/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func generateRandomAvatar(name string) (string, error) {
	url := fmt.Sprintf("https://api.dicebear.com/8.x/bottts/svg?seed=%s", name)
	avatarURL := url
	return avatarURL, nil
}

func AddContactHandler(w http.ResponseWriter, r *http.Request) {
	var contact database.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Get the user ID from the request context
	userID, ok := r.Context().Value("userID").(primitive.ObjectID)
	if !ok {
		log.Println("Failed to get user ID from request context")
		http.Error(w, "Failed to get user ID from request context", http.StatusInternalServerError)
		return
	}

	avatarURL, err := generateRandomAvatar(contact.Name)
	if err != nil {
		log.Println("Failed to generate avatar:", err)
		http.Error(w, "Failed to generate avatar", http.StatusInternalServerError)
		return
	}

	contact.ProfileImage = avatarURL

	err = database.AddContact(contact, userID)
	if err != nil {
		log.Println("Failed to add contact to database:", err)
		http.Error(w, "Failed to add contact to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Contact added successfully"})
}

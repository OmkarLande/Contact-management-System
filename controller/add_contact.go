package controller

import (
	"contact-management-system/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GenerateRandomAvatar(name string) (string, error) {
	url := fmt.Sprintf("https://api.dicebear.com/8.x/bottts/svg?seed=%s", name)
	avatarURL := url
	return avatarURL, nil
}

func AddContactHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var contact database.Contact
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			log.Println("Failed to decode request body:", err)
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		userExists, err := database.UserExists(contact.UserID)
		if err != nil {
			log.Println("Failed to check user existence:", err)
			http.Error(w, "Failed to check user existence", http.StatusInternalServerError)
			return
		}
		if !userExists {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		avatarURL, err := GenerateRandomAvatar(contact.Name)
		if err != nil {
			log.Println("Failed to generate avatar:", err)
			http.Error(w, "Failed to generate avatar", http.StatusInternalServerError)
			return
		}
		contact.ProfileImage = avatarURL

		addedContact, err := database.AddContact(contact, contact.UserID)
		if err != nil {
			log.Println("Failed to add contact to database:", err)
			http.Error(w, "Failed to add contact to database", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		response, _ := json.Marshal(map[string]interface{}{"message": "Contact added successfully", "contactData": addedContact})
		w.Write(response)
	}
}

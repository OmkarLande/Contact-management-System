package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"contact-management-system/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ViewContactsHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(primitive.ObjectID)
	if !ok {
		log.Println("Failed to get user ID from request context")
		http.Error(w, "Failed to get user ID from request context", http.StatusInternalServerError)
		return
	}

	contacts, err := database.GetContacts(userID)
	if err != nil {
		log.Println("Failed to retrieve contacts from database:", err)
		http.Error(w, "Failed to retrieve contacts from database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

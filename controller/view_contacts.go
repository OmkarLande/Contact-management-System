package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"contact-management-system/database"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ViewContactsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDString := params["userID"]

	userID, err := primitive.ObjectIDFromHex(userIDString)
	if err != nil {
		log.Println("Invalid userID:", err)
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}

	userExists, err := database.UserExists(userID)
	if err != nil {
		log.Println("Failed to check user existence:", err)
		http.Error(w, "Failed to check user existence", http.StatusInternalServerError)
		return
	}
	if !userExists {
		http.Error(w, "User not found", http.StatusNotFound)
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

package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"contact-management-system/database"
)

func ViewContactsHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve contacts from the database
	contacts, err := database.GetContacts()
	if err != nil {
		log.Println("Failed to retrieve contacts from database:", err)
		http.Error(w, "Failed to retrieve contacts from database", http.StatusInternalServerError)
		return
	}

	// Respond with the retrieved contacts
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

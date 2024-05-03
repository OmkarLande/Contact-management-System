package handlers

import (
	"contact-management-system/database"
	"contact-management-system/middleware"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteContactHandler(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContextOrSession(r)

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

	err = database.DeleteContact(objectID, userID)
	if err != nil {
		log.Println("Failed to delete contact:", err)
		http.Error(w, "Failed to delete contact", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Contact deleted successfully"})
}

package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"contact-management-system/database"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("contact-management-system"))

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	user, err := database.GetUserByUsername(credentials.Username)
	if err != nil {
		log.Println("Failed to retrieve user:", err)
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}
	if user == nil {
		log.Println("User not found")
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		log.Println("Invalid password")
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	session, err := store.New(r, "contact-management-system-session")
	if err != nil {
		log.Println("Failed to create session:", err)
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	session.Values["userID"] = user.ID.Hex()

	err = session.Save(r, w)
	if err != nil {
		log.Println("Failed to save session:", err)
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Login successful", "user": user})
}

func LogoutUserHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		log.Println("Failed to get session:", err)
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		log.Println("Failed to delete session:", err)
		http.Error(w, "Failed to delete session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}

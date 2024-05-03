package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"contact-management-system/database"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if !isValidUsername(user.Username) {
		log.Println("UserName Should be unique and minimum 5 characters long.")
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}
	if !isValidPassword(user.Password) {
		log.Println("Password should be minimum 8 characters long.")
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}
	if !isValidEmail(user.Email) {
		log.Println("Email is not valid.")
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	existingUser, err := database.GetUserByUsername(user.Username)
	if err != nil {
		log.Println("Failed to check existing user:", err)
		http.Error(w, "Failed to check existing user", http.StatusInternalServerError)
		return
	}
	if existingUser != nil {
		log.Println("Username already exists")
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password:", err)
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)
	user.ProfileImage = fmt.Sprintf("https://api.dicebear.com/8.x/bottts/svg?seed=%s", user.Username)

	err = database.CreateUser(&user)
	if err != nil {
		log.Println("Failed to create user:", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func isValidUsername(username string) bool {
	if len(username) < 5 {
		return false
	}
	existingUser, err := database.GetUserByUsername(username)
	if err != nil {
		log.Println("Failed to check existing user:", err)
		return false
	}
	return existingUser == nil
}

func isValidPassword(password string) bool {
	return len(password) >= 8
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

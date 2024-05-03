package main

import (
	"contact-management-system/database"
	"contact-management-system/handlers"
	"contact-management-system/middleware"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	connectionString := "mongodb+srv://admin:AGXdEDgYfmZcmLJt@cluster0.h1aicec.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	database.ConnectToDatabase(connectionString)

	router := mux.NewRouter()

	authMiddleware := &middleware.AuthenticationMiddleware{}

	router.Use(authMiddleware.Skip("/login-user", "logout-user", "/register-user").Apply)

	router.HandleFunc("/add-contact", handlers.AddContactHandler).Methods("POST")
	router.HandleFunc("/view-contacts", handlers.ViewContactsHandler).Methods("GET")
	router.HandleFunc("/edit-contact/{contact_id}", handlers.EditContactHandler).Methods("PUT")
	router.HandleFunc("/delete-contact/{contact_id}", handlers.DeleteContactHandler).Methods("DELETE")
	router.HandleFunc("/register-user", handlers.RegisterUserHandler).Methods("POST")
	router.HandleFunc("/login-user", handlers.LoginUserHandler).Methods("POST")
	router.HandleFunc("/logout-user", handlers.LogoutUserHandler).Methods("POST")

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

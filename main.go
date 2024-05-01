package main

import (
	"contact-management-system/database"
	"contact-management-system/handlers"
	"log"
	"net/http"
)

func main() {
	connectionString := "mongodb+srv://admin:AGXdEDgYfmZcmLJt@cluster0.h1aicec.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	database.ConnectToDatabase(connectionString)

	http.HandleFunc("/add-contact", handlers.AddContactHandler)
	http.HandleFunc("/view-contacts", handlers.ViewContactsHandler)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

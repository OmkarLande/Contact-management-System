package main

import (
	"contact-management-system/controller"
	"contact-management-system/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	var mongodbURI = os.Getenv("MONGODB_URI")
	connectionString := mongodbURI
	database.ConnectToDatabase(connectionString)

	router := mux.NewRouter()

	router.HandleFunc("/add-contact", controller.AddContactHandler()).Methods("POST")

	router.HandleFunc("/view-contacts/{userID}", controller.ViewContactsHandler).Methods("GET")
	router.HandleFunc("/edit-contact/{contact_id}/{user_id}", controller.EditContactHandler).Methods("PUT")
	router.HandleFunc("/delete-contact/{contact_id}", controller.DeleteContactHandler).Methods("DELETE")
	router.HandleFunc("/register-user", controller.RegisterUserHandler).Methods("POST")
	router.HandleFunc("/login-user", controller.LoginUserHandler).Methods("POST")
	router.HandleFunc("/logout-user", controller.LogoutUserHandler).Methods("POST")

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(router))
}

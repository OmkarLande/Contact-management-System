package main

import (
	"github.com/OmkarLande/PRODIGY_SD_03/d"
)

func main() {
	// Connect to MongoDB
	connectionString := "mongodb+srv://admin:AGXdEDgYfmZcmLJt@cluster0.h1aicec.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	ConnectToDatabase(connectionString)
}

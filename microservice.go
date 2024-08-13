package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// User represents a user entity

type User struct {
	ID int `json:"id"`

	Username string `json:"username"`

	Email string `json:"email"`
}

var users = []User{

	{ID: 1, Username: "user1", Email: "user1@example.com"},

	{ID: 2, Username: "user2", Email: "user2@example.com"},
}

func main() {

	http.HandleFunc("/users", getUsers)

	log.Println("Server started on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {

		http.Error(w, "Failed to encode users", http.StatusInternalServerError)

		return

	}

}

//Creating a Simple Microservice Using HTTP
//By following the following steps, you can create a simple microservice using either HTTP or gRPC in Go. The choice between HTTP or gRPC depends on factors such as performance requirements, compatibility with existing systems, and developer preferences.
//
//Define Service Interface: Define the endpoints (HTTP methods and paths) that the microservice will expose. For example, /users for fetching user data.
//
//Implement Service Logic: Implement the service logic for each endpoint. This typically involves handling incoming HTTP requests, parsing request parameters, processing data, and generating HTTP responses.
//
//Add Middleware (Optional): Add middleware for cross-cutting concerns such as authentication, logging, error handling, etc.
//
//Set Up Server: Set up an HTTP server (e.g., using Go's http package) and register the service handlers for each endpoint.
//
//Start Server: Start the HTTP server to listen for incoming requests on a specified port.
//
//Create Client (Optional): If necessary, create a client to interact with the HTTP endpoints of the microservice.
//
//Test Service: Test the microservice endpoints using tools like curl, Postman, or automated tests.
//
//Deploy Service: Deploy the HTTP microservice to a server or cloud platform where it can be accessed by clients.
//
//Monitor and Maintain: Monitor the microservice in production to ensure uptime, performance, and reliability. Implement logging, monitoring, and alerting mechanisms to detect and respond to issues. Regularly update and maintain the microservice to address bugs, add new features, and improve performance.

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Health check
	router.HandleFunc("/api/health", healthCheck).Methods("GET")

	// Auth routes
	router.HandleFunc("/api/auth/register", registerUser).Methods("POST")
	router.HandleFunc("/api/auth/login", loginUser).Methods("POST")

	// Capsule tent routes
	router.HandleFunc("/api/tents", getTents).Methods("GET")
	router.HandleFunc("/api/tents/{id}", getTentByID).Methods("GET")
	router.HandleFunc("/api/tents", createTent).Methods("POST")

	// Booking routes
	router.HandleFunc("/api/bookings", createBooking).Methods("POST")
	router.HandleFunc("/api/bookings/{id}", getBooking).Methods("GET")
	router.HandleFunc("/api/bookings/user/{userID}", getUserBookings).Methods("GET")
	router.HandleFunc("/api/bookings/{id}/cancel", cancelBooking).Methods("PUT")

	// Payment routes
	router.HandleFunc("/api/payments/{bookingID}", processPayment).Methods("POST")
	router.HandleFunc("/api/payments/{paymentID}", getPaymentStatus).Methods("GET")

	// Guest routes
	router.HandleFunc("/api/guests", registerGuest).Methods("POST")
	router.HandleFunc("/api/guests/{id}", getGuestInfo).Methods("GET")

	port := ":8080"
	fmt.Printf("Server running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// Health check endpoint
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"OK","message":"Nashik Kumbh Tent Booking API is running"}`)
}

// Auth handlers
func registerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"message":"User registered successfully"}`)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message":"Login successful","token":"jwt_token_here"}`)
}

// Tent handlers
func getTents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"tents":[{"id":1,"name":"Deluxe Capsule","price":5000,"capacity":2}]}`)
}

func getTentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"id":1,"name":"Deluxe Capsule","price":5000,"capacity":2,"description":"Comfortable capsule tent"}`)
}

func createTent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"message":"Tent created successfully"}`)
}

// Booking handlers
func createBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"bookingID":"BOOK123","status":"pending","totalAmount":5000}`)
}

func getBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"bookingID":"BOOK123","tentID":1,"userID":1,"checkIn":"2026-05-21","checkOut":"2026-05-25","status":"confirmed"}`)
}

func getUserBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"bookings":[{"bookingID":"BOOK123","status":"confirmed"}]}`)
}

func cancelBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message":"Booking cancelled successfully"}`)
}

// Payment handlers
func processPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"paymentID":"PAY123","status":"success","message":"Payment processed"}`)
}

func getPaymentStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"paymentID":"PAY123","status":"completed","amount":5000}`)
}

// Guest handlers
func registerGuest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"guestID":"GUEST123","message":"Guest registered successfully"}`)
}

func getGuestInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"guestID":"GUEST123","name":"John Doe","email":"john@example.com","country":"USA"}`)
}

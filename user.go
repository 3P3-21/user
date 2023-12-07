package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

var db *sql.DB

func initDB() {
	var err error
	connStr := "user=username dbname=dbname sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

func main() {
	initDB()
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/register", RegisterHandler).Methods("POST")
	router.HandleFunc("/login", LoginHandler).Methods("POST")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	RoleID    int
	CreatedAt time.Time
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var user User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash the password", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO users (email, password_hash, first_name, last_name, role_id) VALUES ($1, $2, $3, $4, $5)",
		user.Email, string(hashedPassword), user.FirstName, user.LastName, user.RoleID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			http.Error(w, "Email already exists", http.StatusConflict)
		} else {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var credentials struct {
		Email    string
		Password string
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var storedPassword string
	err := db.QueryRow("SELECT password_hash FROM users WHERE email = $1", credentials.Email).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusUnauthorized)
		} else {
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(credentials.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	sessionID := "some-random-session-id"

	_, err = db.Exec("INSERT INTO sessions (session, expires_at) VALUES ($1, $2)", sessionID, time.Now().Add(time.Hour*24))
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"session_id": sessionID}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

package handler

import (
	"encoding/json"
	"net/http"

	"RestfulAPI/config"
)

// User model
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// GetAllUsers handler to retrieve all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := config.GetDB().Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUser handler to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = config.GetDB().Exec("INSERT INTO users (username, email) VALUES (?, ?)", newUser.Username, newUser.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

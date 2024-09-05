package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

// User model
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Handler function for /users route
func handleUsersPage(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Fetch users from the database
	users, err := fetchUsers(db)
	if err != nil {
		http.Error(w, "Unable to fetch users", http.StatusInternalServerError)
		return
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/users.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	// Render the template with the list of users
	err = tmpl.Execute(w, users)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

// Fetch users from the database
func fetchUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name FROM user_details")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users, nil
}

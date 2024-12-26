package user

import (
	companyDB "company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	db := companyDB.Validate(w)
	defer db.Close()

	var user structs.User
	json.NewDecoder(r.Body).Decode(&user)

	err := createUser(db, user)

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully!!!")
}

func createUser(db *sql.DB, user structs.User) error {

	password, error := companyDB.HashPassword(user.Password)

	if error != nil {
		panic("Error hash password")
	}

	query := "INSERT INTO user (id, name, username, password) values (?, ?, ?, ?)"
	_, err := db.Exec(query, 0, user.Name, user.Username, password)

	if err != nil {
		return err
	}

	return nil
}

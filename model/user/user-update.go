package user

import (
	companyDB "company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := companyDB.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	userId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Impossible change to int", http.StatusInternalServerError)
	}

	var user structs.User
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Impossible decode", http.StatusInternalServerError)
	}

	error := updateUserById(db, userId, user)
	if error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "User updated successfully")
}

func updateUserById(db *sql.DB, userId int, user structs.User) error {

	password, error := companyDB.HashPassword(user.Password)

	if error != nil {
		panic("Error hash password")
	}

	query := "update user set name = ?, username = ?, password = ? where id = ?"
	_, err := db.Exec(query, user.Name, user.Username, password, userId)

	if err != nil {
		return err
	}

	return nil
}

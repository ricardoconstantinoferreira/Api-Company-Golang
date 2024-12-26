package user

import (
	"company/db"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	userId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Error userId", http.StatusBadRequest)
		return
	}

	employee := deleteUserById(db, userId)

	if employee != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "User deleted successfully")
}

func deleteUserById(db *sql.DB, userId int) error {
	query := "delete from user where id = ?"
	_, err := db.Exec(query, userId)

	if err != nil {
		return err
	}

	return nil
}

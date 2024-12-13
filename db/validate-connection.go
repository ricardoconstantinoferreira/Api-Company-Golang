package db

import (
	"database/sql"
	"net/http"
)

func Validate(w http.ResponseWriter) *sql.DB {
	db, err := GetConnect()

	if err != nil {
		http.Error(w, "Failed in connection with database", http.StatusInternalServerError)
	}

	return db
}

package user

import (
	"company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetListUserHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	user, err := getListUser(db)

	if err != nil {
		http.Error(w, "No one user find", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	userId, err := strconv.Atoi(id)

	users, err := getUserById(db, userId)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUserById(db *sql.DB, id int) (*structs.User, error) {
	query := "select id, name, username, password from user where id = ?"

	result := db.QueryRow(query, id)

	user := &structs.User{}

	err := result.Scan(&user.Id, &user.Name, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func getListUser(db *sql.DB) (*map[int]structs.User, error) {
	results, err := db.Query("select id, name, username, password from user")

	if err != nil {
		panic(err.Error())
	}

	resultado := map[int]structs.User{}

	user := structs.User{}
	cont := 0

	for results.Next() {

		err = results.Scan(&user.Id, &user.Name, &user.Username, &user.Password)

		if err != nil {
			panic(err.Error())
		}

		resultado[cont] = user
		cont += 1
	}

	return &resultado, nil
}

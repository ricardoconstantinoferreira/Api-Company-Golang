package handler

import (
	"company/db"
	"company/model/employee"
	"company/structs"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createToken(userID, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var employeeSt structs.LoginEmployee

	json.NewDecoder(r.Body).Decode(&employeeSt)

	employee, err := employee.GetEmployeeByName(employeeSt.Username, w)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "invalid username or password", http.StatusUnauthorized)
		} else {
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	valid := db.VerifyPassword(employeeSt.Password, employee.Password)

	if !valid {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := createToken(strconv.Itoa(employee.Id), employee.Username)

	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})

}

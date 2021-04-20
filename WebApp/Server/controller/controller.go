package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mskKandula/middleware"
	"github.com/mskKandula/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	Db  *sql.DB
	err error
	// credentials = make(map[string]string)
)

func Signup(w http.ResponseWriter, r *http.Request) {

	var user model.User

	json.NewDecoder(r.Body).Decode(&user)

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword := string(hash)

	query, err := Db.Prepare("INSERT INTO USERS(name, age, email, phoneNo, password) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Panic(err)
	}
	query.Exec(user.Name, user.Age, user.Email, user.PhoneNo, hashedPassword)

	// credentials[user.Email] = user.Password

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)

}
func Login(w http.ResponseWriter, r *http.Request) {

	var (
		validation model.Validation
		password   string
	)

	json.NewDecoder(r.Body).Decode(&validation)

	// if credentials[validation.Email] != validation.Password {

	// 	w.WriteHeader(http.StatusUnauthorized)

	// 	return
	// }

	row := Db.QueryRow("select password from USERS where email=?", validation.Email)

	err = row.Scan(&password)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	// if validation.Password != password {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(validation.Password)); err != nil {

		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString, expirationTime, err := middleware.Auth(validation)

	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Token Expires at : ", expirationTime)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	fmt.Fprintf(w, tokenString)
}
func Logout(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		MaxAge: -1,
	})
}

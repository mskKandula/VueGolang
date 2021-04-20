package middleware

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mskKandula/model"
)

func Auth(creds model.Validation) (string, time.Time, error) {

	var err error

	//Creating Access Token
	os.Setenv("jwtKey", "7yt65U745TR57lo9h%$fre#$TR43EW") //this should be in an env file

	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true

	atClaims["email"] = creds.Email

	atClaims["password"] = creds.Password

	expirationTime := time.Now().Add(time.Minute * 5)

	atClaims["expireAt"] = expirationTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	tokenString, err := token.SignedString([]byte(os.Getenv("jwtKey")))

	if err != nil {
		return "", expirationTime, err
	}

	return tokenString, expirationTime, nil

}

func Decode(tokenString string) (interface{}, error) {

	// Initialize a new instance of `Claims`
	claims := jwt.MapClaims{}
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("jwtKey")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}

	if !token.Valid {
		return "", err
	}

	email := claims["email"]

	return email, nil
}

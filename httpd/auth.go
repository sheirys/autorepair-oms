package httpd

import (
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type JwtTokenResponse struct {
	Token string `json:"token"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	cred := Credentials{}

	// generate and sign token
	signed, err := app.generateToken(cred)
	if err != nil {
		logrus.WithError(err).Error("cannot generate token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(JwtTokenResponse{Token: signed})
}

// checkToken should be used in middleware to verify that jwt token is valid.
func (app *Application) checkToken(w http.ResponseWriter, r *http.Request) {
}

func (app *Application) generateToken(c Credentials) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": "bybis",
		"exp":      time.Now().Add(12 * time.Hour),
		"iat":      time.Now(),
		"nbf":      time.Now(),
	})
	return token.SignedString(app.SignSecret)
}

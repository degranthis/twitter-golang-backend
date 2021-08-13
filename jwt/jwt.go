package jwt

import (
	"time"

	"github.com/degranthis/twitter-golang-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GeneroJWT se usa para generar el jwt */
func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("MastersDesarrollo")

	payload := jwt.MapClaims{
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}

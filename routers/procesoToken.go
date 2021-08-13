package routers

import (
	"errors"
	"strings"

	"github.com/degranthis/twitter-golang-backend/bd"

	"github.com/degranthis/twitter-golang-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email */
var Email string

/*IDUsuario */
var IDUsuario string

/*ProcesoToken proceso token para extrar sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersDesarrollo")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token inválido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)

		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token inválido")
	}

	return claims, false, string(""), err

}

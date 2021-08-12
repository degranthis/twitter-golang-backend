package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la funcion para encriptar password */
func EncriptarPassword(pass string) (string, error) {

	salt := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), salt)

	return string(bytes), err

}

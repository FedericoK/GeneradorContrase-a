package main

import (
	//Generan numeros aleatorios seguros
	"crypto/rand"
	"strings"

	//se usan para manejar numeros grandes
	"math/big"
)

// Creamos una funcion para generar la contraseña que reciba un numero para el largo como parametro
// y devuelva un string y error.
// La funcion empiza con mayuscula para que sea publica a los otros archivos del paquete main.
func GenerarContra(length int) (string, error) {

	//Se define chars, contiene todos los caracteres que se pueden usar en la contraseña
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	//Variable password como un string builder para hacer la contraseña
	var password strings.Builder

	//un for que se repite tantas vece como la longitud de la contraseña deseada
	for i := 0; i < length; i++ {

		//Se crea la variable randomNum para elegir un caracter al azar de chars
		//ran.Int retorna un valor entre 0 y max donde max es la longitud de chars
		// rand reader shared instance of a cryptographically secure random number generator
		//big.NewInt aloca un nuevo int
		randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))

		//Verificamos que no tengamos ningun error
		if err != nil {
			//si error es distinto de 0 retornamos en blanco el passwrod y el error generado
			return "", err
		}

		//colocamos la letra de Chars en paswword
		password.WriteByte(chars[randomNum.Int64()])
	}
	//si no fallo la contraseña se retorna el password creado y error en 0
	return password.String(), nil
}

package main

import "fmt"

func main() {

	//declaramos la variable de la longitud
	var length int

	//Pedimos que se ingrese la longitud
	fmt.Print("Ingrese la longitud de la contraseña:")

	//Usamos la funcion fmt.Scan ya que solo vamos a necesita una cadena sin espacios
	_, err1 := fmt.Scan(&length)

	//Se verifica si lo que se ingresa es un numero
	if err1 != nil {
		//Si da error da mensaje de error y se termina
		fmt.Println("Error al leer la longitud. Por favor, ingrese un numero.")
		return
	}

	//Llamamos a la funcion GenerarContra para generara la contraseña
	//Ver generador.go
	password, err2 := GenerarContra(length)

	//Se verifica que no haya errores al generar la contraseña
	if err2 != nil {
		fmt.Println("Hubo un error al generar la contraseña: ", err2)
		return
	}

	//Imprimimos la contraseña
	fmt.Println("Contraseña generada: ", password)

}

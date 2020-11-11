package api_back

import (
	"fmt"
	"io/ioutil"
)

// Esta funcion se utiliza para identificar si el token a expirado
// se utiliza comparando el body de cada request, si son iguales entonces el token ha
// expirado por lo que es necesario realizar un refresh token
func tokenError () string {
	return "{\"cause\":[],\"message\":\"Invalid token\",\"error\":\"not_found\",\"status\":401}"
}

// Funcion para abrir los archivos json
func openJson(filename string) string {

	// lee el la plantilla .json y lo almacena en un string
	jsonData, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return ""
	}

	return  string(jsonData)
}


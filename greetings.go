package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especifica
func Hello(name string) (string, error) {
	//Verficar si name es vacio
	if name == "" {
		return "", errors.New("nombre vacio")
	}
	//Devuelve un saludo que incluye el nombre en un mensaje
	message := fmt.Sprintf(randomFormat(), name)
	//Probar test
	//message := fmt.Sprint(randomFormat())
	return message, nil
}

// funcion para retornar un mapa que asocia cada nombre con un mensaje de saludo
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// Funcion para devolver saludos aleatorios
func randomFormat() string {
	formats := []string{
		"Hola, %v! Bienvenido!",
		"Que bueno verte, %v!",
		"Saludos, %v! Encantado de conocerte!",
	}

	return formats[rand.Intn(len(formats))]
}

# Saludos en Go

Este paquete proporciona una forma simple de obtener saludos personalizados

## Instalacion
Ejectua el siguiente comando para instalar el paquete:
```bash
go get -u github.com/juanchosandox90/greetings.git
```

## Uso
Aqui tienes un ejemplo de como utilizar el paquete en tu codigo:

```go
package main

import (
    "fmt"
    "github.com/juanchosandox90/greetings"
)

func main(){
    message, err := greetings.Hello("John Doe")

    if err != nil {
        fmt.Println("Ocurrio un error:", err)
        return
    }

    fmt.Println(message)
}

```

Este ejemplo importa el paquete github.com/juanchosandox90/greetings y llama la funcion Hello para crear
un saludo personalizado. Si ocurre un error, se imprime un mensaje de error.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//obtiene la fecha actua en segundos unix
	segundos := time.Now().Unix()
	//generar un numero random con respecto a los segundos
	rand.Seed(segundos)
	//Generador de un número entero entre el 1 y el 100
	objetivo := rand.Intn(100) + 1
	//Imprimir en terminal lo que hace el sistema
	fmt.Println("Trata de adivinar seleccionar un número del 1 al 100")
	//instancia de lectura de datos en el terminal
	lector := bufio.NewReader(os.Stdin)
	//configurando para imprimir de forma predetrminada el mesaje de error
	terminado := false // ***************************** ":=" se declara variable y se asigna el valor. "llamado Short variable declaration"
	//For solo puede ejecutarce solo 10 veces
	for intentos := 0; intentos < 10; intentos++ {
		fmt.Println("Tienes: ", 10-intentos, " veces aún.")
		//comienza en juego, preguntando que trate de adivinar
		fmt.Println("Trata de adivinar: ")
		//realiza una lectura de un string hasta que el usuario de un enter
		input, err := lector.ReadString('\n')
		// validación si existe algun error durante la lectura de datos
		if err != nil {
			log.Fatal(err)
		}
		// remueve la linea que se agrego cuando se dio el enter
		input = strings.TrimSpace(input) // *********************** "=" se le asigna valor
		//convertir el string de ingreso a un interger
		intento, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		// Validacion del juego si es menor, mayor o lo adivinó
		if intento < objetivo {
			fmt.Println("Opps, tu número es mas bajo que el número random")
		} else if intento > objetivo {
			fmt.Println("Opps, tu número es más alto que el número random")
		} else {
			//previene el mensage de error
			terminado = true
			fmt.Println("Buen trabajo, ¡Adivinaste!")
			//sale del loop
			break
		}
	}
	//hace la validación de que la variavle succes este en el boleano true y si esta en falase mas de 10 veces
	if !terminado {
		fmt.Println("Lo siento, no adivinaste el número el cual era: ", objetivo)
	}
}

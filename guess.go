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

func main(){
	//obtiene la feha y el tiempo actual en tiempo unix
	seconds := time.Now().Unix()
	//generar un numero random con respecto a los segundos 
	rand.Seed(seconds)
	//Generador de un número entero entre el 1 y el 100
	target := rand.Intn(100) + 1
	//Imprimir en terminal lo que hace el sistema 
	fmt.Println("Trata de adivinar seleccionar un número del 1 al 100")
	//instancia de lectura de datos en el terminal 
	reader := bufio.NewReader(os.Stdin)
	//configurando para imprimir de forma predetrminada el mesaje de error
	success := false
	//For solo puede ejecutarce solo 10 veces 
	for guesses := 0; guesses < 10; guesses++{
		fmt.Println("Tienes: ", 10-guesses, " veces aún.")
		//comienza en juego, preguntando que trate de adivinar 
		fmt.Println("Trata de adivinar: ")
		//realiza una lectura de un string hasta que el usuario de un enter 
		input, err := reader.ReadString('\n')
		// validación si existe algun error durante la lectura de datos 
		if err != nil{
			log.Fatal(err)
		}
		// remueve la linea que se agrego cuando se dio el enter 
		input = strings.TrimSpace(input) // *************** investigar cual es la diferencia el igual y el igual := 
		//convertir el string de ingreso a un interger
		guess, err := strconv.Atoi(input)
		if err != nil{
			log.Fatal(err)
		}
		// Validacion del juego si es menor, mayor o lo adivinó 
		if guess < target{
			fmt.Println("Opps, tu número es mas bajo")
		}else if guess > target{
			fmt.Println("Opps, tu número es más alto")
		}else{
			//previene el mensage de error 
			success = true 
			fmt.Println("Buen trabajo, ¡Adivinaste!")
			//sale del loop
			break
		}
	}
	//hace la validación de que la variavle succes este en el boleano true y si esta en falase mas de 10 veces 
	if !success{
		fmt.Println("Lo siento, no adivinaste el número. Número de veces que quedan: ", target)
	}
}
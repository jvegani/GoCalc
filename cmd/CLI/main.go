/*
* File: calculadora.go
*
* Author1:  Juan Vega, (jvega@vozy.co)
* Date:     May 2022
* Partner:  I worked alone
* Course:   Vozy assesment
*
* Summary of File:
*
*   This file contains some functions for arithmetic operations
*
 */

package main

import (
	"fmt"

	"github.com/JuanVegaN/GoCalc.git/pkg/mathops"
)

func main() {
	var opcion, subopcion int
	var num1, num2 float64
	menu := []string{"Suma", "Resta", "Multiplicacion", "Division", "Raiz Cuadrada", "Salir"}
	for opcion != 6 {
		fmt.Print("\n\n----------------------------------------------------------")
		fmt.Printf("\n\t\tBienvenido al Programa Calculadora\n\nPor favor, ingrese el numero de la operacion que desea realizar:\n\n")

		for i, v := range menu {
			fmt.Printf("%d. %v\n", i+1, v)
		}

		switch fmt.Scanln(&opcion); opcion {
		case 1:
			fmt.Printf("\nHas Elegido la opcion %s:", menu[opcion-1])
			for {
				fmt.Printf("\nPara sumar dos numeros enteros, marque 1, para sumar dos numeros flotantes, marque 2: ")
				fmt.Scan(&subopcion)
				if subopcion == 1 {
					fmt.Printf("\nIngrese el primer valor entero: ")
					fmt.Scan(&num1)
					fmt.Printf("\nIngrese el segundo valor entero: ")
					fmt.Scan(&num2)
					fmt.Printf("\nLa suma entre %d y %d es: %d\n", int(num1), int(num2), mathops.Add(int(num1), int(num2)))
					break
				} else if subopcion == 2 {
					fmt.Printf("\nIngrese el primer valor flotante: ")
					fmt.Scan(&num1)
					fmt.Printf("\nIngrese el segundo valor flotante: ")
					fmt.Scan(&num2)
					fmt.Printf("\nLa suma entre %f y %f es: %f\n", num1, num2, mathops.AddFloats(num1, num2))
					break
				} else {
					fmt.Print("Opcion invalida, intente nuevamente")
				}
			}

		case 2:
			fmt.Printf("\nHas Elegido la opcion %s:", menu[opcion-1])
			for {
				fmt.Printf("\nPara restar dos numeros enteros, marque 1, para restar dos numeros flotantes, marque 2: ")
				fmt.Scan(&subopcion)
				if subopcion == 1 {
					fmt.Printf("\nIngrese el primer valor entero: ")
					fmt.Scan(&num1)
					fmt.Printf("\nIngrese el segundo valor entero: ")
					fmt.Scan(&num2)
					fmt.Printf("\nLa resta entre %d y %d es: %d\n", int(num1), int(num2), mathops.Substract(int(num1), int(num2)))
					break
				} else if subopcion == 2 {
					fmt.Printf("\nIngrese el primer valor flotante: ")
					fmt.Scan(&num1)
					fmt.Printf("\nIngrese el segundo valor flotante: ")
					fmt.Scan(&num2)
					fmt.Printf("\nLa resta entre %f y %f es: %f\n", num1, num2, mathops.SubstractFloats(num1, num2))
					break
				} else {
					fmt.Print("Opcion invalida, intente nuevamente")
				}
			}

		case 3:
			fmt.Printf("\nHas Elegido la opcion %s:", menu[opcion-1])
			for {
				fmt.Printf("\nPara multiplicar dos numeros enteros, marque 1, para multiplicar dos numeros flotantes, marque 2: ")
				fmt.Scan(&subopcion)
				if subopcion == 1 {
					fmt.Printf("\nIngrese el primer valor entero: ")
					fmt.Scan(&num1)
					fmt.Printf("\nIngrese el segundo valor entero: ")
					fmt.Scan(&num2)
					fmt.Printf("\nLa multiplicacion entre %d y %d es: %d\n", int(num1), int(num2), mathops.Mult(int(num1), int(num2)))
					break
				} else if subopcion == 2 {
					fmt.Printf("\nIngrese el primer valor flotante: ")
					fmt.Scan(&num1)
					fmt.Printf("\nIngrese el segundo valor flotante: ")
					fmt.Scan(&num2)
					fmt.Printf("\nLa multiplicacion entre %f y %f es: %f\n", num1, num2, mathops.MultFloats(num1, num2))
					break
				} else {
					fmt.Print("Opcion invalida, intente nuevamente")
				}
			}

		case 4:
			fmt.Printf("\nHas Elegido la opcion %s:", menu[opcion-1])
			for {
				fmt.Printf("\nPara dividir dos numeros enteros, marque 1, para dividir dos numeros flotantes, marque 2: ")
				fmt.Scan(&subopcion)
				if subopcion == 1 {
					fmt.Printf("\nIngrese el primer valor entero: ")
					fmt.Scan(&num1)
					fmt.Printf("\nIngrese el segundo valor entero: ")
					fmt.Scan(&num2)
					fmt.Printf("\nLa division entre %d y %d es: %d\n", int(num1), int(num2), mathops.Div(int(num1), int(num2)))
					break
				} else if subopcion == 2 {
					fmt.Printf("\nIngrese el primer valor flotante: ")
					fmt.Scan(&num1)
					fmt.Printf("\nIngrese el segundo valor flotante: ")
					fmt.Scan(&num2)
					fmt.Printf("\nLa division entre %f y %f es: %f\n", num1, num2, mathops.DivFloats(num1, num2))
					break
				} else {
					fmt.Print("Opcion invalida, intente nuevamente")
				}
			}

		case 5:
			fmt.Printf("\nHas Elegido la opcion %s:", menu[opcion-1])
			fmt.Printf("\nPara calcular la raiz cuadrada se utiliza el metodo de newthon\n")
			fmt.Printf("\nIngrese el valor (puede ser entero o flotante) para el cual desea conocer su raiz cuadrada: ")
			fmt.Scan(&num1)
			fmt.Printf("\nLa raiz cuadrada de %f es: %v", num1, mathops.Sqrt(num1))
			//break

		case 6:
			//break

		default:
			fmt.Print("Opcion invalida, intente nuevamente")
		}

	}
}

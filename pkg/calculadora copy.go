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

import "fmt"

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
					fmt.Printf("\nLa suma entre %d y %d es: %d\n", int(num1), int(num2), Add(int(num1), int(num2)))
					break
				} else if subopcion == 2 {
					fmt.Printf("\nIngrese el primer valor flotante: ")
					fmt.Scan(&num1)
					fmt.Printf("\nIngrese el segundo valor flotante: ")
					fmt.Scan(&num2)
					fmt.Printf("\nLa suma entre %f y %f es: %f\n", num1, num2, AddFloats(num1, num2))
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
					fmt.Printf("\nLa resta entre %d y %d es: %d\n", int(num1), int(num2), Substract(int(num1), int(num2)))
					break
				} else if subopcion == 2 {
					fmt.Printf("\nIngrese el primer valor flotante: ")
					fmt.Scan(&num1)
					fmt.Printf("\nIngrese el segundo valor flotante: ")
					fmt.Scan(&num2)
					fmt.Printf("\nLa resta entre %f y %f es: %f\n", num1, num2, SubstractFloats(num1, num2))
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
					fmt.Printf("\nLa multiplicacion entre %d y %d es: %d\n", int(num1), int(num2), Mult(int(num1), int(num2)))
					break
				} else if subopcion == 2 {
					fmt.Printf("\nIngrese el primer valor flotante: ")
					fmt.Scan(&num1)
					fmt.Printf("\nIngrese el segundo valor flotante: ")
					fmt.Scan(&num2)
					fmt.Printf("\nLa multiplicacion entre %f y %f es: %f\n", num1, num2, MultFloats(num1, num2))
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
					fmt.Printf("\nLa division entre %d y %d es: %d\n", int(num1), int(num2), Div(int(num1), int(num2)))
					break
				} else if subopcion == 2 {
					fmt.Printf("\nIngrese el primer valor flotante: ")
					fmt.Scan(&num1)
					fmt.Printf("\nIngrese el segundo valor flotante: ")
					fmt.Scan(&num2)
					fmt.Printf("\nLa division entre %f y %f es: %f\n", num1, num2, DivFloats(num1, num2))
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
			fmt.Printf("\nLa raiz cuadrada de %f es: %v", num1, Sqrt(num1))
			break

		case 6:
			break

		default:
			fmt.Print("Opcion invalida, intente nuevamente")
		}

	}
}

func Add(a, b int) int {
	// recuperado de https://iq.opengenus.org/addition-using-bitwise-operations/
	for b != 0 {
		carry := a & b // Carry value is calculated
		a = a ^ b      // Sum value is calculated and stored in a
		b = carry << 1 // The carry value is shifted towards left by a bit
	}
	return a // returns the final sum
}

func Substract(a, b int) int {
	// recuperado de https://www.geeksforgeeks.org/subtract-two-numbers-without-using-arithmetic-operators/?ref=rp
	for b != 0 {
		borrow := ^(a) & b
		a = a ^ b
		b = borrow << 1
	}
	return a // returns the final subs
}

func Div(a, b int) int {
	cont := 0
	if !(b == 0) {
		for i := 0; a-b >= 0; i++ {
			cont++
			a = a - b
		}
	}
	return cont
}

func Mult(a, b int) int {
	var c int
	if a != 0 && b != 0 {
		for i := 0; i < a; i++ {
			c = c + b
		}
		return c
	}
	return 0
}

func Sqrt(x float64) float64 {
	z_ant := 1.1
	var i int
	for z := float64(1); !(z/z_ant == 1.0); z -= (z*z - x) / (2 * z) {
		i++
		z_ant = z //how to simplify these two lines?
		fmt.Printf("Calculando Raiz Cuadrada de %.2f --> Iteracion %d: %.6f\n", x, i, z_ant)
	}
	return z_ant
}

func AddFloats(a, b float64) (x float64) {
	x = a + b
	return
}

func MultFloats(a, b float64) float64 {
	return a * b
}

func DivFloats(a, b float64) float64 {
	return a / b
}

func SubstractFloats(a, b float64) (x float64) {
	x = a - b
	return
}

/*
* File: sqrt.go
*
* Author1:  Juan Vega, (jvega@vozy.co)
* Date:     May 2022
* Partner:  I worked alone
* Course:   Vozy assesment
*
* Summary of File:
*
*   This file contains a function that calculates the square of a given number
*   using the well-known Newton's Method, based on the formula (z*z - x) / (2*z)
*   where x is the given number and z is the number that is is approximating
*   to be the square root for the given number x.
*
 */

package mathops

import (
	"fmt"
)

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

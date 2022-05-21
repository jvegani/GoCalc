/*
* File: mult.go
*
* Author1:  Juan Vega, (jvega@vozy.co)
* Date:     May 2022
* Partner:  I worked alone
* Course:   Vozy assesment
*
* Summary of File:
*
*   This function mults two int numbers by using loop.
*
 */

package mathops

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

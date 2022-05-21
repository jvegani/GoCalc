/*
* File: add.go
*
* Author1:  Juan Vega, (jvega@vozy.co)
* Date:     May 2022
* Partner:  I worked alone
* Course:   Vozy assesment
*
* Summary of File:
*
*   This function adds two int numbers by using bitwise operators.
*
 */

package mathops

func Add(a, b int) int {
	// recuperado de https://iq.opengenus.org/addition-using-bitwise-operations/
	for b != 0 {
		carry := a & b // Carry value is calculated
		a = a ^ b      // Sum value is calculated and stored in a
		b = carry << 1 // The carry value is shifted towards left by a bit
	}
	return a // returns the final sum
}

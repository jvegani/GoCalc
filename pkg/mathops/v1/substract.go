/*
* File: substract.go
*
* Author1:  Juan Vega, (jvega@vozy.co)
* Date:     May 2022
* Partner:  I worked alone
* Course:   Vozy assesment
*
* Summary of File:
*
*   This function substract two int numbers by using bitwise operators.
*
 */
package mathops

func Substract(a, b int) int {
	// recuperado de https://www.geeksforgeeks.org/subtract-two-numbers-without-using-arithmetic-operators/?ref=rp
	for b != 0 {
		borrow := ^(a) & b
		a = a ^ b
		b = borrow << 1
	}
	return a // returns the final subs
}

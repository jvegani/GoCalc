/*
* File: div.go
*
* Author1:  Juan Vega, (jvega@vozy.co)
* Date:     May 2022
* Partner:  I worked alone
* Course:   Vozy assesment
*
* Summary of File:
*
*   This function divs two int numbers by using loop.
*
 */

package mathops

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

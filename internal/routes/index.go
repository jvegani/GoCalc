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
*   This file contains the route localhost:3000/index
*   And some types used in other files
*
 */

package routes

import (
	"fmt"
	"net/http"
)

type JsonReq struct {
	Float          bool
	FirstOperator  float64
	SecondOperator float64
}

type JsonResp struct {
	Operation      string
	FirstOperator  float64
	SecondOperator float64
	Result         float64
	Message        string
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome GoCalc API")
}

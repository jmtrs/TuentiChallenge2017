/* Problem 1 - PiZZA
 * Author: jmtrs
 * Languaje: GoLang
 * Execution: go run 1.go < a.in > a.out
 */

package main

import (
	"fmt"
)

//import "math"

func main() {
	var cases int
	var people int
	// var portion int
	//var pizzas int

	counter := 0

	fmt.Scanf("%d", &cases)
	fmt.Printf("Casos: %d\n", cases)

	for i := 1; i <= cases; i++ {

		fmt.Scanf("%d", &people)
		fmt.Printf("%d\n %d", people, counter)
		counter++
	}
}

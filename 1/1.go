/* Problem 1 - PiZZA
 * Author: jmtrs
 * Languaje: GoLang
 * Execution: go run 1.go < a.in > a.out
 */

package main

import "fmt"

//import "math"

func main() {
	var cases int
	var people string
	// var portion int
	//var pizzas int
	
	counter := 0
	for scanner.Scan() {
	    line := scanner.Text()
	    counter++
	    // do something with your line
	}
	fmt.Printf("Lines read: %d", counter)

	fmt.Scanf("%d", &cases)
	fmt.Printf("Casos: %d\n", cases)
	for i := 1; i <= cases*2; i++ {
		fmt.Scanf("%[^\n]%*s", &people)
		fmt.Printf("%s\n", people)
		i++
		
		
	}
}

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
    //var people int
    var portion int
    //var pizzas int
    
	fmt.Scanf("%d", &cases)
	for i:= 1; i <= cases; cases++{
    	 fmt.Scanf("%d", &portion) 
    	 fmt.Printf("%f\n", portion)
    	   cases++ 
	}
}
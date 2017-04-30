package main
 
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
 
func init() {
	log.SetFlags(log.Lshortfile)
}
 
func main() {
	var cases int
	var people int
	var portions string
	var portionsTotal int
	
	// Open an input file, exit on error.
	inputFile, err := os.Open("a.in")
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}
 
	// Closes the file when we leave the scope of the current function,
	// this makes sure we never forget to close the file if the
	// function can exit in multiple places.
	defer inputFile.Close()
 
	scanner := bufio.NewScanner(inputFile)
 
	// scanner.Scan() advances to the next token returning false if an error was encountered
	counter := 0
	for scanner.Scan() {
		counter++
		if(counter == 1){
			cases,err =  strconv.Atoi(scanner.Text())
			if err != nil {
		        fmt.Println(err)
		    }
  
		}else if (counter%2 == 0){
			people,err =  strconv.Atoi(scanner.Text())
			if err != nil {
		        fmt.Println(err)
		    }
		    fmt.Printf("people: %d\n", people)
			
		}else{
			portionsTotal,err += strconv.Atoi(strings.Fields(portions))
			//portions,err =  strconv.Atoi(scanner.Text())
			if err != nil {
		        fmt.Println(err)
		    }
		    fmt.Printf("portions: %d\n", portionsTotal)
			
		}
			
		//fmt.Println(scanner.Text())
		//fmt.Printf("%d\n", counter)
	}
	
	fmt.Printf("casos: %d", cases)
 
	// When finished scanning if any error other than io.EOF occured
	// it will be returned by scanner.Err().
	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}
}
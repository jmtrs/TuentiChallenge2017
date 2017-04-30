package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	var pizzas float64
	var pizzasdef int
	var aux int
	var count int

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
		if counter == 1 {
			cases, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}

		} else if counter%2 == 0 {
			people, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Printf("people: %d\n", people)

		} else {
			portions = scanner.Text()
			portions := strings.Fields(portions)
			if err != nil {
				fmt.Println(err)
			}

			for i := 0; i < len(portions); i++ {
				aux, err = strconv.Atoi(portions[i])
				if err != nil {
					//proper err handling
				}
				portionsTotal += aux
				//fmt.Println(aux)

				if i == people-1 {
					pizzas = float64(portionsTotal) / 8
					if math.Mod(pizzas, 1) > 0 {
						pizzas++
					}
					pizzasdef = int(pizzas)
					count++
					fmt.Printf("case %d: %d\n", count, pizzasdef)
					aux = 0
					portionsTotal = 0
				}
			}

			//fmt.Printf("portions: %s\n", portions)

		}

		//fmt.Println(scanner.Text())
		//fmt.Printf("%d\n", counter)
	}

	//fmt.Printf("casos: %d", cases)

	// When finished scanning if any error other than io.EOF occured
	// it will be returned by scanner.Err().
	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}
}

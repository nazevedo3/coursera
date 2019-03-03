/*Write a program which prompts the user to enter a floating point number
and prints the integer which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place. */

package main

import (
	"fmt"
	"log"
)

func main() {
	var num float64
	for x := 0; x < 2; x++ {
		fmt.Println("Please enter a floating point number: ")
		_, e := fmt.Scan(&num)
		if e != nil {
			log.Fatal(e)
		}
		fmt.Println(int(num))

	}
}

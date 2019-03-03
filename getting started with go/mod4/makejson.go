// Write a program which prompts the user to first enter a name, and then enter an address.
// Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
// Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	//make map
	m := make(map[string]string)

	n := bufio.NewReader(os.Stdin)
	fmt.Println("What your name? ")
	name, err := n.ReadString('\n')

	name = name[:len(name)-1] //remove the new line that gets returned from ReadString

	a := bufio.NewReader(os.Stdin)
	fmt.Println("What's your address? ")
	address, err := a.ReadString('\n')
	address = address[:len(address)-1] //remove the new line that gets returned from ReadString

	//Assign values to the "name" and "address" keys
	m["name"] = name
	m["address"] = address

	//Marshal the map to json
	bs, err := json.Marshal(m)
	if err != nil {
		log.Fatal("There was an error: ", err)
	}

	//Print the new json object
	os.Stdout.Write(bs)

}

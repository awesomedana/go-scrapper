package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "budae-jjige"}
	// hayeon := person{"hayeon", 28, favFood}
	hayeon := person{name: "hayeon", age: 28, favFood: favFood} //Looks better
	fmt.Println(hayeon.name)
}

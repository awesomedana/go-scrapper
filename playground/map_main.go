package main

import "fmt"

func main() {
	hayeon := map[string]string{"name": "hayeon", "age": "28"} //[key의 type]value의 type
	for key, _ := range hayeon {
		fmt.Println(key)
	}
}

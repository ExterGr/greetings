package main

import "fmt"

func Hello(name string) string {
	res := fmt.Sprintf("Hi %v, bienvenidooo !!", name)
	return res
}

func main() {
	fmt.Println(Hello("Ivan"))
	goodEvening()
}

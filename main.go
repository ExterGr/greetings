package greetings

import "fmt"

func Hello(name string) string {
	goodEvening()
	res := fmt.Sprintf("Hi %v, bienvenidooo !!", name)
	return res
}

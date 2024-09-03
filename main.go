package greetings

import "fmt"

func Hello(name string) string {
		res := fmt.Sprintf("Hi %v, welcome !!", name)
		return res
}

package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const hola = "Hola, "
const bonjour = "Bonjour, "
const eng = "Hello, "

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	return greeting(language) + name
}

func greeting(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = hola
	case french:
		prefix = bonjour
	default:
		prefix = eng
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}

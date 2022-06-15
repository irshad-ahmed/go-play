package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreetings() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
func (englishBot) getGreetings() string {
	return "Hi There!"
}

func (spanishBot) getGreetings() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreetings())
}

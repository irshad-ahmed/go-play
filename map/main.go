package main

import "fmt"

func main() {
	//var colors map[string]string
	//colors := make(map[int]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ksd000",
	}

	colors["white"] = "#ffffff"

	delete(colors, "white")
	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hex := range colors {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}

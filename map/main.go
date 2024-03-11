package main

import (
	"fmt"
)

func main() {
	// Different ways of declaring maps

	// colors := map[string]string{
	// 	"red":   "#ff0002",
	// 	"green": "#f41333",
	// }

	// var colors map
	colors := make(map[string]string)
	colors["white"] = "#ffffff"

	fmt.Println(len(colors))
	printMap(colors)
	delete(colors, "white")
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}
}

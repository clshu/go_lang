package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":  "#ff0000",
		"lime": "#00ff00",
	}

	var colors1 map[string]string

	colors2 := make(map[string]string)

	fmt.Println(colors)
	fmt.Println(colors1)
	fmt.Println(colors2)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%s %s\n", color, hex)
	}
}

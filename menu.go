package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hi")
	fmt.Println("Loading Menu from csv")
	csv, err := os.ReadFile("recipes.csv")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(csv))

}

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hi")
	fmt.Println("Loading Menu from csv")
	raw, err := os.ReadFile("recipes.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(raw))

	reader := csv.NewReader(strings.NewReader(string(raw)))

	recipes, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(recipes)

}

package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// note to self this remove method does not preserve order
// lotsa discussion here https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func remove(s [][]string, i int) [][]string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

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
	fmt.Println(len(recipes))
	fmt.Println(recipes[2])

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	leftovers := 0
	recipe := ""
	menu := make([][]string, 0)
	for _, v := range days {
		if leftovers == 0 {
			choice := rand.Intn(len(recipes))
			recipe = recipes[choice][0]
			fmt.Println("Adding ", v, "recipe :", recipes[choice])
			numDays, err := strconv.Atoi(recipes[choice][1])
			numDays /= 4
			if err != nil {
				panic(err)
			}
			fmt.Println("this recipe lasts ", numDays, " days")
			recipes = remove(recipes, choice)
			if numDays == 3 {
				toappend := []string{v, recipe}
				menu = append(menu, toappend)
				leftovers = 2
			} else if numDays == 2 {
				toappend := []string{v, recipe}
				menu = append(menu, toappend)
				leftovers = 1
			} else if numDays == 1 {
				toappend := []string{v, recipe}
				menu = append(menu, toappend)
				leftovers = 0
			}
		} else if leftovers > 0 {
			toappend := []string{v, recipe}
			menu = append(menu, toappend)
			leftovers -= 1
		}
	}
	fmt.Println(menu)
}

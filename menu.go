package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// note to self this remove method does not preserve order
// lotsa discussion here https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func remove(s [][]string, i int) [][]string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func printMenu(menu [][]string) {
	println("This weeks menu selections:")
	for _, recipe := range menu {
		fmt.Printf("%-12s %s\n", recipe[0], recipe[1])
	}
}

func GetProspectiveRecipes() [][]string {
	raw, err := os.ReadFile("recipes.csv")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(strings.NewReader(string(raw)))

	recipes, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return recipes
}

func main() {

	recipes := GetProspectiveRecipes()

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	rand.Seed(time.Now().UnixNano())

	leftovers := 0
	recipe := ""
	menu := make([][]string, 0)
	for _, day := range days {
		if leftovers == 0 {
			choice := rand.Intn(len(recipes))
			recipe = recipes[choice][0]

			numDays, err := strconv.Atoi(recipes[choice][1])
			numDays /= 4
			if err != nil {
				panic(err)
			}

			recipes = remove(recipes, choice)
			if numDays == 3 {
				toappend := []string{day, recipe}
				menu = append(menu, toappend)
				leftovers = 2
			} else if numDays == 2 {
				toappend := []string{day, recipe}
				menu = append(menu, toappend)
				leftovers = 1
			} else if numDays == 1 {
				toappend := []string{day, recipe}
				menu = append(menu, toappend)
				leftovers = 0
			}
		} else if leftovers > 0 {
			toappend := []string{day, recipe}
			menu = append(menu, toappend)
			leftovers -= 1
		}
	}
	fmt.Println(menu)
	if leftovers > 0 {
		fmt.Println("There are still leftovers.")
	}
	printMenu(menu)
}

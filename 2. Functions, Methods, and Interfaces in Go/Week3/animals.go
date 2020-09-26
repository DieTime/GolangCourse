package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	var animal, info string
	var selected Animal

	// Prepared animals
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		fmt.Print("> ")

		// Read animal name and info type
		_, err := fmt.Scanln(&animal, &info)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Choose animal by name
		switch animal {
		case "cow":
			selected = cow
		case "bird":
			selected = bird
		case "snake":
			selected = snake
		default:
			fmt.Println("Please enter correct line!")
			continue
		}

		// Choose method by info type
		switch info {
		case "eat":
			selected.Eat()
		case "move":
			selected.Move()
		case "speak":
			selected.Speak()
		default:
			fmt.Println("Please enter correct line!")
			continue
		}
	}
}

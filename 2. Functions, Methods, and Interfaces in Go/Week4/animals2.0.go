package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
	GetName() string
}

type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

func (c Cow) GetName() string {
	return c.name
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func (b Bird) GetName() string {
	return b.name
}

type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func (s Snake) GetName() string {
	return s.name
}

func GenerateAnimal(animalType, name string) (Animal, error) {
	switch animalType {
	case "cow":
		return Cow{name}, nil
	case "bird":
		return Bird{name}, nil
	case "snake":
		return Snake{name}, nil
	default:
		return nil, fmt.Errorf("incorrect animal type")
	}
}

func DoQuery(animal Animal, queryType string) error {
	switch queryType {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		return fmt.Errorf("incorrect requested information type")
	}
	return nil
}

func FindByName(animals []Animal, name string) (Animal, bool) {
	for _, item := range animals {
		if item.GetName() == name {
			return item, true
		}
	}
	return nil, false
}

func main() {
	var command, name, info string
	var animals []Animal

	fmt.Println("Examples of queries:")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("newanimal <name> <animal type (cow, bird, snake)>")
	fmt.Println("query <name> <requested information (eat, move, speak)>")
	fmt.Println("-------------------------------------------------------")

MainLoop:
	for {
		fmt.Print("> ")

		// Read animal name and command type
		_, err := fmt.Scanln(&command, &name, &info)

		if err != nil {
			fmt.Println("You need pass 2 parameters after query name!")
			continue MainLoop
		}

		switch command {
		case "newanimal":
			// Check if animal already exists
			_, found := FindByName(animals, name)

			// If animal not exists add new animal to slice
			if !found {
				animal, err := GenerateAnimal(info, name)
				if err != nil {
					fmt.Println(err)
					continue MainLoop
				}
				animals = append(animals, animal)
			} else {
				fmt.Printf("Animal with name '%s' already exists\n", name)
				continue MainLoop
			}
		case "query":
			// Get animal from slice by name
			animal, found := FindByName(animals, name)

			// If animal not found print error
			if !found {
				fmt.Printf("Can't find animal with name '%s'\n", name)
				continue MainLoop
			}

			// Eif animal was founded do query
			err = DoQuery(animal, info)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

/*
Використовуються структури Animal, Cage та Zookeeper.
Метод AddAnimal - додавання звіря до клітки,
Метод CollectAnimals - збір всіх звірів в клітки.
*/
package main

import "fmt"

type Animal struct {
	Name string
}

type Cage struct {
	Animals []Animal
}

type Zookeeper struct {
	Name string
}

// AddAnimal add Animal to Cage
func (c *Cage) AddAnimal(animal Animal) {
	c.Animals = append(c.Animals, animal)
	fmt.Printf("# %s located in cage\n", animal.Name)
}

// CollectAnimals Put all animals to cages
func (z *Zookeeper) CollectAnimals(cages Cage) {
	fmt.Printf("Zookeeper %s collects animals from cages:\n", z.Name)
	for _, animal := range cages.Animals {
		fmt.Printf("- %s\n", animal.Name)
	}
}

func main() {
	// Create animals
	lion := Animal{Name: "Leon"}
	elephant := Animal{Name: "Elephant"}
	tiger := Animal{Name: "Tiger"}

	// create the Cage
	cage := Cage{}

	// adding animals to Cages
	cage.AddAnimal(lion)
	cage.AddAnimal(elephant)
	cage.AddAnimal(tiger)

	// Create Zookeeper
	zookeeper := Zookeeper{Name: "John Doe"}

	// Збір звірів
	zookeeper.CollectAnimals(cage)
}

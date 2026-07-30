package main

import "fmt"

type Practice struct {
	ID int
	Name string
	Frequency string
	Completed bool
}

func main() {

	var practices []Practice

	// practice1 := Practice{1, "Rosary", "Everyday", true}
	practice1 := Practice{
		ID:			1, 
		Name:		"Rosary", 
		Frequency: 	"Everyday", 
		Completed: 	true,
	}

	// practice2 := Practice{2, "Lectio", "Everyday", false}
	practice2 := Practice{
		ID:			2, 
		Name:		"Lectio", 
		Frequency: 	"Everyday", 
		Completed: 	false,
	}

	practices = append(practices, practice1)
	practices = append(practices, practice2)

	// fmt.Println("Pratice name: ", practice1.Name)
	// fmt.Println("Frequency", practice1.Frequency)
	// fmt.Println("Completed: ", practice1.Completed)

	// fmt.Println("\nPratice name: ", practice2.Name)
	// fmt.Println("Frequency", practice2.Frequency)
	// fmt.Println("Completed: ", practice2.Completed)

	// fmt.Println(practices)

	// for i := 0; i<len(practices); i++ {
	// 	fmt.Println("\nPratice name: ", practices[i].Name)
	// 	fmt.Println("Frequency: ", practices[i].Frequency)
	// 	fmt.Println("Completed: ", practices[i].Completed)
	// }

	for _, practice := range practices {
		fmt.Println(practice.Name)
		fmt.Println(practice.Frequency)
		fmt.Println(practice.Completed)
	}
}
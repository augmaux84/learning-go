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

	practice1 := Practice{
		ID: 1, 
		Name: "Rosary", 
		Frequency: "Everyday", 
		Completed: true,
	}

	practice2 := Practice{
		ID:	2, 
		Name: "Lectio", 
		Frequency: "Everyday", 
		Completed: false,
	}

	practices = append(practices, practice1)
	practices = append(practices, practice2)

	ListPractices(practices)
}

func ListPractices(practices []Practice) {

	for _, practice := range practices {
		fmt.Println(practice.Name)
		fmt.Println(practice.Frequency)
		fmt.Println(practice.Completed)
		fmt.Println()
	}
}
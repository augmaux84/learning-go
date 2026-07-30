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

	CompletePractice(practices, 2)
	ListPractices(practices)

	// fmt.Println(practices[1])

}

func ListPractices(practices []Practice) {

	for _, practice := range practices {
		fmt.Println(practice.Name)
		fmt.Println(practice.Frequency)
		fmt.Println(practice.Completed)
		fmt.Println()
	}

}

func CompletePractice(practices []Practice, id int) {

	for i := 1; i<len(practices); i++ {
		if practices[i].ID == id {
			practices[i].Completed = true
			fmt.Println(practices[i].ID, " transformado em: ", practices[i].Completed)
			break
		} else {}
	}

}
package main

import "fmt"

// GET
func ListPractices(practices []Practice) {

	for _, practice := range practices {
		fmt.Println(practice.Name)
		fmt.Println(practice.Frequency)
		fmt.Println(practice.Completed)
		fmt.Println()
	}

}

// POST
func AddPractice(Name string, Frequency string) {
		
		var practices []Practice

		practice := Practice{
			ID: len(practices) + 1, 
			Name: Name, 
			Frequency: Frequency, 
			Completed: false,
		}

		practices = append(practices, practice)
}


// UPDATE
func CompletePractice(practices []Practice, id int) {

	for i := 0; i<len(practices); i++ {
		if practices[i].ID == id {
			practices[i].Completed = true
			fmt.Println(practices[i].ID, " transformado em: ", practices[i].Completed)
			break
		} else {}
	}
	
}

// DELETE
func RemovePractice(practices []Practice, id int) []Practice{
	
    	return append(practices[:id], practices[id+1:]...)
	
}
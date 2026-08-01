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
func AddPractice(Name string, Frequency string, practices []Practice) []Practice {

		var practice Practice

		if len(practices) == 0 {
			practice = Practice{
				ID: len(practices) + 1, 
				Name: Name, 
				Frequency: Frequency, 
				Completed: false,
		}} else {
			maxID := practices[len(practices)-1].ID
			// for i := range practices {
			// 	if practices[i].ID > maxID {
			// 		maxID = practices[i].ID
			// 		break
			// 	} else {}
			// }

			practice = Practice{
				ID: maxID + 1, 
				Name: Name, 
				Frequency: Frequency, 
				Completed: false,
			}
		}

		practices = append(practices, practice) 
		fmt.Println(practices)
		return practices
		
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
	
	for i := range practices {
		if practices[i].ID == id {
			practices = append(practices[:i], practices[i+1:]...)
			break	
		} else {}
	}

	fmt.Println(practices)
	return practices
	
}
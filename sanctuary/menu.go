package main

import "fmt"

func Menu(practices []Practice) {

	for {

		fmt.Println("----- Sanctuary -----")
		fmt.Println("1. List practices")
		fmt.Println("2. Add practice")
		fmt.Println("3. Complete practice")
		fmt.Println("4. Remove practice")
		fmt.Println("0. Exit")
		fmt.Println("Enter the value: ")
		
		var value int
		fmt.Scan(&value)

		if value == 0 {
			fmt.Println("Verso l'alto")
			break
		}

		switch value {
		case 1:
			ListPractices(practices)

		case 2:
			fmt.Println("Enter the name: ")
			var Name string 
			fmt.Scan(&Name)

			fmt.Println("Enter the frequency: ")
			var Frequency string
			fmt.Scan(&Frequency)

			practices = AddPractice(Name, Frequency, practices)

		case 3:
			var ID int
			fmt.Println("Enter the pratice ID:")
			fmt.Scan(&ID)
			CompletePractice(practices, ID)
			
		case 4:
			var ID int
			fmt.Println("Enter the pratice ID:")
			fmt.Scan(&ID)
			
			practices = RemovePractice(practices, ID)			
		}
			
	}
}
package main

import "fmt"

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
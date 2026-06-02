package main

import "fmt"

const maxDocument = 1000

type document struct {
	graduationYear int
	nameStudent, mentor, researchTopic, researchTitle string
	graduationStatus bool
}

type documentStudent struct {
	id int
	dataDocument [maxDocument]document
}

func add() {
	
}

func edit()  {
	
}

func delete() {
	
}

func main() {
	var (
		arsip documentStudent
		option, graduationYear int
		title, nameStudent string
	)

	for {
		fmt.Println("Choose option by number")
		fmt.Println("1. Add Document")
		fmt.Println("2. Edit Document")
		fmt.Println("3. Delete Document")
		fmt.Print("4. Exit")
		fmt.Scan(&option)
		
		if option == 1 {
			add()
		} else if option == 2 {
			edit()
		} else if option == 3 {
			delete()
		} else if option == 4 {
			break
		} else {
			fmt.Print("Invalid option")
			fmt.Scan(&option)
		}
	}
}
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

func add(data *documentStudent) {
	var (
		yearGraduation int
		studentName, titleResearch, Mentor, topicResearch string
	)

	if data.id >= maxDocument {
		fmt.Print("Document full")
	} else {
		fmt.Println("Fill thesis data: name student, graduation year, research topic, research title and mentor")
		fmt.Scan(&studentName, &yearGraduation, &topicResearch, &titleResearch, &Mentor)

		data.dataDocument[data.id] = document{
			graduationYear: yearGraduation,
			nameStudent: studentName, 
			mentor: Mentor, 
			researchTopic: topicResearch, 
			researchTitle: titleResearch,
		}

		data.id++
		fmt.Println("success add document")
	}
	fmt.Println()
}

func edit(data *documentStudent)  {
	var (
		yearGraduation int
		titleDocument, studentName, titleResearch, Mentor, topicResearch string
	)

	fmt.Println("Fill title document: ")
	fmt.Scan(&titleDocument)

	for i := 0; i < data.id; i++ {
		if data.dataDocument[i].researchTitle == titleDocument {
			fmt.Println("Fill data: name student, graduation year, research topic, research title and mentor")
			fmt.Scan(&studentName, &yearGraduation, &topicResearch, &titleResearch, &Mentor)

			data.dataDocument[i] = document{
				graduationYear: yearGraduation,
				nameStudent: studentName, 
				mentor: Mentor, 
				researchTopic: topicResearch, 
				researchTitle: titleResearch,
			}

			fmt.Println("Success edit document")
		} else {
			fmt.Println("Document not found")
		}
	}
	fmt.Println()
}

func delete(data *documentStudent) {
	var titleDocument string

	fmt.Print("Fill title document to delete: ")
	fmt.Scan(&titleDocument)

	for i := 0; i < data.id; i++ {
		if data.dataDocument[i].researchTitle == titleDocument {
			
		} else {
			fmt.Println("Document not found")
		}
	}
	fmt.Println()
}

func showData(data documentStudent) {
	if data.id > 0 {
		for i := 0; i < data.id; i++ {
			var documentStudent = data.dataDocument[i]
			fmt.Println("ID", data.id)

			fmt.Println("Name", documentStudent.nameStudent)
			fmt.Println("Topic", documentStudent.researchTopic)
			fmt.Println("Title", documentStudent.researchTitle)
			fmt.Println("Graduate Year", documentStudent.graduationYear)
			fmt.Println("Mentor", documentStudent.mentor)
		}
	} else {
		fmt.Println("Document empty")
	}
	fmt.Println()
}

func summaryData(data documentStudent) {
	fmt.Println("Data")

	if data.id < 1 {
		fmt.Println("Document empty")
	} else {
		
	}
}

func searchSequential()  {
	
}

func searchBinary()  {
	
}

func sortSelection()  {
	
}

func sortInsertion()  {
	
}

func main() {
	var (
		arsip documentStudent
		option int
	)

	for {
		fmt.Println("Choose option by number")
		fmt.Println("1. Show Document")
		fmt.Println("2. Add Document")
		fmt.Println("3. Edit Document")
		fmt.Println("4. Delete Document")
		fmt.Println("5. Summary data")
		fmt.Println("6. Exit")
		fmt.Scan(&option)
		
		if arsip.id < maxDocument {
			if option == 1 {
				showData(arsip)
			} else if option == 2 {
				add(&arsip)
			} else if option == 3 {
				edit(&arsip)
			} else if option == 4 {
				delete(&arsip)
			} else if option == 5 {
				summaryData(arsip)
			} else if option == 6 {
				fmt.Print("Out")
				break
			} else {
				fmt.Println("Invalid option")
				fmt.Scan(&option)
			}
		}
	}
}
package main

import "fmt"

const maxDocument = 1000

type document struct {
	graduationYear                                    int
	nameStudent, mentor, researchTopic, researchTitle string
	graduationStatus                                  bool
}

type documentStudent struct {
	id           int
	dataDocument [maxDocument]document
}

func add(data *documentStudent) {
	var (
		yearGraduation                                    int
		studentName, titleResearch, Mentor, topicResearch string
		statusGraduation                                  bool
	)

	if data.id >= maxDocument {
		fmt.Print("Document full")
	} else {
		fmt.Println("Fill data: name student, graduation year, research topic, research title , mentor and graduation status")
		fmt.Scan(&studentName, &yearGraduation, &topicResearch, &titleResearch, &Mentor, &statusGraduation)

		data.dataDocument[data.id] = document{
			graduationYear:   yearGraduation,
			nameStudent:      studentName,
			mentor:           Mentor,
			researchTopic:    topicResearch,
			researchTitle:    titleResearch,
			graduationStatus: statusGraduation,
		}

		data.id++
		fmt.Println("success add document")
	}
	fmt.Println()
}

func edit(data *documentStudent) {
	var (
		yearGraduation                                                   int
		titleDocument, studentName, titleResearch, Mentor, topicResearch string
		statusGraduation                                                 bool
	)

	fmt.Println("Fill title document: ")
	fmt.Scan(&titleDocument)

	for i := 0; i < data.id; i++ {
		if data.dataDocument[i].researchTitle == titleDocument {
			fmt.Println("Fill data: name student, graduation year, research topic, research title , mentor and graduation status")
			fmt.Scan(&studentName, &yearGraduation, &topicResearch, &titleResearch, &Mentor, &statusGraduation)

			data.dataDocument[i] = document{
				graduationYear:   yearGraduation,
				nameStudent:      studentName,
				mentor:           Mentor,
				researchTopic:    topicResearch,
				researchTitle:    titleResearch,
				graduationStatus: statusGraduation,
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
			for j := 0; j < data.id-1; j++ {
				data.dataDocument[j] = data.dataDocument[j+1]
			}

			data.id--
			fmt.Println("Success delete document")
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
			fmt.Println("ID", i+1)

			fmt.Println("Name", documentStudent.nameStudent)
			fmt.Println("Topic", documentStudent.researchTopic)
			fmt.Println("Title", documentStudent.researchTitle)
			fmt.Println("Graduate Year", documentStudent.graduationYear)
			fmt.Println("Mentor", documentStudent.mentor)
			fmt.Println("Status", documentStudent.graduationStatus)
		}
	} else {
		fmt.Println("Document empty")
	}
	fmt.Println()
}

func searchSequential(data *documentStudent) {
	var titleDocument string
	fmt.Print("Fill title document to search: ")
	fmt.Scan(&titleDocument)

	found := false
	for i := 0; i < data.id; i++ {
		if titleDocument == data.dataDocument[i].researchTitle {
			fmt.Println("Document found")
			fmt.Println("Name", data.dataDocument[i].nameStudent)
			fmt.Println("Topic", data.dataDocument[i].researchTopic)
			fmt.Println("Title", data.dataDocument[i].researchTitle)
			fmt.Println("Graduate Year", data.dataDocument[i].graduationYear)
			fmt.Println("Mentor", data.dataDocument[i].mentor)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Document not found")
	}
	fmt.Println()
}

func searchBinary(data *documentStudent) {
	var titleDocument string
	fmt.Print("Fill title document to search: ")
	fmt.Scan(&titleDocument)

	low := 0
	high := data.id - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		if data.dataDocument[mid].researchTitle == titleDocument {
			fmt.Println("Document found")
			fmt.Println("Name", data.dataDocument[mid].nameStudent)
			fmt.Println("Topic", data.dataDocument[mid].researchTopic)
			fmt.Println("Title", data.dataDocument[mid].researchTitle)
			fmt.Println("Graduate Year", data.dataDocument[mid].graduationYear)
			fmt.Println("Mentor", data.dataDocument[mid].mentor)
			found = true
			break
		} else if data.dataDocument[mid].researchTitle < titleDocument {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if !found {
		fmt.Println("Document not found")
	}
	fmt.Println()

}

func sortSelection(data *documentStudent) {
	for i := 0; i < data.id-1; i++ {
		minIndex := i
		for j := i + 1; j < data.id; j++ {
			if data.dataDocument[j].researchTitle < data.dataDocument[minIndex].researchTitle {
				minIndex = j
			}
		}
		if minIndex != i {
			data.dataDocument[i], data.dataDocument[minIndex] = data.dataDocument[minIndex], data.dataDocument[i]
		}
	}
}

func sortInsertion(data *documentStudent) {
	for i := 1; i < data.id; i++ {
		key := data.dataDocument[i]
		j := i - 1
		for j >= 0 && data.dataDocument[j].researchTitle > key.researchTitle {
			data.dataDocument[j+1] = data.dataDocument[j]
			j--
		}
		data.dataDocument[j+1] = key
	}
}

func StatisticbyYear(data *documentStudent) {
	yearCount := make(map[int]int)
	for i := 0; i < data.id; i++ {
		year := data.dataDocument[i].graduationYear
		yearCount[year]++
	}

	fmt.Println("Graduation Year Statistics:")
	for year, count := range yearCount {
		fmt.Printf("Year: %d, Count: %d\n", year, count)
	}

}

func main() {
	var (
		arsip  documentStudent
		option int
	)

	for {
		fmt.Println("Choose option by number")
		fmt.Println("1. Show Document")
		fmt.Println("2. Add Document")
		fmt.Println("3. Edit Document")
		fmt.Println("4. Delete Document")
		fmt.Println("5. Search Document Sequential (By title)")
		fmt.Println("6. Search Document Binary (By title)")
		fmt.Println("7. Sort Document Selection (By Title)")
		fmt.Println("8. Sort Document Insertion (By Title)")
		fmt.Println("9. Statistic by Year")
		fmt.Println("10. Exit")
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
				searchSequential(&arsip)
			} else if option == 6 {
				searchBinary(&arsip)
			} else if option == 7 {
				sortSelection(&arsip)
				showData(arsip)
			} else if option == 8 {
				sortInsertion(&arsip)
				showData(arsip)
			} else if option == 9 {
				StatisticbyYear(&arsip)
			} else if option == 10 {
				fmt.Println("Exiting...")
				break
			} else {
				fmt.Println("Invalid option")
				fmt.Scan(&option)
			}
		}
	}
}

package main

import "fmt"

const maxDocument = 1000

type document struct {
	id, graduationYear                                int
	nameStudent, mentor, researchTopic, researchTitle string
	graduationStatus                                  bool
}

type documentStudent struct {
	totalDocument int
	dataDocument  [maxDocument]document
}

type yearSummary struct {
	year  int
	count int
}

func add(data *documentStudent) {
	var (
		yearGraduation                                    int
		studentName, titleResearch, Mentor, topicResearch string
		statusGraduate                                    bool
	)

	if data.totalDocument >= maxDocument {
		fmt.Print("Document full")
	} else {
		fmt.Println("Fill name student")
		fmt.Scan(&studentName)
		fmt.Println("Fill graduation year")
		fmt.Scan(&yearGraduation)
		fmt.Println("Fill research topic")
		fmt.Scan(&topicResearch)
		fmt.Println("Fill research title")
		fmt.Scan(&titleResearch)
		fmt.Println("Fill mentor")
		fmt.Scan(&Mentor)
		fmt.Println("Fill graduate status")
		fmt.Scan(&statusGraduate)

		data.dataDocument[data.totalDocument] = document{
			id:               data.totalDocument + 1,
			graduationYear:   yearGraduation,
			nameStudent:      studentName,
			mentor:           Mentor,
			researchTopic:    topicResearch,
			researchTitle:    titleResearch,
			graduationStatus: statusGraduate,
		}

		data.totalDocument++
		fmt.Println("success add document")
	}
	fmt.Println()
}

func edit(data *documentStudent) {
	var (
		foundIndex                                                       int = -1
		yearGraduation                                                   int
		titleDocument, studentName, titleResearch, Mentor, topicResearch string
		statusGraduate                                                   bool
	)

	fmt.Println("Fill title document: ")
	fmt.Scan(&titleDocument)

	for i := 0; i < data.totalDocument; i++ {
		if data.dataDocument[i].researchTitle == titleDocument {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Println("Document not found")
	} else {
		fmt.Println("Fill name student")
		fmt.Scan(&studentName)
		fmt.Println("Fill graduation year")
		fmt.Scan(&yearGraduation)
		fmt.Println("Fill research topic")
		fmt.Scan(&topicResearch)
		fmt.Println("Fill research title")
		fmt.Scan(&titleResearch)
		fmt.Println("Fill mentor")
		fmt.Scan(&Mentor)
		fmt.Println("Fill graduate status")
		fmt.Scan(&statusGraduate)

		data.dataDocument[foundIndex] = document{
			id:               data.dataDocument[foundIndex].id,
			graduationYear:   yearGraduation,
			nameStudent:      studentName,
			mentor:           Mentor,
			researchTopic:    topicResearch,
			researchTitle:    titleResearch,
			graduationStatus: statusGraduate,
		}

		fmt.Println("Success edit document")
	}
	fmt.Println()
}

func delete(data *documentStudent) {
	var (
		foundIndex    int = -1
		titleDocument string
	)

	fmt.Print("Fill title document to delete: ")
	fmt.Scan(&titleDocument)

	for i := 0; i < data.totalDocument; i++ {
		if data.dataDocument[i].researchTitle == titleDocument {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Println("Document not found")
	} else {
		for j := foundIndex; j < data.totalDocument-1; j++ {
			data.dataDocument[j] = data.dataDocument[j+1]
		}

		data.totalDocument--
		fmt.Println("Success delete document")
	}
	fmt.Println()
}

func showData(data documentStudent) {
	if data.totalDocument > 0 {
		for i := 0; i < data.totalDocument; i++ {
			var documentStudent = data.dataDocument[i]
			fmt.Println("ID", documentStudent.id)

			fmt.Println("Name 	  :", documentStudent.nameStudent)
			fmt.Println("Topic 	  :", documentStudent.researchTopic)
			fmt.Println("Title 	  :", documentStudent.researchTitle)
			fmt.Println("Graduate Year:", documentStudent.graduationYear)
			fmt.Println("Mentor    :", documentStudent.mentor)
			fmt.Println("Status    :", documentStudent.graduationStatus)
		}
	} else {
		fmt.Println("Document empty")
	}
	fmt.Println()
}

func summaryData(data documentStudent) {
	var (
		summary                       [maxDocument]yearSummary
		totalSummary, foundYear, year int
	)

	for i := 0; i < data.totalDocument; i++ {
		year = data.dataDocument[i].graduationYear
		foundYear = -1

		for j := 0; j < totalSummary; j++ {
			if summary[j].year == year {
				summary[j].count++
				foundYear = 0
				break
			}
		}

		if foundYear == -1 {
			summary[totalSummary].year = year
			summary[totalSummary].count++
			totalSummary++
		}
	}

	fmt.Println("Graduation Year Statistics:")
	for i := 0; i < totalSummary; i++ {
		fmt.Println("Tahun", summary[i].year, " Count", summary[i].count)
	}
	fmt.Println()
}

func searchSequential(data *documentStudent) {
	var titleDocument string
	fmt.Print("Fill title document to search: ")
	fmt.Scan(&titleDocument)

	found := false
	for i := 0; i < data.totalDocument; i++ {
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

func searchBinary(data documentStudent) {
	var title string
	fmt.Print("Fill title document to search: ")
	fmt.Scan(&title)

	for i := 1; i < data.totalDocument; i++ {
		key := data.dataDocument[i]
		j := i - 1
		for j >= 0 && data.dataDocument[j].researchTitle > key.researchTitle {
			data.dataDocument[j+1] = data.dataDocument[j]
			j--
		}
		data.dataDocument[j+1] = key
	}

	low := 0
	high := data.totalDocument - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		doc := data.dataDocument[mid]

		if doc.researchTitle == title {
			fmt.Println("Document found")
			fmt.Println("Name", doc.nameStudent)
			fmt.Println("Topic", doc.researchTopic)
			fmt.Println("Title", doc.researchTitle)
			fmt.Println("Graduate Year", doc.graduationYear)
			fmt.Println("Mentor", doc.mentor)
			found = true
			break
		} else if doc.researchTitle < title {
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

func sortSelectionbygraduationYear(data documentStudent) {
	for i := 0; i < data.totalDocument-1; i++ {
		min := i
		for j := i + 1; j < data.totalDocument; j++ {
			if data.dataDocument[j].graduationYear < data.dataDocument[min].graduationYear {
				min = j
			}
		}
		if min != i {
			data.dataDocument[i], data.dataDocument[min] = data.dataDocument[min], data.dataDocument[i]
		}
	}
	showData(data)
}

func sortInsertionbyYear(data documentStudent) {
	for i := 1; i < data.totalDocument; i++ {
		key := data.dataDocument[i]
		j := i - 1

		for j >= 0 && data.dataDocument[j].graduationYear > key.graduationYear {
			data.dataDocument[j+1] = data.dataDocument[j]
			j--
		}
		data.dataDocument[j+1] = key
	}
	showData(data)
}

// func StatisticbyYear(data *documentStudent) {
// 	yearCount := make(map[int]int)
// 	for i := 0; i < data.totalDocument; i++ {
// 		year := data.dataDocument[i].graduationYear
// 		yearCount[year]++
// 	}

// 	fmt.Println("Graduation Year Statistics:")
// 	for year, count := range yearCount {
// 		fmt.Printf("Year: %d, Count: %d\n", year, count)
// 	}

// }

func main() {
	var (
		arsip  documentStudent
		option int
	)

	generateDummyData(&arsip)

	for {
		fmt.Println("Choose option by number")
		fmt.Println("1. Show Document")
		fmt.Println("2. Add Document")
		fmt.Println("3. Edit Document")
		fmt.Println("4. Delete Document")
		fmt.Println("5. Summary data")
		fmt.Println("6. Search Document (Sequential)")
		fmt.Println("7. Search Document (Binary)")
		fmt.Println("8. Sort by Year (Selection)")
		fmt.Println("9. Sort by Year (Insertion)")
		fmt.Println("10. Exit")
		fmt.Scan(&option)

		if arsip.totalDocument < maxDocument {
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
				searchSequential(&arsip)
			} else if option == 7 {
				searchBinary(arsip)
			} else if option == 8 {
				sortSelectionbygraduationYear(arsip)
			} else if option == 9 {
				sortInsertionbyYear(arsip)
			} else if option == 10 {
				fmt.Print("Out")
				break
			} else {
				fmt.Println("Nonvalid option")
			}
		}
	}
}

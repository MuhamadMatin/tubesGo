package main

import "fmt"

const maxDocument = 1000

type document struct {
	id, graduationYear int
	nameStudent, mentor, researchTopic, researchTitle string
	graduationStatus bool
}

type documentStudent struct {
	totalDocument int
	dataDocument [maxDocument]document
}

type yearSummary struct {
	year int
	count int
}

func add(data *documentStudent) {
	var (
		yearGraduation int
		studentName, titleResearch, Mentor, topicResearch string
		statusGraduate bool
	)

	if data.totalDocument >= maxDocument {
		fmt.Print("!!! Document Full !!!")
	} else {
		fmt.Println()
		fmt.Println("==============")
		fmt.Println(" ADD DOCUMENT")
		fmt.Println("==============")
		fmt.Println()

		fmt.Print("Name            : ")
		fmt.Scan(&studentName) 
		fmt.Print("Topic           : ") 
		fmt.Scan(&topicResearch) 
		fmt.Print("Title           : ") 
		fmt.Scan(&titleResearch) 
		fmt.Print("Mentor          : ") 
		fmt.Scan(&Mentor) 
		fmt.Print("Graduation Year : ") 
		fmt.Scan(&yearGraduation) 
		fmt.Print("Graduate Status (true/false) : ") 
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
		fmt.Println("^^^ Success Add Document ^^^")
	}
	fmt.Println()
}

func edit(data *documentStudent) {
	var (
		foundIndex int = -1
		yearGraduation int
		titleDocument, studentName, titleResearch, Mentor, topicResearch string
		statusGraduate bool
	)

	fmt.Println()
	fmt.Println("===============")
	fmt.Println(" EDIT DOCUMENT")
	fmt.Println("===============")
	fmt.Println()
	fmt.Print("Fill Title Document To Edit: ")
	fmt.Scan(&titleDocument)

	for i := 0; i < data.totalDocument; i++ {
		if data.dataDocument[i].researchTitle == titleDocument {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Print("!!! Document Not Found !!!")
	} else {
		fmt.Print("Name            : ")
		fmt.Scan(&studentName) 
		fmt.Print("Topic           : ") 
		fmt.Scan(&topicResearch) 
		fmt.Print("Title           : ") 
		fmt.Scan(&titleResearch) 
		fmt.Print("Mentor          : ") 
		fmt.Scan(&Mentor) 
		fmt.Print("Graduation Year : ") 
		fmt.Scan(&yearGraduation) 
		fmt.Print("Graduate Status (true/false) : ") 
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

		fmt.Println("^^^ Success Edit Document ^^^")
	}
	fmt.Println()
}

func delete(data *documentStudent) {
	var (
		foundIndex    int = -1
		titleDocument string
	)

	fmt.Println()
	fmt.Println("=================")
	fmt.Println(" DELETE DOCUMENT")
	fmt.Println("=================")
	fmt.Println()
	fmt.Print("Fill Title Document To Delete: ")
	fmt.Scan(&titleDocument)

	for i := 0; i < data.totalDocument; i++ {
		if data.dataDocument[i].researchTitle == titleDocument {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Println("!!! Document Not Found !!!")
	} else {
		for j := foundIndex; j < data.totalDocument-1; j++ {
			data.dataDocument[j] = data.dataDocument[j+1]
		}

		data.totalDocument--
		fmt.Println("^^^ Success Delete Document ^^^")
	}
	
	fmt.Println()
}

func showData(data documentStudent) {
	if data.totalDocument > 0 {
		fmt.Println()
		fmt.Println("==============")
		fmt.Println(" ALL DOCUMENT")
		fmt.Println("==============")
		fmt.Println()
		for i := 0; i < data.totalDocument; i++ {
			var documentStudent = data.dataDocument[i]
			
			fmt.Println("ID", documentStudent.id)
			fmt.Println("Name             :", documentStudent.nameStudent) 
			fmt.Println("Topic            :", documentStudent.researchTopic) 
			fmt.Println("Title            :", documentStudent.researchTitle) 
			fmt.Println("Mentor           :", documentStudent.mentor) 
			fmt.Println("Graduate Year    :", documentStudent.graduationYear) 
			fmt.Print("Graduation Status: ") 
			if documentStudent.graduationStatus == false { 
				fmt.Println("Not Graduated") 
			} else { 
				fmt.Println("Graduate") 
			}
			fmt.Println()
		}
	} else {
		fmt.Println("!!! Document Empty !!!")
	}
}

func summaryData(data documentStudent) {
	var (
		summary [maxDocument]yearSummary
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

	fmt.Println()
	fmt.Println("===================")
	fmt.Println(" SUMMARIES BY YEAR")
	fmt.Println("===================")
	fmt.Println()
	for i := 0; i < totalSummary; i++ {
		fmt.Println("Tahun", summary[i].year, " Count", summary[i].count)
	}
	fmt.Println()
}

func searchSequential(data *documentStudent) {
	var titleDocument string
	fmt.Print("Fill Title Document To Search: ")
	fmt.Scan(&titleDocument)

	found := false
	for i := 0; i < data.totalDocument; i++ {
		if titleDocument == data.dataDocument[i].researchTitle {
			fmt.Println("^^^ Document Found ^^^")

			fmt.Println("ID", data.dataDocument[i].id)
			fmt.Println("Name             :", data.dataDocument[i].nameStudent) 
			fmt.Println("Topic            :", data.dataDocument[i].researchTopic) 
			fmt.Println("Title            :", data.dataDocument[i].researchTitle) 
			fmt.Println("Mentor           :", data.dataDocument[i].mentor) 
			fmt.Println("Graduate Year    :", data.dataDocument[i].graduationYear) 
			fmt.Print("Graduation Status: ") 
			if data.dataDocument[i].graduationStatus == false { 
				fmt.Println("Not Graduated") 
			} else { 
				fmt.Println("Graduate") 
			}
			found = true
			break
		}
	}
	if !found {
		fmt.Println("!!! Document Not Found !!!")
	}
	fmt.Println()
}

func searchBinary(data documentStudent) {
	var title string
	fmt.Print("Fill Title Document To Search: ")
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
			fmt.Println("^^^ Document Found ^^^")

			fmt.Println("ID", doc.id)
			fmt.Println("Name             :", doc.nameStudent) 
			fmt.Println("Topic            :", doc.researchTopic) 
			fmt.Println("Title            :", doc.researchTitle) 
			fmt.Println("Mentor           :", doc.mentor) 
			fmt.Println("Graduate Year    :", doc.graduationYear) 
			fmt.Print("Graduation Status: ") 
			if doc.graduationStatus == false { 
				fmt.Println("Not Graduated") 
			} else { 
				fmt.Println("Graduate") 
			}
			found = true
			break
		} else if doc.researchTitle < title {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("!!! Document Not Found !!!")
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

func main() {
	var (
		arsip documentStudent
		option int
	)

	// uncomment when you generate dummy data
	generateDummyData(&arsip)

	for {
		fmt.Println("==========")
		fmt.Println(" SkripsIn")
		fmt.Println("==========")
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
				fmt.Print("See ya")
				break
			} else {
				fmt.Println("!!! Invalid Option, Please Input Again !!!")
			}
		}
	}
}

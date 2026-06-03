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

func generateDummyData(data *documentStudent) {
	data.dataDocument[0] = document{
		id:               1,
		graduationYear:   2023,
		nameStudent:      "Budi",
		mentor:           "Pak_Sutrisno",
		researchTopic:    "AI",
		researchTitle:    "Face_Recognition",
		graduationStatus: true,
	}
	data.dataDocument[1] = document{
		id:               2,
		graduationYear:   2024,
		nameStudent:      "Siti",
		mentor:           "Bu_Ratna",
		researchTopic:    "Web",
		researchTitle:    "E-Commerce_Optimization",
		graduationStatus: false,
	}
	data.dataDocument[2] = document{
		id:               3,
		graduationYear:   2023,
		nameStudent:      "Andi",
		mentor:           "Pak_Sutrisno",
		researchTopic:    "Network",
		researchTitle:    "5G_Security",
		graduationStatus: true,
	}
	data.dataDocument[3] = document{
		id:               4,
		graduationYear:   2024,
		nameStudent:      "Rina",
		mentor:           "Pak_Budi",
		researchTopic:    "Mobile",
		researchTitle:    "Flutter_UI_Testing",
		graduationStatus: true,
	}
	data.dataDocument[4] = document{
		id:               5,
		graduationYear:   2022,
		nameStudent:      "Joko",
		mentor:           "Bu_Ratna",
		researchTopic:    "Data_Science",
		researchTitle:    "Predictive_Analysis",
		graduationStatus: true,
	}
	data.dataDocument[5] = document{
		id:               6,
		graduationYear:   2025,
		nameStudent:      "Ayu",
		mentor:           "Pak_Anton",
		researchTopic:    "IoT",
		researchTitle:    "Smart_Farming",
		graduationStatus: false,
	}
	data.dataDocument[6] = document{
		id:               7,
		graduationYear:   2023,
		nameStudent:      "Candra",
		mentor:           "Pak_Anton",
		researchTopic:    "Cyber_Security",
		researchTitle:    "Malware_Detection",
		graduationStatus: true,
	}
	data.dataDocument[7] = document{
		id:               8,
		graduationYear:   2024,
		nameStudent:      "Dewi",
		mentor:           "Bu_Siska",
		researchTopic:    "Game_Dev",
		researchTitle:    "Procedural_Generation",
		graduationStatus: false,
	}
	data.dataDocument[8] = document{
		id:               9,
		graduationYear:   2022,
		nameStudent:      "Eko",
		mentor:           "Pak_Budi",
		researchTopic:    "Cloud_Computing",
		researchTitle:    "Docker_Orchestration",
		graduationStatus: true,
	}
	data.dataDocument[9] = document{
		id:               10,
		graduationYear:   2025,
		nameStudent:      "Fajar",
		mentor:           "Bu_Siska",
		researchTopic:    "Blockchain",
		researchTitle:    "Smart_Contracts",
		graduationStatus: false,
	}

	data.totalDocument = 10
}

func add(data *documentStudent) {
	var (
		yearGraduation int
		studentName, titleResearch, Mentor, topicResearch string
		statusGraduate bool
	)

	if data.totalDocument >= maxDocument {
		fmt.Print("Document full")
	} else {
		fmt.Println("Fill thesis data: name student, graduation year, research topic, research title, mentor and graduate status")
		fmt.Scan(&studentName, &yearGraduation, &topicResearch, &titleResearch, &Mentor, &statusGraduate)

		data.dataDocument[data.totalDocument] = document{
			id: data.totalDocument+1,
			graduationYear: yearGraduation,
			nameStudent: studentName, 
			mentor: Mentor, 
			researchTopic: topicResearch, 
			researchTitle: titleResearch,
			graduationStatus: statusGraduate,
		}

		data.totalDocument++
		fmt.Println("success add document")
	}
	fmt.Println()
}

func edit(data *documentStudent)  {
	var (
		yearGraduation int
		titleDocument, studentName, titleResearch, Mentor, topicResearch string
		statusGraduate bool
	)

	fmt.Println("Fill title document: ")
	fmt.Scan(&titleDocument)

	for i := 0; i < data.totalDocument; i++ {
		if data.dataDocument[i].researchTitle == titleDocument {
			fmt.Println("Fill data: name student, graduation year, research topic, research title, mentor and graduate status")
			fmt.Scan(&studentName, &yearGraduation, &topicResearch, &titleResearch, &Mentor, &statusGraduate)

			data.dataDocument[i] = document{
				graduationYear: yearGraduation,
				nameStudent: studentName, 
				mentor: Mentor, 
				researchTopic: topicResearch, 
				researchTitle: titleResearch,
				graduationStatus: statusGraduate,
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

	for i := 0; i < data.totalDocument; i++ {
		if data.dataDocument[i].researchTitle == titleDocument {
			for j := 0; j < data.totalDocument - 1; j++ {
				data.dataDocument[j] = data.dataDocument[j+1]
			}

			data.totalDocument--
			fmt.Println("Success delete document")
		} else {
			fmt.Println("Document not found")
		}
	}
	fmt.Println()
}

func showData(data documentStudent) {
	if data.totalDocument > 0 {
		for i := 0; i < data.totalDocument; i++ {
			var documentStudent = data.dataDocument[i]
			fmt.Println("ID", documentStudent.id)

			fmt.Println("Name:", documentStudent.nameStudent)
			fmt.Println("Topic:", documentStudent.researchTopic)
			fmt.Println("Title:", documentStudent.researchTitle)
			fmt.Println("Graduate Year:", documentStudent.graduationYear)
			fmt.Println("Mentor:", documentStudent.mentor)
			fmt.Println("Status:", documentStudent.graduationStatus)
		}
	} else {
		fmt.Println("Document empty")
	}
	fmt.Println()
}

func summaryData(data documentStudent) {
	fmt.Println("Data")

	if data.totalDocument < 1 {
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

	generateDummyData(&arsip)

	for {
		fmt.Println("Choose option by number")
		fmt.Println("1. Show Document")
		fmt.Println("2. Add Document")
		fmt.Println("3. Edit Document")
		fmt.Println("4. Delete Document")
		fmt.Println("5. Summary data")
		fmt.Println("6. Exit")
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
				fmt.Print("Out")
				break
			} else {
				fmt.Println("Nonvalid option")
				fmt.Scan(&option)
			}
		}
	}
}
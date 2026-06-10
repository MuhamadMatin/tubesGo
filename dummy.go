package main

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
<p align="center"><a href="https://go.dev" target="_blank"><img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg" width="300" alt="Go Logo"></a></p>
<h1 align="center">📁 Sistem Informasi Inventaris Dokumen Skripsi (SkripsIn)</h1>
<p align="center">
  Kelola arsip tugas akhir mahasiswa dari terminal. SkripsIn melacak judul skripsi, penulis, dan tahun kelulusan. Dirancang untuk staf administrasi program studi dan petugas perpustakaan.
</p>

---

## Catatan Penting
 
- Instal [Go 1.18+](https://go.dev/dl) sebelum menjalankan aplikasi.
- Berjalan di terminal. Tanpa browser, tanpa database eksternal.
- Ganti spasi dengan underscore (`_`) saat input data. Contoh: `ini_Budi`.

---
 
## Instalasi
 
1. Clone repository
```bash
git clone https://github.com/MuhamadMatin/tubesGo.git
cd tubesGo
```
 
2. Pastikan Go terinstal
```bash
go version
```
*(Unduh di [https://go.dev/dl](https://go.dev/dl) jika belum terinstal).*
 
3. Jalankan aplikasi
```bash
go run .
```
 
4. Build menjadi executable (opsional)
```bash
# Build
go build -o skripsin main.go
 
# Linux / macOS
./skripsin
 
# Windows
skripsin.exe
```
 
---

## Alur Program

```mermaid
flowchart TD
    %% --- Styling ---
    classDef startEnd fill:#ff9999,stroke:#333,stroke-width:2px;
    classDef menu fill:#99ccff,stroke:#333,stroke-width:2px;
    classDef process fill:#ffffff,stroke:#333,stroke-width:1px;

    %% --- Main Flow ---
    S([Start]):::startEnd --> INIT[generateDummyData]:::process
    INIT --> MENU[/Show Menu/]:::menu
    MENU --> IN[/Input option/]:::process
    IN --> OPT{Pilih Option?}

    %% --- Option Routing ---
    OPT -->|1| SH_CHK
    OPT -->|2| AD_CHK
    OPT -->|3| ED_IN
    OPT -->|4| DE_IN
    OPT -->|5| SM_LOOP
    OPT -->|6| SQ_IN
    OPT -->|7| BI_IN
    OPT -->|8| SS_LOOP
    OPT -->|9| SI_LOOP
    OPT -->|10| EXIT([Exit]):::startEnd
    OPT -->|else| ERR[/Invalid Option/]
    ERR --> MENU

    %% --- Subgraphs ---

    subgraph g_show [1. Show Data]
        SH_CHK{total > 0?}
        SH_CHK -->|yes| SH_PRINT[print each doc]
        SH_CHK -->|no| SH_EMPTY[Document Empty]
    end

    subgraph g_add [2. Add Data]
        AD_CHK{total >= max?}
        AD_CHK -->|yes| AD_FULL[Document Full]
        AD_CHK -->|no| AD_IN[/input fields/] --> AD_SAVE[store & total++]
    end

    subgraph g_edit [3. Edit Data]
        ED_IN[/input title/] --> ED_FIND{found?}
        ED_FIND -->|no| ED_404[Not Found]
        ED_FIND -->|yes| ED_FIELDS[/input new fields/] --> ED_UPD[update doc]
    end

    subgraph g_del [4. Delete Data]
        DE_IN[/input title/] --> DE_FIND{found?}
        DE_FIND -->|no| DE_404[Not Found]
        DE_FIND -->|yes| DE_SHIFT[shift docs left] --> DE_DEC[total--]
    end

    subgraph g_sum [5. Summary Data]
        SM_LOOP[loop docs] --> SM_CHK{year in summary?}
        SM_CHK -->|yes| SM_INC[count++]
        SM_CHK -->|no| SM_NEW[add year entry]
        SM_INC --> SM_NXT{more docs?}
        SM_NEW --> SM_NXT
        SM_NXT -->|yes| SM_LOOP
        SM_NXT -->|no| SM_PRINT[print by year]
    end

    subgraph g_seq [6. Search Sequential]
        SQ_IN[/input title/] --> SQ_LOOP[i = 0..total]
        SQ_LOOP --> SQ_CHK{title match?}
        SQ_CHK -->|yes| SQ_OK[showData result]
        SQ_CHK -->|no| SQ_NXT{i < total?}
        SQ_NXT -->|yes| SQ_LOOP
        SQ_NXT -->|no| SQ_404[Not Found]
    end

    subgraph g_bin [7. Search Binary]
        BI_IN[/input title/] --> BI_SORT[insertion sort by title]
        BI_SORT --> BI_INIT[low=0 high=n-1]
        BI_INIT --> BI_LOOP{low <= high?}
        BI_LOOP -->|no| BI_404[Not Found]
        BI_LOOP -->|yes| BI_MID[mid = low+high / 2]
        BI_MID --> BI_CMP{doc.title vs title}
        BI_CMP -->|match| BI_OK[showData result]
        BI_CMP -->|less| BI_LO[low = mid+1] --> BI_LOOP
        BI_CMP -->|greater| BI_HI[high = mid-1] --> BI_LOOP
    end

    subgraph g_sel [8. Sort Selection]
        SS_LOOP[i = 0..n-2] --> SS_MIN[find min from i+1..n]
        SS_MIN --> SS_CMP{min != i?}
        SS_CMP -->|yes| SS_SWAP[swap i, min] --> SS_NEXT{i++, done?}
        SS_CMP -->|no| SS_NEXT
        SS_NEXT -->|no| SS_LOOP
        SS_NEXT -->|yes| SS_SHOW[showData]
    end

    subgraph g_ins [9. Sort Insertion]
        SI_LOOP[i = 1..n-1] --> SI_KEY[key = doc i]
        SI_KEY --> SI_SHIFT[shift docs where year > key.year]
        SI_SHIFT --> SI_INS[doc j+1 = key]
        SI_INS --> SI_NEXT{i++, done?}
        SI_NEXT -->|no| SI_LOOP
        SI_NEXT -->|yes| SI_SHOW[showData]
    end

    %% --- Clean Return Paths ---
    BACK_TO_MENU((Kembali ke Menu)):::menu

    SH_PRINT & SH_EMPTY --> BACK_TO_MENU
    AD_FULL & AD_SAVE --> BACK_TO_MENU
    ED_404 & ED_UPD --> BACK_TO_MENU
    DE_404 & DE_DEC --> BACK_TO_MENU
    SM_PRINT --> BACK_TO_MENU
    SQ_OK & SQ_404 --> BACK_TO_MENU
    BI_OK & BI_404 --> BACK_TO_MENU
    SS_SHOW --> BACK_TO_MENU
    SI_SHOW --> BACK_TO_MENU

    BACK_TO_MENU --> MENU
```

---
 
## Data Dummy
 
SkripsIn memuat 10 data awal saat start:
 
| ID | Nama Mahasiswa | Topik           | Judul Penelitian        | Pembimbing   | Tahun | Status   |
|----|----------------|-----------------|-------------------------|--------------|-------|----------|
| 1  | Budi           | AI              | Face_Recognition        | Pak_Sutrisno | 2023  | ✅ Lulus |
| 2  | Siti           | Web             | E-Commerce_Optimization | Bu_Ratna     | 2024  | ❌ Belum |
| 3  | Andi           | Network         | 5G_Security             | Pak_Sutrisno | 2023  | ✅ Lulus |
| 4  | Rina           | Mobile          | Flutter_UI_Testing      | Pak_Budi     | 2024  | ✅ Lulus |
| 5  | Joko           | Data_Science    | Predictive_Analysis     | Bu_Ratna     | 2022  | ✅ Lulus |
| 6  | Ayu            | IoT             | Smart_Farming           | Pak_Anton    | 2025  | ❌ Belum |
| 7  | Candra         | Cyber_Security  | Malware_Detection       | Pak_Anton    | 2023  | ✅ Lulus |
| 8  | Dewi           | Game_Dev        | Procedural_Generation   | Bu_Siska     | 2024  | ❌ Belum |
| 9  | Eko            | Cloud_Computing | Docker_Orchestration    | Pak_Budi     | 2022  | ✅ Lulus |
| 10 | Fajar          | Blockchain      | Smart_Contracts         | Bu_Siska     | 2025  | ❌ Belum |
 
---

## Tech Stack

| Layer    | Teknologi                                        |
| -------- | ------------------------------------------------ |
| Bahasa   | [Go (Golang)](https://go.dev)                    |
| I/O      | [fmt](https://pkg.go.dev/fmt)                    |
| Struktur | Array statis, Struct                             |
| Output   | Command Line Interface (CLI)                     |

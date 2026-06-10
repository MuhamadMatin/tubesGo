<p align="center"><a href="https://go.dev" target="_blank"><img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg" width="300" alt="Go Logo"></a></p>
<h1 align="center">📁 Sistem Informasi Inventaris Dokumen Skripsi (SkripsIn)</h1>
<p align="center">
  Sistem Informasi Inventaris Dokumen Skripsi adalah aplikasi untuk mengelola arsip tugas
akhir mahasiswa secara digital. Data utama yang digunakan adalah data judul skripsi, data
penulis, dan data tahun lulus. Pengguna aplikasi adalah staf administrasi program studi atau
petugas perpustakaan.
</p>

---

## ⚠️ Penting
 
- Pastikan [Go](https://go.dev/dl) sudah terinstal di sistem Anda (versi **1.18** atau lebih baru)
- Aplikasi berjalan di **terminal / command line** — tidak memerlukan browser atau database
- Input yang mengandung **spasi** harus diganti dengan `_` (underscore), contoh: `ini_Budi`
- **Binary Search** hanya bekerja pada data yang **sudah diurutkan** terlebih dahulu
---
 
## 🚀 Instalasi
 
### 1. Clone repository
 
```bash
git clone https://github.com/MuhamadMatin/tubesGo.git
cd tubesGo
```
 
### 2. Verifikasi instalasi Go
 
```bash
go version
```
 
> Belum punya Go? Unduh di [https://go.dev/dl](https://go.dev/dl)
 
### 3. Jalankan program
 
```bash
go run .
```
 
### 4. Build menjadi executable (opsional)
 
```bash
# Build
go build -o skripsin main.go
 
# Jalankan (Linux / macOS)
./skripsin
 
# Jalankan (Windows)
skripsin.exe
```
 
Akses program melalui terminal setelah dijalankan.
 
---

 
## 📋 Data Dummy
 
Aplikasi sudah dilengkapi **10 data dummy** yang otomatis dimuat saat pertama kali dijalankan:
 
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

## 🛠️ Tech Stack

| Layer    | Teknologi                                        |
| -------- | ------------------------------------------------ |
| Bahasa   | [Go (Golang)](https://go.dev)                    |
| I/O      | [fmt](https://pkg.go.dev/fmt)                    |
| Tipe Data | Array statis + Struct                           |
| Antarmuka | CLI (Command Line Interface)                   |

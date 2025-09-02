# danantara

Selamat datang di proyek **danantara**!

## Deskripsi
danantara adalah aplikasi berbasis Go yang menggunakan framework Gin untuk web server dan GORM sebagai ORM untuk database.

## Prasyarat

Pastikan sudah meng-install:
- [Go](https://go.dev/doc/install) minimal versi 1.18
- Database (misal: MySQL, PostgreSQL, atau SQLite, sesuai konfigurasi GORM yang digunakan)

## Langkah Setup

### 1. Clone Repository
```bash
git clone https://github.com/hamas105projects/danantara.git
cd danantara
```

### 2. Instalasi Dependensi
Jalankan perintah berikut untuk mengunduh semua dependensi:
```bash
go mod tidy
```

### 3. Konfigurasi Database
Edit file konfigurasi database (misal di `.env` atau langsung di source code, sesuai implementasi).

Contoh konfigurasi (MySQL):
```
DB_USER=root
DB_PASS=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=danantara
SERVER_PORT=8080
```

### 4. Menjalankan Proyek

Jalankan aplikasi dengan:
```bash
go run main.go
```

Jika menggunakan file utama lain, sesuaikan dengan nama file entry point.

## Struktur Folder

- `main.go` : Entry point aplikasi
- `handler/` : Logic untuk menangani request
- `usecase/` : Logic semua data dan relasi database
- `entity/` : Struktur data dan ORM
- `routes/` : Routing aplikasi


## Kontribusi

1. Fork repository ini
2. Buat branch baru (`git checkout -b fitur-baru`)
3. Commit perubahan Anda (`git commit -m 'Menambahkan fitur X'`)
4. Push ke branch Anda (`git push origin fitur-baru`)
5. Ajukan pull request

## Lisensi

Cek file LICENSE untuk detail lisensi.

---

Jika ada pertanyaan, silakan buka issue di repository ini.
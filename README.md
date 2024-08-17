# Project api-example

Project ini merupakan contoh project untuk membuat REST API dengan menggunakan Golang. Terdapat beberapa teknologi yang digunakan dalam project ini, antara lain: Echo, Sqlx, dan Goose. Project ini menggunakan arsitektur gabungan antara Clean Architecture dan Domain Driven Design.

## Getting Started

Sebelum memulai, pastikan Anda telah menginstal Go pada sistem Anda. Jika belum, Anda dapat mengunduh dan menginstal Go dari [situs resmi Go](https://golang.org/dl/).

Selain Go, terdapat tools lain yang perlu diinstal, yaitu:
- [Bruno](https://www.usebruno.com/downloadshttps://github.com/go-bruno/bruno) (Posmant Alternative).
- Make, untuk menjalankan perintah di Makefile.
    - Windows: Bisa menggunakan [chocolatey](https://chocolatey.org/packages/make) `choco install make`
    - Linux: `sudo apt install make`
    - Mac: `brew install make`
- [Air](https://github.com/air-verse/air), untuk menjalankan server dengan hot reload.

### How To Run
1. Clone repository ini ke komputer Anda.
2. Buka terminal dan arahkan ke direktori project.
3. Jalankan perintah `go mod tidy` untuk menginstal dependensi.
4. Buat file `.env` dengan nilai-nilai seperti pada file `.env.example`
5. Ubah nilai pada file `.env` sesuai dengan konfigurasi Anda.
6. Jalankan perintah `make migrate-up` atau `go run cmd/migration/main.go up` untuk menjalankan migrasi database (pastikan database sudah tersedia).
7. Jalankan perintah `make run` atau `go run cmd/api/main.go` untuk menjalankan server.

## Project Structure
```
.
├── bruno
│   ├── environments
│   └── User
├── cmd
│   ├── api
│   └── migration
├── constants
├── internal
│   ├── domain
│   ├── repository
│   ├── server
│   └── service
├── migrations
├── pkg
│   ├── mySqlExt
│   └── util
└── tmp
```

### bruno
Folder ini berisi file-file configurasi untuk Bruno. Sebelum bisa menggunakan Bruno, kita perlu masuk ke folder bruno dan jalankan perintah `npm i` untuk menginstall dependensi. Jika tidak ingin menggunakan aplikasi Bruno GUI, kita bisa menggunakan bruno cli yang bisa kita install menggunakan perintah `npm i -g @usebruno/cli`.

### cmd
Folder yang berisi command-command yang bisa kita jalankan seperti `api` dan `migration`.

### constants
Folder yang berisi konstanta-konstanta yang digunakan di project ini.

### internal
Folder yang berisi semua logic dari project ini. Di dalam folder `domain` terdapat kumpulan structure yang digunakan untuk mendefinisikan structure data yang akan digunakan pada `handler`, `service`, dan `repository`.

Folder `repository` berisi package-package yang berurusan dengan service diluar project, seperti database, cache, dan juga aplikasi lainnya.

Folder `server` berisi konfigurasi http server, seperti `routing` dan `handlers`

Folder `service` berisi logic utama dari project ini.

### migrations
Folder yang berisi file-file migrasi database. Untuk membuat file migrasi baru kita bisa menjalankan perintah `go run cmd/migration/main.go create migration_file_name`.

### pkg
Folder yang berisi package-package yang digunakan di project ini, contohnya seperti function `utility` dan juga mysql extension.
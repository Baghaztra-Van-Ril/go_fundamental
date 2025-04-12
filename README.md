> pertemuan 6  
### Introduction and General Comparison  
# GoLang  

## Tugas:  
Studi kasus dengan kombinasi goroutine dan channel bahasa GO  

---

## 📌 Judul Studi Kasus  
**Pengambilan Data dari Banyak Situs Web secara Paralel**

---

## 🧩 Deskripsi Kasus

Sebuah aplikasi ingin mengambil konten dari banyak situs berita secara cepat (misalnya untuk keperluan agregator berita).  

Jika situs-situs tersebut diakses satu per satu, proses akan lama karena bergantung pada kecepatan masing-masing server.

Untuk mempercepat, kita bisa menjalankan permintaan ke semua situs secara **paralel menggunakan goroutine**, dan menyatukan hasilnya melalui **channel**.

---

## 🛠️ Solusi

- Gunakan goroutine untuk menjalankan request HTTP ke setiap situs secara paralel.  
- Gunakan channel untuk mengirim hasil response dari masing-masing goroutine ke fungsi utama.
- Fungsi utama akan menerima dari channel dan mengumpulkan hasilnya.

---

## 💡 Manfaat

- Jauh lebih cepat dibanding sequential.  
- Skalabel: bisa menambahkan banyak situs tanpa perubahan logika besar.  
- Tetap sederhana karena idiomatis dalam Go.

---

## 📁 Struktur Direktori

```
go_fundamental/
├── internal/
│   └── scraper/
│       └── client.go
├── service/
│   └── scrape_service.go
├── main.go
```

## Jalankan Aplikasi 
Inisialisasi **modul Go**
```bash
go mod init go_fundamental
go mod tidy
```
Jalankan proyek
```bash
go run main.go
```


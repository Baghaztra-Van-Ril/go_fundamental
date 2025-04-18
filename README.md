## 📚 **Pertemuan 6**  
### **Introduction and General Comparison**  
# **GoLang**  

## Tugas:  
Studi kasus dengan kombinasi goroutine dan channel bahasa Go  

---

## 📌 **Judul Studi Kasus**  
**Pengambilan Data dari Banyak Situs Web secara Paralel**

---

## 🧩 **Deskripsi Kasus**  

Sebuah aplikasi ingin mengambil konten dari banyak situs berita secara cepat (misalnya untuk keperluan agregator berita).  

Jika situs-situs tersebut diakses satu per satu, prosesnya akan memakan waktu lama karena bergantung pada kecepatan masing-masing server.

Untuk mempercepat proses ini, kita dapat menjalankan permintaan ke semua situs secara **paralel menggunakan goroutine**, dan menyatukan hasilnya melalui **channel**.

---

## 🛠️ **Solusi**

1. Gunakan **goroutine** untuk menjalankan request HTTP ke setiap situs secara paralel.
2. Gunakan **channel** untuk mengirim hasil response dari masing-masing goroutine ke fungsi utama.
3. Fungsi utama akan menerima data dari channel dan mengumpulkan hasilnya.

---

## 💡 **Manfaat**

- Proses yang **lebih cepat** dibandingkan eksekusi secara berurutan (sequential).
- **Skalabel**, dapat menambahkan lebih banyak situs tanpa mengubah banyak logika.
- **Sederhana** dalam bahasa Go.

---

## 📁 **Struktur Direktori**

```
go_fundamental/
├── internal/                    # Struktur dari GO untuk mencegak isinya diakses dari luar package
│   ├── scraper/
│   |   └── client.go            # Fungsi untuk mengambil konten dari situs
|   └── service/
│       └── scrape_service.go    # Logika pengolahan dan worker pool
├── main.go                      # Fungsi utama untuk menjalankan aplikasi
```

---

## 🏃‍♂️ **Jalankan Aplikasi**

1. **Inisialisasi modul Go** (jika belum):
```bash
go mod init go_fundamental
go mod tidy
```

2. **Jalankan proyek**:
```bash
go run main.go
```


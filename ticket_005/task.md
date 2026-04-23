### 🎫 TICKET: GO-005
**Title:** Implement Request Logger Middleware
**Story Points:** 2
**Objective:** Membuat *middleware* kustom menggunakan *Closures* untuk mencatat (*log*) setiap *request* HTTP yang masuk beserta waktu pemrosesannya.

**Acceptance Criteria:**
1. Buat fungsi baru bernama `LoggerMiddleware`. Fungsi ini harus menerima parameter `next` dengan tipe data `http.HandlerFunc`, dan *return value*-nya juga harus berupa `http.HandlerFunc`.
2. Di dalam fungsi yang dikembalikan (*closure*), catat waktu dimulainya *request* menggunakan `start := time.Now()`.
3. Izinkan *request* masuk ke "Brankas" dengan mengeksekusi fungsi aslinya: `next(w, r)`.
4. Segera setelah `next(w, r)` selesai, cetak log ke terminal menggunakan `fmt.Printf`. Log harus menampilkan: HTTP Method, URL Path yang diakses, dan durasi eksekusi.
   * *Contoh output terminal: `[POST] /orders - 1.2ms`*
   * *Hint: Durasi bisa didapatkan menggunakan `time.Since(start)`.*
5. Di dalam fungsi `main()`, modifikasi pendaftaran *routing* Anda dengan membungkusnya menggunakan *middleware* ini.
   * Formatnya menjadi: `http.HandleFunc("/orders", LoggerMiddleware(handleOrder))`

**PM Notes:**
* Pastikan Anda menambahkan package `"time"` pada blok `import` di bagian atas file.
* Pola penulisan untuk kriteria nomor 1 di Go biasanya terlihat seperti ini:
  ```go
  func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
      return func(w http.ResponseWriter, r *http.Request) {
          // Logika Anda di sini...
      }
  }
  ```

**Apakah instruksi di tiket ini sudah cukup jelas untuk langsung Anda kerjakan, atau Anda ingin mendiskusikan bagian struktur pengembalian fungsinya terlebih dahulu?**
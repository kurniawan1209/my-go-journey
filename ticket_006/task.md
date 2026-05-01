### 🎫 TICKET: GO-006
**Title:** Refactor Codebase into Packages (Models & Handlers)
**Story Points:** 3
**Objective:** Memecah `main.go` yang monolitik menjadi struktur folder/package yang modular dan *scalable*.

**Acceptance Criteria:**
1.  Buat struktur folder baru di dalam proyek Anda seperti ini:
    ```text
    /order-api
    ├── main.go
    ├── /models
    │   └── order.go
    └── /handlers
        └── order_handler.go
    ```
2.  **Package Models (`/models/order.go`):** * Deklarasikan `package models` di baris paling atas.
    * Pindahkan definisi struct `Order`, `OrderResponse`, dan `OrderDB` ke file ini.
    * Pindahkan inisialisasi `var globalOrderDB = &OrderDB{...}` ke sini. Namun, karena variabel ini akan diakses oleh *handler* nanti (di luar folder *models*), Anda **wajib** mengubah huruf pertamanya menjadi kapital (misal: `GlobalOrderDB`).
    * Pindahkan *method* `Set` dan `Get` ke sini. Pastikan nama *method* tetap diawali huruf kapital.
3.  **Package Handlers (`/handlers/order_handler.go`):**
    * Deklarasikan `package handlers` di baris paling atas.
    * Pindahkan fungsi `handleOrder` dan `handleGetOrder` ke sini. 
    * Ubah nama fungsinya menjadi diawali huruf kapital (misal: `HandleOrder` dan `HandleGetOrder`) agar bisa dipanggil oleh `main.go`.
    * Di dalam file ini, Anda harus mengimpor folder `models` Anda (contoh impor: `"order-api/models"` jika nama modul Anda `order-api`). Panggil *database* menggunakan `models.GlobalOrderDB`.
4.  **Package Main (`main.go`):**
    * File ini sekarang hanya bertugas sebagai titik kumpul (*Entry Point*).
    * Sisakan fungsi `main()` dan `LoggerMiddleware` di sini.
    * Impor package `handlers` Anda. Daftarkan rute dengan memanggil *handler* dari *package* tersebut: `http.HandleFunc("/orders", LoggerMiddleware(handlers.HandleOrder))`.

**Catatan PM:** Ini adalah murni tugas *refactoring*. **Tidak boleh ada logika bisnis yang berubah.** Pastikan server tetap berjalan di port 8086 dan post/get order tetap berfungsi seperti biasa setelah pemecahan file.

Silakan kerjakan pemecahan arsitektur ini. Karena ada 3 file berbeda, Anda bisa menyalin isi masing-masing file (`main.go`, `models/order.go`, dan `handlers/order_handler.go`) dan menempelkannya di bawah untuk saya *review*!
### 🎫 TICKET: GO-004
**Objective:** Melakukan *upgrade* pada *Mutex* untuk mendukung *concurrent reads*, serta membuat endpoint baru untuk mengambil data pesanan berdasarkan ID.

**Acceptance Criteria (AC):**
1. **Upgrade Mutex:** Di dalam struct `OrderDB`, ubah tipe data `mu` dari `sync.Mutex` menjadi `sync.RWMutex`.
2. **Method Get:** Tambahkan *method* baru pada `OrderDB` di bawah *method* `Set`:
   ```go
   func (db *OrderDB) Get(id string) (Order, bool) {
       // Gunakan db.mu.RLock() dan defer db.mu.RUnlock()
       // Ambil data dari map: order, ok := db.orders[id]
       // Return order dan ok
   }
   ```
3. **Handler Baru:** Buat fungsi *handler* bernama `handleGetOrder(w http.ResponseWriter, r *http.Request)`.
   * Tolak semua *request* selain `http.MethodGet` (kembalikan HTTP 405).
   * Ambil ID dari parameter URL: `id := r.URL.Query().Get("id")`.
   * Panggil `globalOrderDB.Get(id)`.
   * Jika parameter *boolean* mengembalikan `false`, kirimkan respons JSON dengan HTTP 404 (Not Found).
   * Jika `true`, kirimkan respons JSON dengan HTTP 200 (OK) dan tampilkan data order tersebut.
4. **Routing:** Di dalam fungsi `main()`, daftarkan *handler* baru ini ke *routing* khusus: `http.HandleFunc("/order", handleGetOrder)`.

**Catatan PM:** Fokus pada penggunaan `RLock()` (Read Lock) di dalam *method* `Get`. Ini akan memastikan bahwa saat ada request GET yang masuk, mereka bisa membaca data secara bersamaan tanpa saling memblokir, selama tidak ada request POST yang sedang menulis data.

Silakan modifikasi kode Anda sesuai dengan instruksi di atas dan berikan seluruh kodenya di sini untuk saya *review* dan *merge*.
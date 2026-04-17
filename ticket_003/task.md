### 🎫 TICKET: GO-003
**Title:** In-Memory Transactional Storage with Mutex
**Story Points:** 3 
**Objective:** Menyimpan data `Order` yang tervalidasi ke dalam *Map* global secara aman untuk menghindari *race condition* saat API menerima *traffic* tinggi.

**Acceptance Criteria:**
1. Hapus baris duplikat `json.NewDecoder` dari tiket sebelumnya.
2. Buat sebuah variabel global berupa *Map* dengan nama `orderDB` untuk menyimpan data pesanan (Key: `string` dari `ID`, Value: `Order`).
3. Buat sebuah variabel global `sync.Mutex` dengan nama `dbMutex`.
4. Di dalam fungsi `handleOrder` (setelah JSON divalidasi), gunakan `dbMutex.Lock()` sebelum menyimpan `body` ke dalam `orderDB`.
5. Pastikan Anda menggunakan `dbMutex.Unlock()` segera setelah data berhasil ditulis ke *Map*.

**PM Notes & Constraints:**
* **Mengapa Mutex?** *Map* di Go **tidak** *thread-safe* secara default. Jika dua *request* masuk di milidetik yang sama dan mencoba menulis ke *Map* tanpa *Lock*, aplikasi Go Anda akan langsung *panic* dan *crash*. *Mutex* (Mutual Exclusion) menjamin hanya satu proses yang boleh menulis ke *database* bohongan kita dalam satu waktu.
* **Hints to Google:** Cari tahu tentang `go map string struct` dan `golang sync.Mutex lock unlock`.

Silakan hapus baris yang salah tadi, tambahkan implementasi *database* di memori ini, dan berikan kode terbarunya untuk *review* Sprint 2.
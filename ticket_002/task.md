### 🎫 TICKET: GO-002
**Title:** Handle POST Request & Parse JSON Payload
**Story Points:** 2
**Objective:** Menerima data pesanan (order) masuk, memvalidasi method HTTP, dan melakukan *decoding* JSON ke dalam struct Go.

**Acceptance Criteria:**
1. Buat struct baru bernama `Order` yang memiliki 3 field: `ID` (string), `CustomerName` (string), dan `TotalAmount` (float64). Berikan JSON tags yang sesuai.
2. Buat endpoint baru di `/orders`.
3. Endpoint ini **hanya** boleh menerima request dengan method `POST`. Jika ada yang menembak menggunakan `GET`, kembalikan HTTP Status 405 (Method Not Allowed).
4. Tangkap payload JSON dari request body, parse (decode) ke dalam struct `Order`.
5. Kembalikan response JSON yang berisi pesan sukses dan menampilkan kembali data order yang barusan ditangkap.

**PM Notes & Constraints:**
* **Method Checking:** Gunakan kondisi `if r.Method != http.MethodPost` untuk memfilter.
* **Decoding JSON:** Jangan dibaca sebagai raw string. Gunakan `json.NewDecoder(r.Body).Decode(&namavariabel)`. Perhatikan penggunaan pointer `&` di sini, ini krusial di Go untuk memodifikasi variabel yang sudah ada.

Silakan kerjakan dan berikan kodenya di sini. Uji menggunakan Postman atau cURL dengan payload JSON sebelum Anda serahkan.
# 🎫 E-Ticketing API - Golang & PostgreSQL

Sistem backend e-ticketing sederhana untuk transportasi publik.  
Mendukung fitur login admin, input terminal, serta skema check-in & check-out menggunakan kartu prepaid.

---

## 🚀 Fitur Utama

- 🔐 **Login Admin** (JWT Authentication)
- 🏁 **Input Terminal** 
- 🎫 **Simulasi Check-In dan Check-Out** kartu
- 💾 **Database PostgreSQL**
- 🧱 Struktur proyek modular (controllers, models, middleware)

---

## 🛠️ Teknologi

- **Golang** (Gin Framework)
- **PostgreSQL**
- **JWT Authentication**
- **DBeaver / pgAdmin** (opsional, untuk pengelolaan DB)

---

## 📁 Struktur Folder

e-ticketing-api/
├── main.go
├── .env
├── go.mod
│
├── controllers/
│ ├── auth.go
│ └── terminal.go
│
├── middleware/
│ └── jwt.go
│
├── database/
│ └── db.go
├── postman/
│ └── e-ticketing.postman_collection.json

## 📊 Jaringan Online mode dan OFFline mode

💻Penjelasan Saat Ada Jaringan (Online Mode)
- Penumpang tap kartu di Gate Check-in.
- Gate mengirim data ke server (terminal_id, user_id, timestamp).
- Server mencatat entry, menunggu check-out.
- Saat check-out, server hitung tarif berdasarkan jarak terminal check-in & check-out.
- Saldo di database dikurangi → respons sukses → gate terbuka.
- Riwayat transaksi disimpan di log.

💻Penjelasan Saat Tidak Ada Jaringan (Offline Mode)
Solusi Offline:
- Setiap Gate Validator menyimpan:
- Cache entry (Check-in/Check-out)
- Saldo terakhir kartu (cached)
- Transaksi dicatat lokal dan disinkronkan ke server saat internet kembali tersedia.
- Conflict Resolution: Menggunakan timestamp + ID unik transaksi.



---

# Yard Planning System

Sistem sederhana untuk mengatur posisi container pada yard dengan logika slot–row–tier. Project ini menyediakan endpoint untuk **suggestion**, **placement**, dan **pickup** container.

---

## 📁 Struktur Project

```
yard-planning/
├── main.go          # Entry point: start server & route handler
├── db.go            # Koneksi database
├── model.go         # Struct model: Yard, Block, YardPlan
├── suggestion.go    # Endpoint /suggestion untuk mencari posisi kosong
├── placement.go     # Endpoint /placement untuk menempatkan container
└── pickup.go        # Endpoint /pickup untuk pickup container
```

---

## 🗄 Struktur Database

### Tabel:

* **yards**
  Menyimpan data yard.

* **blocks**
  Menyimpan data blok (relasi melalui `yard_id`).

* **yard_planning**
  Menyimpan lokasi container: slot, row, tier, container info, dan status pickup.

---

# 🔍 Suggestion Endpoint

### **POST** `/suggestion`

Digunakan untuk mencari posisi kosong pada yard berdasarkan urutan:

1. Saya menggunakan Pendekatan Agar Semua Slot Terisi semua terlebih dahulu Ascending dimulai dari 1. Misal Slot 1, Row 1, Tier 1
2. Jika slot terisi semua, lalu lanjut ke Row, sama Mulai dr 1 dan seterusnya Jadi misal Slot terakhir, Row 1 terisi semua lalu ke Row 2
3. Lalu jika semua Slots terisi lalu row terisi semua juga maka Tier naik ke 2 dan seterusnya

### Contoh Request

```json
{
  "yard": "YRD1",
  "container_number": "ALFI000002",
  "container_size": 20,
  "container_height": 8.6,
  "container_type": "DRY"
}
```

### Contoh Response

```json
{
  "suggested_position": {
    "block": "LC01",
    "row": 1,
    "slot": 1,
    "tier": 1
  }
}
```

---

# 📦 Placement Endpoint

### **POST** `/placement`

Digunakan untuk menempatkan container pada posisi tertentu.

### Contoh Request

```json
{
  "yard": "YRD1",
  "container_number": "ALFI000004",
  "block": "LC01",
  "slot": 3,
  "row": 2,
  "tier": 1,
  "container_size": 20
}
```

### Contoh Response

```json
{
  "message": "container placed successfully"
}
```

### Alur Logika

1. Akan Melakukan Pengecekan ke Table yard_plans sudah terisi belum
2. Jika Kosong maka langsung ter insert
3. Jika Terisi maka akan ter validasi

> NOTE : saya menambahkan container_size di req body, karena ada case jika container berukuran 40ft maka butuh 2 slot.

---

# 🚚 Pickup Endpoint

### **POST** `/pickup`

Digunakan untuk menandai bahwa container sudah diambil.

### Contoh Request

```json
{
  "yard": "YRD1",
  "container_number": "ALFI000002"
}
```

### Contoh Response

```json
{
  "message": "container picked up successfully"
}
```

### Logika

* Kolom `is_picked` menentukan status container.
* `true` → container sudah diambil
* `false` → container masih terpasang

> Catatan: Jika boleh mungkin Lebih baik menggunakan **yard_id** agar lebih unik daripada memakai `yard` code/name.

---

# 🛠 Instalasi & Menjalankan Server

1. Clone repository:

   ```bash
   git clone https://github.com/denizsetyawan/yard-planning-test.git
   ```
2. Import database dari file SQL.
3. Edit konfigurasi database di `db.go`.
4. Jalankan server:

   ```bash
   go run .
   ```
5. Uji endpoint menggunakan Postman / Thunder Client / cURL.

---

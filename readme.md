# Yard Planning
yard-planning/
├── main.go         # entry point, start server & route handler
├── db.go           # koneksi database
├── model.go        # struct model untuk table: Yard, Block, YardPlan
├── suggestion.go   # endpoint /suggestion (cari posisi kosong)
└── placement.go    # endpoint /placement untuk menempatkan container
└── pickup.go    # endpoint /pickup untuk pickup container

## Tabel
yards - untuk menampung data yards
blocks - untuk menampung data blok (terhubung ke yards by yard_id)
yard_planning - untuk menampung posisi container, slot, row, tier dll

## Suggestion

curl -X POST "http://localhost:8080/suggestion"

Contoh Body :
{
	"yard":  "YRD1",
	"container_number":  "ALFI000002",
	"container_size":  20,
	"container_height":  8.6,
	"container_type":  "DRY"
}

Contoh Response :
{
	"suggested_position":  {
		"block":  "LC01",
		"row":  1,
		"slot":  1,
		"tier":  1
	}
}

Alur :
1. Saya menggunakan Pendekatan Agar Semua Slot Terisi semua terlebih dahulu Ascending dimulai dari 1. Misal Slot 1, Row 1, Tier 1
2. Jika slot terisi semua, lalu lanjut ke Row, sama Mulai dr 1 dan seterusnya Jadi misal Slot terakhir, Row 1 terisi semua lalu ke Row 2
3. Lalu jika semua Slots terisi lalu row terisi semua juga maka Tier naik ke 2 dan seterusnya

## Placement

curl -X POST "http://localhost:8080/placement"

Contoh Body :
{
	"yard":  "YRD1",
	"container_number":  "ALFI000004",
	"block":  "LC01",
	"slot":  3,
	"row":  2,
	"tier":  1,
	"container_size":  20
}

Response :
{
	"message":  "container placed successfully"
}

Alur :
1. Akan Melakukan Pengecekan ke Table yard_plans sudah terisi belum
2. Jika Kosong maka langsung ter insert
3. Jika Terisi maka akan ter validasi

NOTE : saya menambahkan container_size di req body, karena ada case jika container berukuran 40ft maka butuh 2 slot.

## Pickup

curl -X POST "http://localhost:8080/pickup"

Contoh Body :
{
	"yard":  "YRD1",
	"container_number":  "ALFI000002"
}

Response :
{
	"message":  "container picked up successfully"
}

Alur :
1. jika di pickup di yard_plans ada kolom is_picked, true berarti sudah di pickup, jika false maka belum di pickup dan status masih ditempati container

NOTE : Jika bisa mengubah req body mungkin saya akan memilih memasukkan yard_id daripada menggunakan nama yard/code. agar lebih unique.

# Instalasi
1. Clone repository
2. Import database
3. Masuk ke db.go setting untuk konfigurasi database bisa di sesuaikan
4. jalankan via terminal di folder utama (yard-planning) run "go run ."
5. jalankan sesuai endpoint yang ingin diuji via postman atau lainnya
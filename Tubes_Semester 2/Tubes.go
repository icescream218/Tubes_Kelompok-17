package main

import (
	"fmt"
)

type kendaraan struct {
	ID          string
	Type        string
	waktuMasuk  int
	WaktuKeluar int
}

type Admin struct {
	Username string
	Password string
}

type ParkingAttendant struct {
	ID       string
	Name     string
	Password string
}

var parkingData = []kendaraan{}
var admins = []Admin{{Username: "admin1", Password: "pass123"}}
var parkingAttendants = []ParkingAttendant{}

func loginAdmin(username, password string) bool {
	for _, admin := range admins {
		if admin.Username == username && admin.Password == password {
			return true
		}
	}
	return false
}

func addParkingAttendant(id, name, password string) {
	length := len(parkingAttendants)
	parkingAttendants = make([]ParkingAttendant, length+1)
	parkingAttendants[length] = ParkingAttendant{ID: id, Name: name, Password: password}
}

func editParkingAttendant(id, namaBaru, passwordBaru string) {
	for i := 0; i < len(parkingAttendants); i++ {
		if parkingAttendants[i].ID == id {
			parkingAttendants[i].Name = namaBaru
			parkingAttendants[i].Password = passwordBaru
			return
		}
	}
}

func deleteParkingAttendant(id string) {
	for i := 0; i < len(parkingAttendants); i++ {
		if parkingAttendants[i].ID == id {
			copy(parkingAttendants[i:], parkingAttendants[i+1:])             // Pindahkan elemen satu indeks ke belakang.
			parkingAttendants = parkingAttendants[:len(parkingAttendants)-1] // Potong slice.
			return
		}
	}
}

func addkendaraan(id string, tipeKendaraan string, waktuMasuk int) {
	parkingData = append(parkingData, kendaraan{ID: id, Type: tipeKendaraan, waktuMasuk: waktuMasuk})
}

func editkendaraan(id string, newWaktuKeluar int) {
	for i := 0; i < len(parkingData); i++ {
		if parkingData[i].ID == id {
			parkingData[i].WaktuKeluar = newWaktuKeluar
			return
		}
	}
}

func calculateParkingbiaya(kendaraan kendaraan) int {
	durasi := kendaraan.WaktuKeluar - kendaraan.waktuMasuk
	biaya := durasi * 1000 // asumsi biaya per jam
	if durasi > 10 {
		biaya += 2000
	}
	return biaya
}

func displaykendaraansAndEarnings() {
	totalEarnings := 0
	for i := 0; i < len(parkingData); i++ {
		biaya := calculateParkingbiaya(parkingData[i])
		fmt.Printf("ID: %s, Type: %s, biaya: %d\n", parkingData[i].ID, parkingData[i].Type, biaya)
		totalEarnings += biaya
	}
	fmt.Printf("Total Earnings: %d\n", totalEarnings)
}

func main() {
	var username, password string
	fmt.Println("Masukkan username admin:")
	fmt.Scan(&username)
	fmt.Println("Masukkan password admin:")
	fmt.Scan(&password)

	if loginAdmin(username, password) {
		fmt.Println("Admin berhasil login.")
		for {
			fmt.Println("Pilih operasi:")
			fmt.Println("1. Tambah Petugas Parkir")
			fmt.Println("2. Edit Petugas Parkir")
			fmt.Println("3. Hapus Petugas Parkir")
			fmt.Println("4. Tambah Kendaraan")
			fmt.Println("5. Edit Waktu Keluar Kendaraan")
			fmt.Println("6. Tampilkan Kendaraan dan Pendapatan")
			fmt.Println("7. Keluar")
			var choice int
			fmt.Scan(&choice)

			switch choice {
			case 1:
				var id, name, password string
				fmt.Println("Masukkan ID petugas:")
				fmt.Scan(&id)
				fmt.Println("Masukkan nama petugas:")
				fmt.Scan(&name)
				fmt.Println("Masukkan password petugas:")
				fmt.Scan(&password)
				addParkingAttendant(id, name, password)
			case 2:
				var id, namaBaru, passwordBaru string
				fmt.Println("Masukkan ID petugas yang akan diubah:")
				fmt.Scan(&id)
				fmt.Println("Masukkan nama baru:")
				fmt.Scan(&namaBaru)
				fmt.Println("Masukkan password baru:")
				fmt.Scan(&passwordBaru)
				editParkingAttendant(id, namaBaru, passwordBaru)
			case 3:
				var id string
				fmt.Println("Masukkan ID petugas yang akan dihapus:")
				fmt.Scan(&id)
				deleteParkingAttendant(id)
			case 4:
				var id, tipeKendaraan string
				var waktuMasuk int
				fmt.Println("Masukkan ID kendaraan:")
				fmt.Scan(&id)
				fmt.Println("Masukkan tipe kendaraan:")
				fmt.Scan(&tipeKendaraan)
				fmt.Println("Masukkan waktu masuk (jam):")
				fmt.Scan(&waktuMasuk)
				addkendaraan(id, tipeKendaraan, waktuMasuk)
			case 5:
				var id string
				var WaktuKeluar int
				fmt.Println("Masukkan ID kendaraan untuk mengubah waktu keluar:")
				fmt.Scan(&id)
				fmt.Println("Masukkan waktu keluar (jam):")
				fmt.Scan(&WaktuKeluar)
				editkendaraan(id, WaktuKeluar)
			case 6:
				displaykendaraansAndEarnings()
			case 7:
				fmt.Println("Keluar dari program.")
				return
			default:
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			}
		}
	} else {
		fmt.Println("Login gagal!")
	}
}

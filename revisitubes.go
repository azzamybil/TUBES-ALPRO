package main

import (
	"fmt"
)

const NMAX = 20

type EPL struct {
	nama         string
	pertandingan int
	menang       int
	kalah        int
	seri         int
	golmasuk     int
	golkemasukan int
	selisihgol   int
	point        int
}

type tabEPL [NMAX]EPL

func tambahDataKlub(A *tabEPL, klubCount *int) {
	var nklub int
	fmt.Print("Masukkan jumlah klub yang akan ditambahkan: ")
	fmt.Scanln(&nklub)

	if *klubCount+nklub > NMAX {
		fmt.Println("Data klub melebihi kapasitas maksimum.")
		return
	}
	fmt.Print("Masukan nama klub: ")
	for i := 0; i < nklub; i++ {
		fmt.Scan(&A[*klubCount].nama)
		*klubCount++
	}
	fmt.Println("Klub berhasil ditambahkan")
}

// Pencarian dengan Sequential search
func findClubIndex(A *tabEPL, klubCount int, klubName string) int {
	for i := 0; i < klubCount; i++ {
		if A[i].nama == klubName {
			return i
		}
	}
	return -1
}

func IsiHasilPertandingan(A *tabEPL, klubCount int) {
	var minggu int
	var klub1, klub2 string
	var gol1, gol2 int

	fmt.Print("Masukan jumlah minggu pertandingan: ")
	fmt.Scanln(&minggu)

	if minggu < 1 || minggu > 38 {
		fmt.Println("Jumlah minggu harus antara 1 dan 38")
		return
	}
	for i := 1; i <= minggu; i++ {

		fmt.Printf("Minggu %d:\n", i)
		for j := 0; j < 10; j++ {

			fmt.Scan(&klub1, &gol1, &gol2, &klub2)

			var index1, index2 int
			index1 = findClubIndex(A, klubCount, klub1)
			index2 = findClubIndex(A, klubCount, klub2)

			index1 = -1
			index2 = -1

			for k := 0; k < klubCount; k++ {
				if A[k].nama == klub1 {
					index1 = k
				} else if A[k].nama == klub2 {
					index2 = k
				}
			}

			if index1 == -1 || index2 == -1 {
				fmt.Println("Nama klub tidak ditemukan")
				return
			}
			A[index1].pertandingan++
			A[index2].pertandingan++

			A[index1].golmasuk += gol1
			A[index2].golmasuk += gol2

			A[index1].golkemasukan += gol2
			A[index2].golkemasukan += gol1

			A[index1].selisihgol = A[index1].golmasuk - A[index1].golkemasukan
			A[index2].selisihgol = A[index2].golmasuk - A[index2].golkemasukan

			if gol1 > gol2 {
				A[index1].menang++
				A[index1].point += 3
				A[index2].kalah++
			} else if gol1 < gol2 {
				A[index2].menang++
				A[index2].point += 3
				A[index1].kalah++
			} else {
				A[index1].seri++
				A[index1].point++
				A[index2].seri++
				A[index2].point++
			}
		}
	}
	fmt.Println("Hasil Pertandingan berhasil diisi")
}

func sorting(A *tabEPL, klubCount int) {

	// pengurutan klasemen berdasakarkan poin dan selish gol dengan selection sort

	var temp EPL
	var idx int

	for i := 1; i < klubCount-1; i++ {
		idx = i - 1
		for j := i + 1; j < klubCount; j++ {
			if A[idx].point < A[j].point || (A[idx].point == A[j].point && A[idx].selisihgol < A[j].selisihgol) {
				idx = j
			}
		}
		temp = A[i-1]
		A[i-1] = A[idx]
		A[idx] = temp
	}

}

func tampilkanperingkatklub(A tabEPL, klubCount int) {
	// header table

	fmt.Println(" No.| Klub | T | M | K | S | GM | GK | SG | Poin|")
	fmt.Println("----|------|---|---|---|---|----|----|----|-----|")
	// Penampilan data klub
	for i := 0; i < klubCount; i++ {
		fmt.Printf("%-4d| %-5s| %-2d| %-2d| %-2d| %-2d| %-3d| %-3d| %-3d| %-5d|\n",
			i+1, A[i].nama, A[i].pertandingan, A[i].menang, A[i].seri, A[i].kalah,
			A[i].golmasuk, A[i].golkemasukan, A[i].selisihgol, A[i].point)

	}
}

func tampilkanMenu() {
	fmt.Println("EPL Manager Menu")
	fmt.Println("1. Tambah Data Klub")
	fmt.Println("2. Ubah Data Klub")
	fmt.Println("3. Hapus Data Klub")
	fmt.Println("4. Isi Hasil Pertandingan")
	fmt.Println("5. Tampilkan Peringkat Klub")
	fmt.Println("6. Keluar")
	fmt.Print("Pilih menu: ")
}

// func ubahdataklub() {

// }

// func hapusDataKlub() {

// }

func main() {
	var pilihan int
	var klub tabEPL
	var klubCount int

	for {
		tampilkanMenu()
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			tambahDataKlub(&klub, &klubCount)
		case 2:
			// ubahDataKlub() //  lu kerjain brian
		case 3:
			// hapusDataKlub() // lu kerjain brian
		case 4:
			IsiHasilPertandingan(&klub, klubCount)
			sorting(&klub, klubCount)
		case 5:
			tampilkanperingkatklub(klub, klubCount)
		case 6:
			fmt.Println("Terima kasih telah menggunakan EPL Manager.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

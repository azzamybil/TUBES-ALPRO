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

type Ptd struct {
	club1 string
	club2 string
	gol1  int
	gol2  int
}

type tabPtd [10]Ptd

type Minggu struct {
	pertandingan tabPtd
}

type tabMinggu [38]Minggu

func tambahDataKlub(A *tabEPL, klubCount *int) {
	var nklub int
	fmt.Print("Masukkan jumlah klub yang akan ditambahkan: ")
	fmt.Scanln(&nklub)

	if *klubCount+nklub > NMAX {
		fmt.Println("Data klub melebihi kapasitas maksimum.")
		return
	} else if *klubCount+nklub <= NMAX {
		fmt.Println("Masukan nama klub: ")
		for i := 0; i < nklub; i++ {
			fmt.Scanln(&A[*klubCount].nama)
			*klubCount++
		}
	} else {
		fmt.Print("input salah, coba lagi")
		return
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

func IsiHasilPertandingan(A *tabEPL, B *tabMinggu, klubCount int, mingguCount *int, pertandinganCount *int) {
	var minggu int
	var klub1, klub2 string
	var gol1, gol2 int

	fmt.Print("Masukan jumlah minggu pertandingan: ")
	fmt.Scanln(&minggu)

	if minggu < 1 || minggu > 38 {
		fmt.Println("Jumlah minggu harus antara 1 dan 38")
		return
	}
	for i := 0; i < minggu; i++ {
		if *mingguCount >= 38 {
			fmt.Println("Jumlah minggu melebihi kapasitas maksimum.")
			return
		}
		fmt.Printf("\nMinggu %d:\n", *mingguCount+1)
		for j := 0; j < 10; j++ {
			fmt.Scanln(&klub1, &gol1, &gol2, &klub2)
			index1 := findClubIndex(A, klubCount, klub1)
			index2 := findClubIndex(A, klubCount, klub2)

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
			B[*mingguCount].pertandingan[j].club1 = A[index1].nama
			B[*mingguCount].pertandingan[j].club2 = A[index2].nama
			B[*mingguCount].pertandingan[j].gol1 = gol1
			B[*mingguCount].pertandingan[j].gol2 = gol2
			*pertandinganCount++
		}
		*mingguCount++
	}
	fmt.Println("Hasil Pertandingan berhasil diisi")
}

func SelectionDescending(A *tabEPL, klubCount int) {
	// Pengurutan klasemen berdasarkan poin dan selisih gol dengan selection sort

	var temp EPL
	var idx int

	for i := 0; i < klubCount-1; i++ {
		idx = i
		for j := i + 1; j < klubCount; j++ {
			if A[idx].point < A[j].point || (A[idx].point == A[j].point && A[idx].selisihgol < A[j].selisihgol) {
				idx = j
			}
		}
		temp = A[i]
		A[i] = A[idx]
		A[idx] = temp
	}

}

func InsertionAscending(A *tabEPL, klubCount int) {
	for i := 1; i < klubCount; i++ {
		j := i
		for j > 0 {
			if A[j-1].point > A[j].point || (A[j-1].point == A[j].point && A[j-1].selisihgol > A[j].selisihgol) {
				A[j-1], A[j] = A[j], A[j-1]
			}
			j--
		}
	}
}

func tampilkanPeringkatKlub(A tabEPL, klubCount int) {

	fmt.Println(" No.| Klub | P | M | S | K | GM | GK | SG | Poin|")
	fmt.Println("----|------|---|---|---|---|----|----|----|-----|")
	// Penampilan data klub
	for i := 0; i < klubCount; i++ {
		fmt.Printf("%-4d| %-5s| %-2d| %-2d| %-2d| %-2d| %-3d| %-3d| %-3d| %-4d|\n",
			i+1, A[i].nama, A[i].pertandingan, A[i].menang, A[i].seri, A[i].kalah,
			A[i].golmasuk, A[i].golkemasukan, A[i].selisihgol, A[i].point)
	}
}

func tampilkanPertandingan(B tabMinggu, mingguCount int) {
	for i := 0; i < mingguCount; i++ {
		fmt.Println("\nMinggu ", i+1, ":")
		for j := 0; j < 10; j++ {
			if B[i].pertandingan[j].club1 == "" && B[i].pertandingan[j].club1 == "" {

			} else {
				fmt.Println(B[i].pertandingan[j].club1, B[i].pertandingan[j].gol1, B[i].pertandingan[j].gol2, B[i].pertandingan[j].club2)
			}
		}
	}
}

func ubahDataKlub(A *tabEPL, klubCount int) {
	var klubName string
	var pertandingan, menang, kalah, seri, gol, kebobolan int

	fmt.Print("Masukkan nama klub yang akan diubah hasil pertandingannya: ")
	fmt.Scanln(&klubName)

	index := findClubIndex(A, klubCount, klubName)
	if index == -1 {
		fmt.Println("Nama klub tidak ditemukan")
		return
	}

	fmt.Print("Masukkan jumlah pertandingan baru: ")
	fmt.Scanln(&pertandingan)
	A[index].pertandingan = pertandingan

	fmt.Print("Masukkan jumlah menang baru: ")
	fmt.Scanln(&menang)
	A[index].menang = menang

	fmt.Print("Masukkan jumlah kalah baru: ")
	fmt.Scanln(&kalah)
	A[index].kalah = kalah

	fmt.Print("Masukkan jumlah seri baru: ")
	fmt.Scanln(&seri)
	A[index].seri = seri

	fmt.Print("Masukkan jumlah gol baru: ")
	fmt.Scanln(&gol)
	A[index].golmasuk = gol

	fmt.Print("Masukkan jumlah kebobolan baru: ")
	fmt.Scanln(&kebobolan)
	A[index].golkemasukan = kebobolan

	
	A[index].selisihgol = A[index].golmasuk - A[index].golkemasukan

	
	A[index].point = (A[index].menang * 3) + A[index].seri

	fmt.Println("Hasil pertandingan klub berhasil diubah")
}

func ubahPertandingan(A *tabEPL, B *tabMinggu, klubCount int, mingguCount int) {
	var minggu, pertandingan, gol1, gol2 int
	var klub1, klub2 string

	fmt.Print("Masukkan Minggu berapa data yang ingin diubah (1-38): ")
	fmt.Scanln(&minggu)

	if minggu < 1 || minggu > mingguCount {
		fmt.Println("Minggu tidak valid.")
		return
	}

	fmt.Print("Masukkan Pertandingan ke berapa data yang ingin diubah (1-10): ")
	fmt.Scanln(&pertandingan)

	if pertandingan < 1 || pertandingan > 10 {
		fmt.Println("Pertandingan tidak valid.")
		return
	}

	fmt.Print("Masukkan nama club pertama: ")
	fmt.Scanln(&klub1)
	fmt.Print("Masukkan gol club pertama: ")
	fmt.Scanln(&gol1)
	fmt.Print("Masukkan nama club kedua: ")
	fmt.Scanln(&klub2)
	fmt.Print("Masukkan gol club kedua: ")
	fmt.Scanln(&gol2)

	index1 := findClubIndex(A, klubCount, klub1)
	index2 := findClubIndex(A, klubCount, klub2)

	if index1 == -1 || index2 == -1 {
		fmt.Println("Nama klub tidak ditemukan")
		return
	}

	PertanadinganLama := B[minggu-1].pertandingan[pertandingan-1]
	idx1 := findClubIndex(A, klubCount, PertanadinganLama.club1)
	idx2 := findClubIndex(A, klubCount, PertanadinganLama.club2)

	A[idx1].pertandingan--
	A[idx2].pertandingan--

	A[idx1].golmasuk -= PertanadinganLama.gol1
	A[idx1].golkemasukan -= PertanadinganLama.gol2

	A[idx2].golmasuk -= PertanadinganLama.gol2
	A[idx2].golkemasukan -= PertanadinganLama.gol1

	A[idx1].selisihgol = A[idx1].golmasuk - A[idx1].golkemasukan
	A[idx2].selisihgol = A[idx2].golmasuk - A[idx2].golkemasukan

	if PertanadinganLama.gol1 > PertanadinganLama.gol2 {
		A[idx1].menang--
		A[idx1].point -= 3
		A[idx2].kalah--
	} else if PertanadinganLama.gol1 < PertanadinganLama.gol2 {
		A[idx2].menang--
		A[idx2].point -= 3
		A[idx1].kalah--
	} else {
		A[idx1].seri--
		A[idx1].point--
		A[idx2].seri--
		A[idx2].point--
	}

	A[index1].pertandingan++
	A[index2].pertandingan++
	A[index1].golmasuk += gol1
	A[index1].golkemasukan += gol2
	A[index2].golmasuk += gol2
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

	B[minggu-1].pertandingan[pertandingan-1].club1 = klub1
	B[minggu-1].pertandingan[pertandingan-1].club2 = klub2
	B[minggu-1].pertandingan[pertandingan-1].gol1 = gol1
	B[minggu-1].pertandingan[pertandingan-1].gol2 = gol2

	fmt.Println("Pertandingan berhasil diubah")
}

func hapusDataKlub(A *tabEPL, klubCount *int) {
	var klubName string
	fmt.Print("Masukkan nama klub yang akan dihapus: ")
	fmt.Scanln(&klubName)

	index := findClubIndex(A, *klubCount, klubName)
	if index == -1 {
		fmt.Println("Nama klub tidak ditemukan")
		return
	}

	for i := index; i < *klubCount-1; i++ {
		A[i] = A[i+1]
	}
	A[*klubCount-1] = EPL{}
	*klubCount--
	fmt.Println("Data klub berhasil dihapus")
}
func hapusDataPertandingan(A *tabEPL, B *tabMinggu, klubCount *int, mingguCount *int, pertandinganCount *int) {
	var minggu, pertandingan int

	fmt.Print("Masukkan Minggu berapa data yang ingin dihapus (1-38): ")
	fmt.Scanln(&minggu)

	if minggu < 1 || minggu > *mingguCount {
		fmt.Println("Minggu tidak valid.")
		return
	}

	fmt.Print("Masukkan Pertandingan ke berapa data yang ingin dihapus (1-10): ")
	fmt.Scanln(&pertandingan)

	if pertandingan < 1 || pertandingan > 10 {
		fmt.Println("Pertandingan tidak valid.")
		return
	}

	minggu -= 1
	pertandingan -= 1

	klub1 := B[minggu].pertandingan[pertandingan].club1
	klub2 := B[minggu].pertandingan[pertandingan].club2
	gol1 := B[minggu].pertandingan[pertandingan].gol1
	gol2 := B[minggu].pertandingan[pertandingan].gol2

	index1 := findClubIndex(A, *klubCount, klub1)
	index2 := findClubIndex(A, *klubCount, klub2)

	if index1 == -1 || index2 == -1 {
		fmt.Println("Nama klub tidak ditemukan")
		return
	}

	A[index1].pertandingan--
	A[index2].pertandingan--

	A[index1].golmasuk -= gol1
	A[index2].golmasuk -= gol2

	A[index1].golkemasukan -= gol2
	A[index2].golkemasukan -= gol1

	A[index1].selisihgol = A[index1].golmasuk - A[index1].golkemasukan
	A[index2].selisihgol = A[index2].golmasuk - A[index2].golkemasukan

	if gol1 > gol2 {
		A[index1].menang--
		A[index1].point -= 3
		A[index2].kalah--
	} else if gol1 < gol2 {
		A[index2].menang--
		A[index2].point -= 3
		A[index1].kalah--
	} else {
		A[index1].seri--
		A[index1].point--
		A[index2].seri--
		A[index2].point--
	}

	B[minggu].pertandingan[pertandingan] = Ptd{}

	*pertandinganCount--
	fmt.Println("Data pertandingan berhasil dihapus")
}

func findMax(A tabEPL, klubCount int) {
	var idx int
	for i := 1; i < klubCount; i++ {
		if A[i].point > A[i-1].point || (A[i].point == A[i-1].point && A[i].selisihgol > A[i-1].selisihgol) {
			idx = i
		}
	}
	fmt.Println("Peringkat tertnggi adalah : ")
	fmt.Printf("%-4d| %-5s| %-2d| %-2d| %-2d| %-2d| %-3d| %-3d| %-3d| %-4d|\n", 1, A[idx].nama, A[idx].pertandingan, A[idx].menang, A[idx].seri, A[idx].kalah, A[idx].golmasuk, A[idx].golkemasukan, A[idx].selisihgol, A[idx].point)

}

func tampilkanMenu() {
	fmt.Println("===================================")
	fmt.Println("|        EPL Manager Menu         |")
	fmt.Println("===================================")
	fmt.Println("|  1. Tambah Data Klub            |")
	fmt.Println("|  2. Ubah Data Klub              |")
	fmt.Println("|  3. Ubah Data pertandingan      |")
	fmt.Println("|  4. Hapus Data Klub             |")
	fmt.Println("|  5. Hapus Data Pertandingan     |")
	fmt.Println("|  6. Isi Hasil Pertandingan      |")
	fmt.Println("|  7. Tampilkan Pertandingan      |")
	fmt.Println("|  8. Peringkat Descending        |")
	fmt.Println("|  9. Peringkat Ascending         |")
	fmt.Println("| 10. Pemenang Juara satu EPL	  |")
	fmt.Println("| 11. Keluar                      |")
	fmt.Print("Pilih Menu: ")
}

func main() {
	var pilihan int
	var klub tabEPL
	var tanding tabMinggu
	var klubCount, mingguCount, pertandinganCount int

	for {
		tampilkanMenu()
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			tambahDataKlub(&klub, &klubCount)
		case 2:
			ubahDataKlub(&klub, klubCount)
		case 3:
			ubahPertandingan(&klub, &tanding, klubCount, mingguCount)
		case 4:
			hapusDataKlub(&klub, &klubCount)
		case 5:
			hapusDataPertandingan(&klub, &tanding, &klubCount, &mingguCount, &pertandinganCount)
		case 6:
			IsiHasilPertandingan(&klub, &tanding, klubCount, &mingguCount, &pertandinganCount)
		case 7:
			tampilkanPertandingan(tanding, mingguCount)
		case 8:
			SelectionDescending(&klub, klubCount)
			tampilkanPeringkatKlub(klub, klubCount)
		case 9:
			InsertionAscending(&klub, klubCount)
			tampilkanPeringkatKlub(klub, klubCount)
		case 10:
			findMax(klub, klubCount)
		case 11:
			fmt.Println("Terima kasih telah menggunakan EPL Manager.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

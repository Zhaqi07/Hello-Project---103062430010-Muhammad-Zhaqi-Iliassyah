// Hello Project - 103062430010 Muhammad Zhaqi Iliassyah

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Musik struct {
	Judul, Artis, Genre string
	Rating              int
}

const maxMusik = 100

var koleksi [maxMusik]Musik
var jumlah int
var input = bufio.NewReader(os.Stdin)

func main() {
	for {
		fmt.Printf("\n1. Lihat Musik")
		fmt.Printf("\n2. Tambah Musik")
		fmt.Printf("\n3. Ubah Musik")
		fmt.Printf("\n4. Hapus Musik")
		fmt.Printf("\n5. Cari Musik")
		fmt.Printf("\n6. Urutkan Musik")
		fmt.Printf("\n7. Keluar")

		fmt.Println()

		fmt.Println("Silahkan Pilih Opsi: ")
		pilihan := baca()
		if pilihan == "1" {
			lihat()
		} else if pilihan == "2" {
			tambah()
		} else if pilihan == "3" {
			ubah()
		} else if pilihan == "4" {
			hapus()
		} else if pilihan == "5" {
			cari()
		} else if pilihan == "6" {
			urut()
		} else if pilihan == "7" {
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func baca() string {
	s, _ := input.ReadString('\n')
	return strings.TrimSpace(s)
}

func bacaInt(pesan string) int {
	for {
		fmt.Print(pesan)
		n, err := strconv.Atoi(baca())
		if err == nil {
			return n
		}
		fmt.Println("Input harus angka.")
	}
}

func lihat() {
	if jumlah == 0 {
		fmt.Println("Belum ada musik.")
		return
	}
	for i := 0; i < jumlah; i++ {
		m := koleksi[i]
		fmt.Printf("%d. %s - %s [%s] %d\n", i+1, m.Judul, m.Artis, m.Genre, m.Rating)
	}
} // maman

func tambah() {
	if jumlah >= maxMusik { // buat ngecek jumlah musik yg ada udah mencapai batas maksimum apa blm.
		fmt.Println("Koleksi penuh.") // kalo udah penuh nanti print ini yg bakal keluar di output.
		return // langsung keluar dari fungsi karna datanya udah mencapai maksimum.
	}
	fmt.Print("Judul: ")
	j := baca() // manggil fungsi baca buat ngebaca input dari user dan disimpan ke variabel j.
	fmt.Print("Artis: ")
	a := baca() // manggil fungsi baca buat ngebaca input dari user dan disimpan ke variabel a.
	fmt.Print("Genre: ")
	g := baca() // manggil fungsi baca buat ngebaca input dari user dan disimpan ke variabel g.
	r := bacaInt("Rating: ") // menampilkan teks "Rating" agar user mengisi angka rating lagu. Nilai yang diketik oleh user akan dibaca sebagai angka (integer) dan disimpan ke variabel r.
	koleksi[jumlah] = Musik{j, a, g, r} // buat menyimpan lagu baru ke dalam daftar koleksi di posisi berikutnya. Lagu baru dibuat dari input yang sudah diketik user, yaitu judul, artis, genre, dan rating. Setelah disimpan, jumlah koleksi lagu (jumlah) akan bertambah.
	jumlah++ // abis masukkin data barunya, jumlah++ ini buat nambahin jumlah musik di variabel koleksi sebanyak 1.
	fmt.Println("Musik ditambahkan!")
}

func ubah() {
	lihat() // buat nampilin semua data musik yang udh disimpen biar si user bisa liat dan milih nomor lagu yg pengen di ubah
	n := bacaInt("Nomor musik: ") - 1 // minta input buat masukkin nomor lagu yg pengen dirubah, karna array mulai dari indeks 0, jadi dikurangin 1 supaya input dari si user sesuai sama indeks array nya.
	if n < 0 || n >= jumlah { // buat ngevalidasi kalau user salah masukkin input contohnya kaya terlalu besar atau terlalu kecil, nanti bakal keluar outputnya "nomor tidak ada" dan keluar dari fungsi.
		fmt.Println("Nomor tidak ada.")
		return
	}
	fmt.Print("Judul baru (kosong=tidak diubah): ") // kalo kita ga ngetik apa-apa trs langsung enter, gaada judul lagu yang diubah.
	j := baca()
	if j != "" {
		koleksi[n].Judul = j
	}
	fmt.Print("Artis baru: ") // ini juga sama tapi ini bedanya buat nama artis.
	a := baca()
	if a != "" {
		koleksi[n].Artis = a
	}
	fmt.Print("Genre baru: ") // ini juga sama, kalo ini buat genre
	g := baca()
	if g != "" {
		koleksi[n].Genre = g
	}
	r := baca() // minta input rating yang baru, kalo user masukkin angka nanti di cek dlu, bisa di konversi ke angka (Atoi) apa ngga, sama angka yg dimasukkin valid apa ngga sama ketentuan dari si inputnya, kalo valid rating dirubah, kalo ga valid ratingnya tetep sama kaya yg sebelumnya.
	if r != "" {
		if ri, err := strconv.Atoi(r); err == nil && ri >= 1 && ri <= 5 {
			koleksi[n].Rating = ri
		}
	}
	fmt.Println("Musik diubah.")
} // intinya kalo input baru yg diminta ga di isi, data yg lama yg udh kesimpen ga berubah.

func hapus() {
	lihat() // Menampilkan daftar musik yang sudah ada agar user tahu nomor musik mana yang ingin dihapus.
	n := bacaInt("Nomor musik: ") - 1 // Meminta user memasukkan nomor musik yang ingin dihapus, lalu dikurangi 1 karena array dimulai dari indeks 0.
	if n < 0 || n >= jumlah { // Validasi: Kalau nomor tidak valid (di luar batas data yang ada), maka tampilkan pesan error dan keluar dari fungsi.
		fmt.Println("Nomor tidak ada.")
		return
	}
	for i := n; i < jumlah-1; i++ { // Proses menggeser semua data setelah musik yang dihapus ke depan, agar lubang (kosong) akibat penghapusan tertutup.
		koleksi[i] = koleksi[i+1]
	}
	jumlah-- // Kurangi jumlah data musik karena 1 sudah dihapus.
	fmt.Println("Musik dihapus.")
}

func cari() {
	fmt.Print("Cari berdasarkan: 1. Artis  2. Rating\nPilih: ")
	p := baca()
	if p == "1" {
		fmt.Print("Masukkan nama artis: ") // User diminta memasukkan nama artis. Nama tersebut diubah menjadi huruf kecil agar pencarian tidak case sensitive.
		a := strings.ToLower(baca())
		for i := 0; i < jumlah; i++ { // Pencarian dilakukan satu per satu (linear search) terhadap semua data, Jika ada nama artis yang mengandung kata yang dicari, maka akan ditampilkan.
			if strings.Contains(strings.ToLower(koleksi[i].Artis), a) {
				m := koleksi[i]
				fmt.Printf("%s - %s [%s] %d\n", m.Judul, m.Artis, m.Genre, m.Rating)
			}
		}
	} else if p == "2" {
		selectionSortRatingAsc() // Pertama-tama, data musik diurutkan dulu dari rating terkecil ke terbesar. Ini penting karena kita akan menggunakan binary search, yang hanya bisa digunakan pada data terurut.
		r := bacaInt("Masukkan rating: ")
		i := binarySearchRating(r) //  User diminta memasukkan rating yang dicari, lalu dilakukan binary search untuk menemukan rating tersebut.
		if i != -1 { // Jika binary search menemukan datanya (i != -1), maka data musik tersebut ditampilkan, kalau tidak ketemu, akan muncul pesan "Tidak ditemukan."
			m := koleksi[i]
			fmt.Printf("Ditemukan: %s - %s [%s] %d\n", m.Judul, m.Artis, m.Genre, m.Rating)
		} else {
			fmt.Println("Tidak ditemukan.")
		}
	}
} // jadi kalo cari dari nama artisnya itu carinya manual pake strings.Contains, tapi kalo carinya berdasarkan rating, di urutin dulu baru pake binari search

func urut() {
	fmt.Print("Urutkan berdasarkan: 1. Rating  2. Artis\nPilih: ")
	k := baca() // hasil inputan user di simpen ke variabel k
	fmt.Print("Urutan: 1. Asc  2. Desc\nPilih: ") // Asc itu naik, Desc itu turun
	u := baca() // hasil inputan user di simpen ke variabel u

	if k == "1" { // berdasarkan rating
		if u == "1" {
			selectionSortRatingAsc() // kalo pilihnya Asc nanti fungsi ini yang bakal dipanggil
		} else {
			selectionSortRatingDesc() // kalo pilihnya Desc nanti fungsi ini yang di panggil
		}
	} else {
		if u == "1" { // berdasarkan artis
			insertionSortArtisAsc()
		} else {
			insertionSortArtisDesc()
		}
	}
	fmt.Println("Data diurutkan:")
	lihat() // buat nampilin daftar musik yang udah di urutin ke user
} // jaki

func selectionSortRatingAsc() {
	for i := 0; i < jumlah-1; i++ {
		min := i
		for j := i + 1; j < jumlah; j++ {
			if koleksi[j].Rating < koleksi[min].Rating {
				min = j
			}
		}
		koleksi[i], koleksi[min] = koleksi[min], koleksi[i]
	}
}

func selectionSortRatingDesc() {
	for i := 0; i < jumlah-1; i++ {
		max := i
		for j := i + 1; j < jumlah; j++ {
			if koleksi[j].Rating > koleksi[max].Rating {
				max = j
			}
		}
		koleksi[i], koleksi[max] = koleksi[max], koleksi[i]
	}
}

func insertionSortArtisAsc() {
	for i := 1; i < jumlah; i++ {
		key := koleksi[i]
		j := i - 1
		for j >= 0 && koleksi[j].Artis > key.Artis {
			koleksi[j+1] = koleksi[j]
			j--
		}
		koleksi[j+1] = key
	}
}

func insertionSortArtisDesc() {
	for i := 1; i < jumlah; i++ {
		key := koleksi[i]
		j := i - 1
		for j >= 0 && koleksi[j].Artis < key.Artis {
			koleksi[j+1] = koleksi[j]
			j--
		}
		koleksi[j+1] = key
	}
}

func binarySearchRating(target int) int {
	low, high := 0, jumlah-1
	for low <= high {
		mid := (low + high) / 2
		if koleksi[mid].Rating == target {
			return mid
		} else if koleksi[mid].Rating < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// sapik

package challenge4

import (
	"fmt"
	"os"
)

type Teman struct {
	NoAbsen      int
	Nama         string
	Alamat       string
	Pekerjaan    string
	AlasanGabung string
}

var daftarTeman = []Teman{
	{
		NoAbsen:      1,
		Nama:         "Andi",
		Alamat:       "Jakarta",
		Pekerjaan:    "Mahasiswa",
		AlasanGabung: "Ingin mempelajari bahasa pemrograman Golang",
	},
	{
		NoAbsen:      2,
		Nama:         "Budi",
		Alamat:       "Bandung",
		Pekerjaan:    "Programmer",
		AlasanGabung: "Ingin meningkatkan skill di bidang pemrograman",
	},
	{
		NoAbsen:      3,
		Nama:         "Cindy",
		Alamat:       "Surabaya",
		Pekerjaan:    "Designer",
		AlasanGabung: "Ingin mempelajari bahasa pemrograman yang baru",
	},
}

func Tugas() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <no_absen>")
		return
	}

	noAbsen := args[1]
	for _, teman := range daftarTeman {
		if fmt.Sprintf("%d", teman.NoAbsen) == noAbsen {
			fmt.Printf("Nama: %s\n", teman.Nama)
			fmt.Printf("Alamat: %s\n", teman.Alamat)
			fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
			fmt.Printf("Alasan Gabung: %s\n", teman.AlasanGabung)
			return
		}
	}
	fmt.Println("Teman tidak ditemukan")
}

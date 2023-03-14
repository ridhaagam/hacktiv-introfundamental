package challenge3

import (
	"fmt"
	"log"
)

func Tugas() {
	fmt.Println()
	log.Println("Tugas Challenge Tiga")
	fmt.Println()

	str := "selamat malam"
	count(str)
}

func count(str string) {
	counts := make(map[string]int)

	for _, c := range str {
		counts[string(c)]++
		fmt.Printf("%c\n", c)
	}
	fmt.Println(counts)
}

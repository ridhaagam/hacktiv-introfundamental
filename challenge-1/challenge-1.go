package challenge1

import (
	"fmt"
	"log"
)

func Tugas() {
	fmt.Println()
	log.Println("Tugas Challenge Satu")
	fmt.Println()

	// menampilkan nilai i : 21 fmt.Printf("%T \n", i) // contoh : fmt.Printf("%v \n", i)
	i := 21
	fmt.Printf("%v \n", i)

	// menampilkan tipe data dari variabel i
	fmt.Printf("%T \n", i)

	// menampilkan tanda %
	fmt.Println("%")

	// menampilkan nilai boolean j : true
	j := true
	fmt.Printf("%v \n", j)

	// spacing //
	fmt.Println()

	// menampilkan nilai boolean j : true
	if j {
		fmt.Printf("%b \n", 21)
	}

	// menampilkan unicode russia : Я (ya)
	fmt.Printf("%c\n", 'Я')

	// menampilkan nilai base 10 : 21 menampilkan nilai base 8 :25
	base := 21
	fmt.Printf("%d\n", base)
	fmt.Printf("%o\n", base)

	// menampilkan nilai base 16 : f
	value := 15
	fmt.Printf("%x\n", value)

	// menampilkan nilai base 16 : F 13
	value1 := 15
	value2 := 19
	fmt.Printf("%X %X\n", value1, value2)

	// menampilkan unicode karakter Я : U+042F var k float64 = 123.456;
	char := 'Я'
	codePoint := '\u042F'
	fmt.Printf("%c %U\n", char, codePoint)

	// spacing //
	fmt.Println()

	// menampilkan float : 123.456000
	k := 123.456
	fmt.Printf("%f\n", k)

	// menampilkan float scientific : 1.234560E+02
	s := 123.456
	fmt.Printf("%e\n", s)
}

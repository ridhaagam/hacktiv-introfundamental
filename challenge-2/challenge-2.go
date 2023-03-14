package challenge2

import (
	"fmt"
	"log"
	"unicode/utf8"
)

func Tugas() {
	fmt.Println()
	log.Println("Tugas Challenge Dua")
	fmt.Println()

	// nilai i
	for i := 0; i < 5; i++ {
		fmt.Printf("Nilai i = %d\n", i)
	}

	// nilai j
	for j := 0; j < 5; j++ {
		fmt.Printf("Nilai j = %d\n", j)
	}

	// str to character
	str := "C_A_ШA_P_B_O"

	for i := 0; i < len(str); {
		// Get the next rune and its byte length
		r, size := utf8.DecodeRuneInString(str[i:])

		if r == utf8.RuneError {
			// Handle invalid UTF-8
			fmt.Printf("invalid UTF-8 byte sequence at index %d\n", i)
			i++
			continue
		}

		// Print the Unicode code point and byte position of the character
		if str[i] != '_' {
			fmt.Printf("character U+%04X '%c' starts at byte position %d\n", r, r, i)
		}

		// Move to the next character
		i += size
	}

	// nilai j
	for j := 6; j < 11; j++ {
		fmt.Printf("Nilai j = %d\n", j)
	}
}

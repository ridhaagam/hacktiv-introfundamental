package challenge5

import (
	"fmt"
	"sync"
)

var (
	data1 = []string{"bisa1", "bisa2", "bisa3"}
	data2 = []string{"coba1", "coba2", "coba3"}

	// Inisialisasi mutex
	mutex sync.Mutex

	// Inisialisasi wait group
	wg sync.WaitGroup
)

func Tugas() {
	// printNormal()
	printMutex()
}

func printNormal() {
	wg.Add(8)

	for i := 1; i <= 4; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(data1, i)
		}(i)
	}

	for i := 1; i <= 4; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(data2, i)
		}(i)
	}

	wg.Wait()
}

func printMutex() {
	wg.Add(8)

	for i := 1; i <= 4; i++ {
		go func(i int) {
			defer wg.Done()

			mutex.Lock()
			fmt.Println(data1, i)
			defer mutex.Unlock()
		}(i)
	}

	for i := 1; i <= 4; i++ {
		go func(i int) {
			defer wg.Done()

			mutex.Lock()
			fmt.Println(data2, i)
			defer mutex.Unlock()
		}(i)
	}

	wg.Wait()
}

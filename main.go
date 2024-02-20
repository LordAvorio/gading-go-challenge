package main

import (
	"bufio"
	"fmt"
	"gading-go-challenge/database"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {

	for {

		lineMenu := strings.Repeat("x", 20)

		fmt.Println(lineMenu)
		fmt.Println("1. Ordered Data")
		fmt.Println("2. Unordered Data")
		fmt.Println("3. Exit")
		fmt.Println(lineMenu)
		fmt.Print("Input Angka Menu : ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		errScanner := scanner.Err()

		if errScanner != nil {
			panic(errScanner)
		}

		switch scanner.Text() {
		case "1":
			handleOrderedData()
		case "2":
			handleUnorderedData()
		case "3":
			fmt.Println("PROGRAM EXIT")
			os.Exit(0)
		default:
			fmt.Println("Tidak ada Pilihan Menu")
		}

	}
}

func handleOrderedData() {
	counterData1 := 0
	counterData2 := 0

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		mutex.Lock()
		go printDataOrdered(database.ListSampleData["SampleData1"], &counterData1, &mutex, &wg)
		mutex.Lock()
		go printDataOrdered(database.ListSampleData["SampleData2"], &counterData2, &mutex, &wg)
	}
	wg.Wait()
}

func handleUnorderedData() {
	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go printDataUnordered(database.ListSampleData["SampleData1"], i, &wg)
		go printDataUnordered(database.ListSampleData["SampleData2"], i, &wg)
	}
	wg.Wait()
}

func printDataUnordered(data any, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(data, index)
}

func printDataOrdered(data any, counter *int, mtx *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	*counter += 1
	fmt.Println(data, *counter)
	mtx.Unlock()
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	loopLength := 0

	for {
		fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
		fmt.Println("x    Please enter your number    x")
		fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

		fmt.Printf("Number : ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		errScan := scanner.Err()
		if errScan != nil {
			panic(errScan)
		}

		inputScanner := scanner.Text()
		numberScanner, errConvert := strconv.Atoi(inputScanner)

		if errConvert != nil || numberScanner <= 0 {
			fmt.Println("Input must be a positive number and higher than 0 !")
			fmt.Println("")
		} else {
			loopLength = numberScanner
			break
		}
	}

	fmt.Println("")
	fmt.Println("Result = ")

	for i := 1; i <= loopLength; i++ {

		checkSquare := math.Sqrt(float64(i))
		checkCube := math.Cbrt(float64(i))

		if checkSquare == math.Trunc(checkSquare) && checkCube == math.Trunc(checkCube) {
			fmt.Println("Squarecube")
		} else if checkSquare == math.Trunc(checkSquare) {
			fmt.Println("Square")
		} else if checkCube == math.Trunc(checkCube) {
			fmt.Println("Cube")
		} else {
			fmt.Println(i)
		}
	}

}

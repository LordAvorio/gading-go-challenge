package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	stringValueMap := map[string]int{}
	lineSeparation := strings.Repeat("*", 25)

	for {
		fmt.Println(lineSeparation)
		fmt.Println("     Word Separation")
		fmt.Println(lineSeparation)
		fmt.Printf("Enter your sentence here : ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		errScanner := scanner.Err()
		if errScanner != nil {
			panic(errScanner)
		}

		stringSplit := strings.Split(scanner.Text(),"")

		if len(stringSplit) == 0 {
			fmt.Println("Input cannot be empty !")
		}else{
			for _, valueString := range stringSplit {
				fmt.Println(valueString)

				value, isValueExist := stringValueMap[valueString]
				if isValueExist {
					stringValueMap[valueString] = value + 1
				}else{
					stringValueMap[valueString] = 1
				}
			}
			fmt.Println(stringValueMap)
			break
		}
	}

}
package main

import (
	"fmt"
	"gading-go-challenge/database"
	"os"
	"strconv"
	"strings"
)

func main() {

	isDataAvailable := checkDataAvalaibility()

	if !isDataAvailable {
		fmt.Println("Data students kosong, silahkan isi data terlebih dahulu.")
	} else {
		inputValue := os.Args[1:]
		handlingSearchData(inputValue)
	}

}

func checkDataAvalaibility() bool {
	result := false
	if len(database.ListOfStudents) > 0 {
		result = true
	}
	return result
}

func handlingSearchData(input []string) {

	isSearchByName := false
	inputName := ""
	result := []database.StudentBiodata{}

	if len(input) == 0 {
		fmt.Println("Tolong masukkan nama atau nomor absen")
	} else {

		inputId, errInteger := strconv.Atoi(input[0])
		if errInteger != nil {
			isSearchByName = true
		}

		if isSearchByName {
			inputName = strings.Join(input[:], " ")
			for _, value := range database.ListOfStudents {
				if value.Name == inputName {
					result = append(result, value)
				}
			}
		} else {
			if len(input) > 1 {
				fmt.Println("Format ID salah")
			} else {
				for _, value := range database.ListOfStudents {
					if value.Id == inputId {
						result = append(result, value)
					}
				}
			}
		}

		if len(result) > 0 {
			for _, valueResult := range result {
				fmt.Printf("ID		: %d \n", valueResult.Id)
				fmt.Printf("Nama		: %s \n", valueResult.Name)
				fmt.Printf("Alamat		: %s \n", valueResult.Address)
				fmt.Printf("Pekerjaan	: %s \n", valueResult.Job)
				fmt.Printf("Alasan		: %s \n", valueResult.Reason)
			}
			fmt.Println(result)
		}else{
			fmt.Println("Data dengan nama/absen tersebut tidak tersedia")
		}
	}
}

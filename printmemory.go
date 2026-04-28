package main 

import "fmt"

func PrintMemory(arr [10]byte) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%02x", arr[i])

		if i != len(arr) - 1 {
		fmt.Print(" ")
		}
	}
	fmt.Println()

	for i := 0; i < len(arr); i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			fmt.Printf("%c", arr[i])
		}else{
			fmt.Print(".")
		}
	}
	fmt.Println()	

}
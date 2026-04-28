package main 

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	for i := 0; i < len(arr); i++ {
		PrintHex(arr[i])

		if i != len(arr)-1{
			z01.PrintRune(' ')
		}
		z01.PrintRune(rune(arr[i]))
		z01.PrintRune('\n')
	}
	z01.PrintRune('\n')
	
}

func PrintHex(b byte) {
	base := "0123456789abcdef"
	z01.PrintRune(rune(base[b/16]))
	z01.PrintRune(rune(base[b%16]))
}
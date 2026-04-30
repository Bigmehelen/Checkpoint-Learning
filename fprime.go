package main

// import (
// 	"fmt"
// 	 "os" 
// 	)

// func Atoi(s string) int {

// 	num := 0
// 	for i:= 0; i < len(s); i++ {
// 		if s[i] < '0' || s[i] > '9' {
// 			return 0
// 		}
// 		num = num * 10 + int(s[i] - '0')
// 	}
// 	return num
// }

// func main() {

// 	if len(os.Args) != 2 {
// 		return
// 	}

// 	n := Atoi(os.Args[1])
// 	first := true
// 	div := 2

// 	for n > 1 {
// 		if n % div == 0{
// 			if !first {
// 				fmt.Print("*")
// 			}
// 			fmt.Print(div)
// 			first = false
// 			n /=div
// 		}
// 		div++
// 	}
// 	fmt.Println()

// }
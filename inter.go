package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	if len(os.Args) != 3 {
// 		return
// 	}

// 	s1 := os.Args[1]
// 	s2 := os.Args[2]

// 	seen := make(map[byte]bool)

// 	for i := 0; i < len(s1); i++ {
// 		c := s1[i]

// 		if seen[c] {
// 			continue
// 		}

// 		for j := 0; j < len(s2); j++ {
// 			if c == s2[j] {
// 				fmt.Print(string(c))
// 				seen[c] = true
// 				break
// 			}
// 		}
// 	}
// 	fmt.Println()		
	
// }
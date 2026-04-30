package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 2 {
		return
	}
	
	for i := 1; i < len(os.Args); i++ {
		s := os.Args[i]
		word := ""

		for j := 0; j < len(s); j++ {
			c := s[j]

			if c >= 'A' && c <= 'Z' {
				c += 32
			}
			

			if (j+1 == len(s) || s[j+1] == ' ') && 
			c >= 'a' && c <= 'z' {
				c -= 32
			}
			word += string(c)

		}
		fmt.Println(word)
	}

}
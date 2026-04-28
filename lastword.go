package main

func LastWord (s string) string {
	result := ""
	n := len(s)
	i := n - 1
	
	if n == 0 {
		return "\n"
	}

	for i >= 0 && s[i] == ' ' {
		i--
	}

	for i >= 0 && s[i] != ' ' {
		i--
	}

	end := i 
	start := i + 1


	for char := start; char <= end; char++ {
		result += string(s[char])
	}
	return result + "\n"

}	
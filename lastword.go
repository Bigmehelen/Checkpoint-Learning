package main

func LastWord (s string) string {
	result := ""
	n := len(s)
	i := n - 1
	end := i 
	start := i + 1

	if n == 0 {
		return "\n"
	}

	for i >= 0 && s[i] == ' ' {
		i--
	}

	for i >= 0 && s[i] != ' ' {
		i--
	}

	for char := start; char <= end; char++ {
		result += string(s[char])
	}
	return result + "\n"

}	
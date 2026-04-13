package main

func CountAlpha(str string) int {
	count := 0
	for i := 0; i < len (str); i++ {
		if (str[i] >= 'a' && str[i] <= 'z') || str[i] >= 'A' && str[i] <= 'Z' {
			count++
		}
	}
	return count
}
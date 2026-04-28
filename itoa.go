package main

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""
	sign := "-"

	if n < 0 {
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = string(rune('0' + digit)) + result
		n /= 10
	}
	return sign + result
}
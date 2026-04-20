package main

func HashCode(s string) string {
	n := len(s)
	result := make([]byte, n)

	for i := 0; i < n; i++ {
		hash := (int(s[i]) + n) % 127

		if hash < 32 {
			hash += 32
		}
		result[i] = byte(hash)

	}
	return string(result) 	
} 
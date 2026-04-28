package main
func UniqueCount(s1, s2 string) int {
	if s1 == "" && s2 == "" {
		return -1
	}

	m1 := make(map[rune]bool)
	m2 := make(map[rune]bool)

	for _, ch := range s1 {
		m1[ch] = true
	}

	for _, ch := range s2 {
		m2[ch] = true
	}

	count := 0

	for ch := range m1 {
		if !m2[ch] {
			count++
		}
	}

	for ch := range m2 {
		if !m1[ch] {
			count++
		}
	}

	return count
}	
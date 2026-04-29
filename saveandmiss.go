package main

func SaveAndMiss(s string, n int) string {
	res := ""
	isSave := true

	for i := 0; i < len(s); i+=n {
		end := i + n

		if end > len(s) {
			end = len(s)
		}

		if isSave {
			res = res + s[i:end]
		}
		isSave = !isSave

	}
	return res + "\n"
}

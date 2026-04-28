package main

func PrintRange(start, end int) string {
	if start < 0 || end < 0 || start > 99 || end > 99 {
		return "Invalid\n"
	}
	result := ""
	for i := start; i <= end; i++ {
		if i != start {
			result += ","
		}

		if i < 10 {
			result += "0" + string(rune('0' + i))
		}else{
			tens := i / 10
			units := i % 10
			result += string(rune('0' + tens)) + string(rune('0' + units))
		}
		
	}
	return result + "\n"
}
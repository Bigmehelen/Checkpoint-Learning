package main

func CamelToSnakeCase(s string) string{

	if len(s) == 0 {
		return ""
	}

	if !isValidCamelCase(s) {
		return s
	}

	var result []rune
	for i, char := range s {
		if i > 0 && char >= 'A' && char <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, toLower(char))
	}
	return string(result)
	
}


func isValidCamelCase(s string) bool {

	for _, char := range s {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
			return false 
		}
	}
	
	LastChar := rune(s[len(s)-1])
	if LastChar >= 'A' && LastChar <= 'Z' {
		return false
	}

	for i := 0; i < len (s)-1; i++ {
		current := rune(s[i])
		next := rune(s[i+1])
		if current >= 'A' && current <= 'Z' && next >= 'A' && next <= 'Z' {
			return false
		}
	}
	return true
}

func toLower (char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + 32
	}
	return char
} 
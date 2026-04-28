package main

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	
	newword := true

	for i:= 0; i < len(s); i++ {
		if s[i] == ' ' {
			newword = true
			continue
		}

		if newword {
			if s[i] >= 'a' && s[i] <= 'z' {
				return false
			}
			newword = false
		}
	}
	return true	
} 
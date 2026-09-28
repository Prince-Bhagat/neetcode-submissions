import "unicode/utf8"


func validate(ch rune)bool{
	if (ch <= '9' && ch >= '0') || (ch >=  'a' && ch <= 'z') || (ch <= 'Z' && ch >= 'A'){
		return true
	}else {
		return false
	}
}
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start := 0
	end := utf8.RuneCountInString(s) -1
	runeSlice := []rune(s)

	for start <= end {
		
		if !validate(runeSlice[start]){
			start++
			continue
		}
		if !validate(runeSlice[end]){
			end--
			continue
		}

		
		if runeSlice[start] != runeSlice[end] {
			return false
		}
		start++
		end--
	}
	return true
}
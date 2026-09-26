func isPalindrome(s string) bool {
	l,r:=0,len(s)-1

	for l<r {
		// if l is still less than r and the character at position l is not alphanumeric increment l
		for l < r && !isAlphaNum(rune(s[l])){
			l++
		}

		// if r is still greater than l and the character at position r is not alphanumeric decrement r
		for r > l && !isAlphaNum(rune(s[r])){
			r--
		}

		// check if the characters are the same and if theyre not return
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
            return false
        }
        l++
        r--
	}

	return true
		
}

func isAlphaNum(c rune) bool {
    return unicode.IsLetter(c) || unicode.IsDigit(c)
}
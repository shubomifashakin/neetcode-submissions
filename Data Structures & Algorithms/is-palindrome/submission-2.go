func isPalindrome(s string) bool {
	// initialize 2 pointers
	// we get an element
	// if the element is not a valid number of letter, we move that forward
	// if both are, we compare both of them and if they are different we return immediately

	left:=0
	right:=len(s)-1

	// so far the left has not crossed into the right
	for left<right{
		// if left is still less than right and the character at the left position is not alphanumeric, increment left
		for left<right && !isAlphaNum(rune(s[left])) {
			left++
		}

		// if right is still greater than left and the character at right is not alphanumeric, decrement right
		for right>left && !isAlphaNum(rune(s[right])) {
			right--
		}

		rightChar:=rune(s[right])
		leftChar:=rune(s[left])

		// if both characters are alphanumeric and theyre not the same, return false
		if(unicode.ToLower(rightChar)!=unicode.ToLower(leftChar)){
			return false
		}

		right--
		left++
	}	

	return true		
}

func isAlphaNum(c rune) bool {
    return unicode.IsLetter(c) || unicode.IsDigit(c)
}
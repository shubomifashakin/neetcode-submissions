func isPalindrome(s string) bool {
	// split the string into individual chars
	split:=strings.Split(strings.ToLower(s),"")
	clean:=[]string{}

	for _,char:=range split {
		// if it is not a letter and its not a digit, ignore it
		if(!unicode.IsLetter([]rune(char)[0]) && !unicode.IsDigit([]rune(char)[0])){
			continue
		}

		// if it is, push to array
		clean=append(clean,char)
	}

	cleanLen:=len(clean)
	// loop through the array and compare it with the back
	for idx:=range cleanLen {
		front:=clean[idx]
		back:=clean[cleanLen-1-idx]

		if front != back {
			return false
		}
	}

	return true
		
}

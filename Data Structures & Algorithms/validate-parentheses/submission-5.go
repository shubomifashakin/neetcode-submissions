func isValid(s string) bool {
	// we need to map the closing tags to the opening tags
	// we loop through the string, if we encounter an opening tag, we just append it to the stack
	// if we encounter a closing tag, we check if the previously added character was the correct opening tag
	mapped:=map[rune]rune{
		'}':'{',
		')':'(',
		']':'[',
	}

	stack:=[]rune{}

	// loop through the characters in the string
	for _,char:=range s {
		// check if the current element in view is a closing tage
		// if it is, check if the previously added element was an opening tag
		// if it was an opening tag, remove it from the stack
		expectedOpening,isClosing:=mapped[char]

		if(isClosing){
			// if the stack is empty then no opening tags exist is in the stack
			if(len(stack)==0){
				return false
			}else{
				// if its not empty then check if the previously added opening tag matches this one
				openingTag:=stack[len(stack)-1]

				
				if (expectedOpening!=openingTag){
					return false
				}else {
					stack=stack[0:len(stack)-1]
				}
			}
		}else{
			stack=append(stack,char)
		}
	}

	// if there is no element in the stack then its correct
	return len(stack)==0
}

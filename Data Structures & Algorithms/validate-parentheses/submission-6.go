func isValid(s string) bool {

	// but we have an expectation
	closingToOpening:=map[rune]rune{
		']':'[',
		'}':'{',
		')':'(',
	}

	// we loop through the string
	// for each element we see in the string, we first of all check if the character is a closing tag, if it is, then we check if the previous element is the respective closing tag
	// if its not then we know the string is incorrect
	stack:=[]rune{}

	for _,val:=range s {
		// check if the current element in view is a closing tag
		opening,isClosing:=closingToOpening[val]

		if(isClosing){
			// if its a closing tag and the stack is empty, then we know the strong is not correct because no prev values exist
			if (len(stack)==0){
				return false
			}else{
				// if prev value is not the corresponding opening tag, return false
				if (stack[len(stack)-1]!=opening){
					return false
				}

				// just remove that previous value from the stack
				stack=stack[0:len(stack)-1]
			}
		}else{
			// if its not a closing tag just append to the stack
			stack=append(stack,val)
		}
	}

	// if we have gone through everything and the stack is not empty, then not everything is closed properly
	return len(stack)==0
}

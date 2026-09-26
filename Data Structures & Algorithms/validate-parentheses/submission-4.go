func isValid(s string) bool {
	if len(s)%2==1{
		return false
	}

	mapped:=map[string]string{
		")":"(",
		"}":"{",
		"]":"[",
	}

	stack:=[]string{}

	// loop through the string
	for _,char:=range s {
		opening,currIsClosing:=mapped[string(char)]

		// if the current element in view is a closing parenthesis
		if currIsClosing {

		// if the stack is not empty and value at the top is the expected opening parenthesis, then we can remove it
			if len(stack)>0 && stack[len(stack)-1] == opening{
				stack=stack[0:len(stack)-1]
			}else{
				return false
			}

		}else{
			// the current element in view is not a closing tag, then we can just add to the stack
			stack=append(stack,string(char))
		}
		
	}

	
	return len(stack) == 0
}

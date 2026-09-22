func groupAnagrams(strs []string) [][]string {
	// we have a map which has an integer array as key and a slice of strings as value
	result:=map[[26]int][]string{}

	for _,str:=range strs {
		key:=[26]int{}

		for _,letter:=range str {
			// increment the character at that position 
			key[letter-'a']++
		}

		// find a key in the map that matches and append the str to it
	    arr,exists:=result[key]

		// if that key already exists, append the string to the slice
		if(exists){
			result[key]=append(arr,str)
			continue
		}
		// if it doesnt, just create the slice with that string as the first element
		result[key]=[]string{str}
	}

	actualRes:=[][]string{}

	for _,val:=range result{
		actualRes=append(actualRes,val)
	}

	return actualRes
}
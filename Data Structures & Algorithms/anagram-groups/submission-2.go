func groupAnagrams(strs []string) [][]string {
	// we are given an array of strings
	// we need to return an array of those string grouped by which are anagrams

	// we need to store a map of all anagrams grouped

	// we loop through the slice and we convert it to the letter index
	// thats what we use as the key of the map

	seen:=map[[26]int][]string{}

	// loop through each word (going to be a string)
	for _,word:=range strs {
		// for each value we get the lkey
		key:=[26]int{}

		// loop through each character in the word (a rune)
		for _,char:=range word {
			// convert each charcter to its alphabet index
			converted:=char-'a'
			// increment the value at that position
			key[converted]++
		}

		// check if that key exists in the map and if it does, we append the word to it
		group,exists:=seen[key]

		if (exists){
			group=append(group,word)

			seen[key]=group
		}else{
			seen[key]=[]string{word}
		}
	}

	result:=[][]string{}
	for _,group:=range seen{
		result=append(result,group)
	}

	return result
}
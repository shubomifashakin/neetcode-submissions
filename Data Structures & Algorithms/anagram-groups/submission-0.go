func groupAnagrams(strs []string) [][]string {
	existing:=map[string][]string{}

  for _,val:=range strs {
	// split the string and arrange it first
	ordered:=strings.Split(val,"")
	sort.Strings(ordered)
	joined:=strings.Join(ordered,"")

	arr,exists:=existing[joined]

	// if it exists, append the string to their array
		if(exists){
			arr=append(arr,val)
			existing[joined]=arr
			continue;
		}

	// if it does not, create the array
		existing[joined]=[]string{val}
	}

	//extract all the arrays from the map
	result:=[][]string{}

	for _,val:= range existing {
		result = append(result, val)
	}
	
	return result
}
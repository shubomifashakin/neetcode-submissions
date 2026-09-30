func longestConsecutive(nums []int) int {
	// an array of integers
	// return the length of the longest consecutive sequence

	// so they are basically asking for the amount of n+1 we can find in the element

	seen:=map[int]bool{}

	for _,num:=range nums{
		_,exists:=seen[num]

		if(exists){
			continue
		}

		seen[num]=true
	}

	// for each element in the array we check if its the start of a sequence
	// if its not the start of a sequence, we skip
	// if its the start of a sequence, we start counting how many n+1s we have for that element

	realCount:=0
	for _,num:=range nums{
		// check if its the start of a sequence
		_,exists:=seen[num-1]

		if(exists){
			continue
		}

		// its the start of a sequence
		// keep counting the amount of n+1s from the sequence
		count:=1
		for{
			// check if n+1 exists
			_,exists:=seen[num+count]

			// if n+1 exists, we increment the sequence count
			if(exists){
				count++
			}else{
				// if it does not, we have finished counting
				break
			}
		}

		if(count>realCount){
			realCount=count
		}
	}

	return realCount

}

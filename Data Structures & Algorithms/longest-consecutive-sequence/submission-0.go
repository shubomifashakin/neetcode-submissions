func longestConsecutive(nums []int) int {
	count:=0
	seen:=map[int]bool{}

	// initialize a hash set
	for _,num:=range nums{
		_,exists:=seen[num]
		if(exists){
			continue
		}

		seen[num]=true
	}

	for _,num:=range nums{
		// check if the element in view has a value thats 1 less than it in the set
		_,exists:=seen[num-1]

		if(exists){
			continue
		}

		innerCount:=1
		// if it does not then its the start of a sequence
		// keep checking if theres a sequence
		for {
			_,exists2:=seen[num+innerCount]
			
			if(exists2){
				innerCount++
			}else{
				break
			}
		}

		// if the old sequence is less than the new, then replace it
		if(innerCount>count){
			count=innerCount
		}
	}

	return count
}

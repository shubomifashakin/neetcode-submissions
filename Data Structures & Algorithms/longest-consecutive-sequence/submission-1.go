func longestConsecutive(nums []int) int {
	// create a set of all the unique numbers
	seen:=map[int]bool{}

	for _,num:=range nums {
		_,exists:=seen[num]
		if(exists){
			continue
		}

		seen[num]=true
	}

	count:=0

	// loop through each num and check if the number is the start of a sequence
	for _,num:=range nums {
		// check if the number is a start of a sequence
		_,exists:=seen[num-1]

		if(exists){
			continue
		}

		// then it is a start of a sequence
		sequenceCount:=1

		// keep checking if the next character exists in the set
		for{
			// check if next exists
			_,nextPresent:=seen[num+sequenceCount]

			if(nextPresent){
				sequenceCount++
			}else{
				break
			}
		}
		
		// if the new sequence count is greater than the old, replace the old
		if(sequenceCount>count){
			count=sequenceCount
		}
	}

	return count
	}


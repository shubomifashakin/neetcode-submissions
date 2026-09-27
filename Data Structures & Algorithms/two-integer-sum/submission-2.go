

func twoSum(nums []int, target int) []int{
	// we have a slice of numbers
	// we have a target value
	// two elements in the slice, need to add up to produce the target

	// we setup a hashmap, which records all elements we have seen
	seen:=map[int]int{}

	// we loop through the slice, we get the number that we need to add to the current element in view to produce the result
	// we then check if that number exists in the hashmap
	// if it does, we return a response

	for idx,num:=range nums {
		// get the number we need to add
		required:=target-num

		// check if it exists
		idxReq,exists:=seen[required]

		// if it exists, return both indexes
		if(exists){
			return []int{idxReq,idx}
		}

		// if it does not, record it
		seen[num]=idx
	}

	return []int{}
	}
func twoSum(numbers []int, target int) []int {
	left:=0
	right:=len(numbers)-1

	// so far left is still less than right, keep running
	for left < right {
		// get the result of the values
		result:=numbers[left]+numbers[right]

		// if the result is greater than the target then reduce the max possible value
		if (result>target){
			right--
			continue
		}

		// if the result is less than the target then increase the max possible value
		if(result<target){
			left++
			continue
		}

		return []int{left+1,right+1}
	}
	return []int{}
}

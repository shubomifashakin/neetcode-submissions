func twoSum(numbers []int, target int) []int {
	// the input array is sorted
	// meaning large numbers exist towards the right and
	// small numbers exists towards the end

	left:=0
	right:=len(numbers)-1

	// so far left is still less than right keep running
	for left<right{
		// calculate the sum
		sum:=numbers[left]+numbers[right]

		// if the sum is greater than the target then the upper limit is too high
		if(sum>target){
			right--
			continue
		}

		// if the sum is less than the target then the lower limit is too low
		if(sum<target){
			left++
			continue
		}

		return []int{left+1,right+1}
	}

	return []int{}
}

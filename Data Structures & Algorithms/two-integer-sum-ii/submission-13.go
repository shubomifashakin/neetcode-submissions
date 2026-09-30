func twoSum(numbers []int, target int) []int {
	// initialize a variable at the start
	// initialize a variable at the end
	start:=0
	end:=len(numbers)-1

	// run a loop, so far the start has not cross the end
	//on each iteration, check if the sum of the two is greater than the target
	// if it is, reduce the end by 1
	// if its smaller, increase the start by 1

	for start<end{
		sum:=numbers[start]+numbers[end]

		if(sum>target){
			end--
			continue
		}

		if(sum<target){
			start++
			continue
		}

		return []int{start+1,end+1}
	}

	return []int{}
}

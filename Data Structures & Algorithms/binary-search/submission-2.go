func search(nums []int, target int) int {
	// initialize 2 pointers
	start:=0
	end:=len(nums)-1


	// keep running the loop until theyve crosses each other
	for start<=end{
		// get the middle element
		mid:=(start+end)/2

		// if our target is greater than our middle element, then we shoudl only consider from mid+1 to end
		// else, if our target is less than our middle element, then we should only consider from start to mid-1
		// if our target is our middle element, then we return true

		if (target>nums[mid]){
			start=mid+1
		}

		if(target<nums[mid]){
			end=mid-1
		}

		if(target == nums[mid]){
			return mid
		}
	}

	return -1
}

func search(nums []int, target int) int {
	// given an array of length n
	// given a target

	// return the index of the target in the array
	// if not present return -1

	start:=0
	end:=len(nums)-1

	// keep running so far start is less than end
	// in each iteration, get the mid point
	// if the value at the mid point is greater than the target, only consider values from mid point +1 to end
	// if the value at the mid point is less than target, only consider values from start to mid -1
	// if the value at the mid point is the same as the target, return the mid point

	for (start<=end){
		mid:=(start+end)/2

		// if middle element is the target just return it
		if(nums[mid]==target){
			return mid
		}

		// if the middle element is greater or equal to start
		if(nums[mid]>=nums[start]){
			if (target> nums[mid] || target<nums[start]){
				start=mid+1
			}else{
				end=mid-1
			}
		}else{
			if(target<nums[mid] || target>nums[end]){
				end=mid-1
			}else{
				start=mid+1
			}
		}

	}

	return -1
}

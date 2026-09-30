func search(nums []int, target int) int {
	// initialize 2 pointers
	start:=0
	end:=len(nums)-1

	// keep the loop running so far the start and end are not at the same point
	for start<=end{
		// get the mid point of the current elements in view
		mid:=(start+end)/2

		// if the number at that mid point is the same as the target, return it
		if(nums[mid]==target){
			return mid
		}

		// if the number at that point is greater than the target, remove everything from that mid point till end from view
		if(nums[mid]>target){
			end=mid-1
		}

		// if the number at that point is less than the target, remove everything from the start all the way to that mid point from view
		if(nums[mid]<target){
			start=mid+1
		}
	}

	// if the element was not found, return -1
	return -1
}

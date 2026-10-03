func search(nums []int, target int) int {
	start:=0
	end:=len(nums)-1

	// keep running so far the start has not crossed the end
	for start<=end{
		midpoint:=(start+end)/2
		midEl:=nums[midpoint]

		if(midEl==target){
			return midpoint	
		}

	// get the sorted set
	// if the start is less than the midpoint, then the left is the sorted set
	if(nums[start]<=midEl){
		// if the target is less than the mid element and greater or equal to
		// the start, then the target is in this set
		// search from start to mid -1
		if(target>=nums[start]&& target<midEl){
			end=midpoint-1
		}else{
			start=midpoint+1
		}
	}else{
		if(target>midEl && target<=nums[end]){
			start=midpoint+1
		}else{
			end=midpoint-1
		}
	}
	}

	return -1
}
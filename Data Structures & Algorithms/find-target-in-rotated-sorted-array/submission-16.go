func search(nums []int, target int) int {
   start:=0
   end:=len(nums)-1


   for start<=end {
	// get the mid point
	mid:=(start+end)/2

	if (nums[mid]==target){
		return mid
	}

	// get the sorted set
	if(nums[start]<=nums[mid]){
		// if our target is less than or equal to our mid point
		// and greater or equal to our first element in this set
		// then our target is here
		if(target>=nums[start] && target<=nums[mid]){
			end=mid-1
		}else{
			start=mid+1
		}
	}else{
	// the right is the sorted set
	// if the target element is greater or equal to our mid point and less than or equal to the last element in the set, the our target is here
	if(target>=nums[mid] && target<=nums[end]){
		start=mid+1
	}else{
		end=mid-1
	}
	}
   }

   return -1
}
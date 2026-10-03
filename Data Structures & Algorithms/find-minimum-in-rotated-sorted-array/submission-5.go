func findMin(nums []int) int {
  start:=0
  end:=len(nums)-1

  for start<end{
    midpoint:=(start+end)/2
    midEl:=nums[midpoint]

    // if the mid element is greater than the end
    // then we know the smaller side is at the right
    // we need to check from mid+1 to end
    if(midEl>nums[end]){
      start=midpoint+1
    }else{
      end=midpoint
    }
  }

  return nums[start]
}
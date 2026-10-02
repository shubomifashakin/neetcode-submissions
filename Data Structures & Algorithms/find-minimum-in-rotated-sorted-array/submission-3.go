func findMin(nums []int) int {
   // set the start and end
   start:=0
   end:=len(nums)-1
    
    for start<end{
        // calculate the mid point
        mid:=(start+end)/2
        // if the element at the mid point is less than the element at the end of the subset, then elements are only becoming bigger from the mid point onwards
        if(nums[mid]<nums[end]){
            end=mid
        }else{
            // else, the elements from that mid point are only becoming smaller onwards
            start=mid+1
        }
    }

    return nums[start]
}
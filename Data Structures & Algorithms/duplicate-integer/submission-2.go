func hasDuplicate(nums []int) bool {
   // an integer array nums 
   // return true if any value appears more than once in the array
   // other wise false
   // we use a hash map to record all the values that we have seen
   // then we loop through the array and for every value in view
   // we check if we have seen it before, if we have not, we just record it and continue
	seen:=map[int]bool{}

   for _, num:=range nums {
	// check if seen before
	_,exists:=seen[num]

	if(exists){
		return true
	}

	// if does not exist, just record that we have seen it
	seen[num]=true
   }

   return false
}

func productExceptSelf(nums []int) []int {
	// initialize 3 arrays which are of the same length as the nums
	total:=len(nums)

	pref:=make([]int,total)
	suff:=make([]int,total)
	res:=make([]int,total)

	// initialize the pref at the first index to 1 since it doesnt have anything before it
	pref[0]=1

	// initialize the suff at the last index to 1 since it doesnt have anything after it
	suff[total-1]=1

	// 
	for i:=1;i<total;i++{
		pref[i]=nums[i-1]*pref[i-1]	
	}

	for i:=total-2;i>=0;i-- {
		suff[i]=nums[i+1]*suff[i+1]
	}

	// multiply the numbers 
	for i:=0;i<total;i++{
		res[i]=pref[i]*suff[i]
	}

	return res
}

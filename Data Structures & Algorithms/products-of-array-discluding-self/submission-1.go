func productExceptSelf(nums []int) []int {
	product:=1
	zeroCount:=0

	// loop through each element
	for _,num:= range nums {
		if (num==0){
			zeroCount=zeroCount+1
		}else{
			product=product*num
		}
	}

	// if there are 2 or more zeros, return an array of 0
	if (zeroCount>=2){
		res := make([]int, len(nums))
		return res
	}

	result:=[]int{}
	
	// if there is one zero in the array, it means only the element that is zero would be nonzero, all other elements would be zero
	// if there is no zero at all, just divide the product with that element
	for _,num:=range nums {
		if (zeroCount ==1){
			if (num==0){
			result=append(result,product)
		}else{
				result=append(result,0)
			}
		}else{
			result=append(result,product/num)
		}
	}

	return result
}

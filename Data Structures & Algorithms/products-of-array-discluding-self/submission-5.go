func productExceptSelf(nums []int) []int {
	// get the product of all elements in the array
	product:=1
	zeroCount:=0

	// for each element just get the product
	for _,num:=range nums {
		if (num==0){
			zeroCount=zeroCount+1
		}else{
			product=product*num
		}
	}

	// if there are more than 2 zeros then no product can be non zero
	if (zeroCount>=2){
		res:=make([]int,len(nums))
		return res
	}

	res:=[]int{}

	// if there is only 1 zero, then only the element that is zero, can have a non zero result
	for _,val:=range nums {
		if(zeroCount==1){
			if(val==0){
				res=append(res,product)
			}else{
				res=append(res,0)
			}
		}else{
		 res=append(res,product/val)
		}
	}

	return res
}

func maxProfit(prices []int) int {
	//profit
	profit:=0

	left:=0
	right:=1

	// keep looping through the array until the right is at the end
	for(right<len(prices)){
		// if its a profitable transaction
		if(prices[right]>prices[left]){
			newProfit:=prices[right]-prices[left]

			if(newProfit>profit){
				profit=newProfit
			}
		}else{
			// move the buying day to the current selling date
			left=right
		}
		right++
	}

	return profit
}

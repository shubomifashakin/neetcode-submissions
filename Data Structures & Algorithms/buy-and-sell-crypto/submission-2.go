func maxProfit(prices []int) int {
	// the ask is the get the maximum profit we can make 
	// to do that we initialze 2 pointers
	// one on the first day and one on the second day
	// on the first day, we check if the result of selling would be profitable, if it wouldnt we move the day to buy to the current date to sell
	// and we move the day to sell forward by one
	// if it would we just record it and keep looping

	buy:=0
	sell:=1
	maxProfit:=0
	
	// keep running this loop until we have reached the last day to sell
	for(sell<len(prices)){
		// check if the transaction is profitable
		if(prices[sell]>prices[buy]){
			profit:=prices[sell]-prices[buy]
			if(profit>maxProfit){
				maxProfit=profit
			}
		}else{
			buy=sell
		}
		sell++
	}

	return maxProfit
}

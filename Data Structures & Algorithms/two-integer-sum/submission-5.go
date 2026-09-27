
// using two pointers
func twoSum(nums []int, target int) []int{
	// initialize 2 pointers
	// one starts at point 0 and the other starts at point 1
	// for each point in significant, increment the other pointer until it finds an element that satisfies the result
	// if it does not, move to the next element and return insig to n+1

	sig:=0
	insig:=1

	// while the sig hasnt reached the end, continue
	for sig<len(nums)-1{
		// check if the values at sig and insig add up to meet target
		for insig<len(nums){
			if (nums[sig]+nums[insig]==target){
				return []int{sig,insig}
			}
			insig++
		}
		sig++
		insig=sig+1
	}

	return []int{}
	}
func searchMatrix(matrix [][]int, target int) bool {
	// given a matrix and a target
	// each row in a matrix is sorted in acending order
	// first integer of every row is greater than the last integer of previous row
	// return true if target exists or false otherwise

	start:=0
	end:=len(matrix)-1

	// keep running the loop so far they have not crossed
	for start<=end{
		// get the mid point of the matrix
		mid:=(start+end)/2
		midRow:=matrix[mid]

		// if the target is greater than or equal to the first element of this row and less than or equal to the last integer of this row, then it should be here
		if(target >=midRow[0] && target<=midRow[len(midRow)-1]){
			// check if the target is here
			// loop through each element in the array
			for _,num:=range midRow{
				if(num==target){
					return true
				}
			}

			return false
		}

		// if the target is less than the first integer of this row, then we shouldnt bother checking anything from this row forward (from start - midRow-1)
		if(target<midRow[0]){
			end=mid-1
			continue
		}

		// if the target is greater than the last element of this row, then we should check from midRow+1 till the end
		if(target>midRow[len(midRow)-1]){
			start=mid+1
		}
	}

	return false
}

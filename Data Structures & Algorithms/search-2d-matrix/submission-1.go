func searchMatrix(matrix [][]int, target int) bool {
	// get the start and end
	start:=0
	end:=len(matrix)-1

	for start<=end{
		// get the mid point
		midVal:=(start+end)/2
		mid:=matrix[midVal]
		firstElemInMid:=mid[0]
		lastElemInMid:=mid[len(mid)-1]

		// if the firstElemInMid is greater than the target, dont consider everything from mid till current end anymore
		if(firstElemInMid>target){
			end=midVal-1
		}

		// if the target is greater than firstElement but less than lastElement in mid, then the target should be in view
		if(target>=firstElemInMid && target<=lastElemInMid){
			// use binary search here also
			innerStart:=0
			innerEnd:=len(mid)-1
	
			for innerStart<=innerEnd{
				innerMid:=(innerStart+innerEnd)/2
				elemAtMid:=mid[innerMid]
				
				// if the target is less than the mid point
				// dont consider anything from that mid point forward
				if(target<elemAtMid){
					innerEnd=innerMid-1
					continue
				}

				if(target>elemAtMid){
					innerStart=innerMid+1
					continue
				}

				if(target==elemAtMid){
					return true
				}
			}

			return false
		}

		// if the target is greater than the lastElementIn mid then we should not consider everything from start to mid anymore
		if(target>lastElemInMid){
			start=midVal+1
		}
	}

	return false
}

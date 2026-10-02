type TimeMap struct {
	Values map[string][][]string
}

func Constructor() TimeMap {
	return TimeMap{
		Values:map[string][][]string{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	// check if the key exists
	val,exists:=this.Values[key]

	converted:=strconv.Itoa(timestamp)
	if(!exists){
		this.Values[key]=[][]string{{value,converted},}
	}else{
		arr:=[]string{value,converted}
		this.Values[key] = append(val, arr)
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	// get the val from the map
	vals,exists:=this.Values[key]

	if(!exists){
		return ""
	}
	
	// use binary search to scan for the element
	start:=0
	end:=len(vals)-1

	res:=""
	for start<=end {
		mid:=(start+end)/2
		midEl:=vals[mid]

		// if the val at the mid element is the timestamp, then just return it
		midTime,_:=strconv.Atoi(midEl[1])
		if(midTime==timestamp){
			return midEl[0]
		}

		// if the timestamp at the mid element is greater than requested
		// only consider from start to mid-1
		if(midTime>timestamp){
			end=mid-1
			continue
		}

		// if the timestamp at the mid element is less than requested
		// only consider from mid+1 to end
		if(midTime<timestamp){
			// but this particular mid element has been the best so far
			res=midEl[0]
			start=mid+1
		}
	}

	return res
}

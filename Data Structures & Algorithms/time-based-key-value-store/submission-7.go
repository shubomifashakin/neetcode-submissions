type TimeMap struct {
	Values map[string][][]string
}

func Constructor() TimeMap {
	return TimeMap{
		Values:map[string][][]string{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	// receives the key, value and the timestamp
	val,exists:=this.Values[key]

	// convert the timestamp to an integer
	converted:=strconv.Itoa(timestamp)
	if(exists){
		// create a new array with the values
		arr:=[]string{value,converted}
		val=append(val,arr)
		// update the current state fir that key
		this.Values[key]=val
	}else{
		arr:=[]string{value,converted}
		this.Values[key]=[][]string{arr}
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	// we need to get the key from the array and
	// if the one for the timestamp is not found
	// we need to get the one that comes immediately before it

	// check if the key exists
	vals,exists:=this.Values[key]

	if(!exists){
		return ""
	}

	
	// we know the timestamp is always increasing
	// so itechnically its sorted
	start:=0
	end:=len(vals)-1

	res:=""
	// keep running the loop so far they have not crossed
	for start<=end{
		mid:=(start+end)/2
		midEl:=vals[mid]

		midTime,_:=strconv.Atoi(midEl[1])

		// if the requested timestamp is the same as the timestamp for the mid
		// return it
		if(midTime==timestamp){
			return midEl[0]
		}

		// if the requested timestamp is greater than the mid point timestamp
		// onnly consider values from mid+1 to end
		if(timestamp>midTime){
				res=midEl[0]
			start=mid+1
		}

		// if the requested timestamp is less than the mid point timestamp
		// only consider values from start to mid-1
		if(timestamp < midTime){
			end=mid-1
		}
	}

	return res
	
}

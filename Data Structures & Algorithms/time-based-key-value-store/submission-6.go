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

	res:=""

	// loop through the values in that key set
	// if the we have reached the one for that timestamp, then return it
	// if we have reached a value greater than the given timestamp, return previous one

	for _,pairs:=range vals{
		// get the timestamp
		timeStr:=pairs[1]

		// convert to number
		converted,_:= strconv.Atoi(timeStr)

		if(converted>timestamp){
			break
		}

		res=pairs[0]
		if(converted==timestamp){
			return pairs[0]
		}

	}

	return res
}

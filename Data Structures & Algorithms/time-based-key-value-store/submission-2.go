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
	
	
	lastVal:=""
	for _, el := range vals {
	
		// keep checking the slice until you find
		strTime := el[1]

		// to int
		converted, _ := strconv.Atoi(strTime)

		if converted > timestamp {
			break
		}

		// if the value at that timestamp was found, just return it
		if converted == timestamp {
			return el[0]
		}
		lastVal = el[0]
	}

	return lastVal
}

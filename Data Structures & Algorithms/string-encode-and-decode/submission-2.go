type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    encoded:=""
    // for each string in the slice
    // we get the length of the string
    // once we have the length, we build a new string
    for _,str:=range strs{
        // get the length of the string
        strLen:=len(str)

        // convert the len to string
        converted:=strconv.Itoa(strLen)
        encoded=encoded+converted+"#"+str
    }

    return encoded

}

func (s *Solution) Decode(encoded string) []string {
    // we have a string
    // we know that the string is of a certain variant
    // the first characters are the number of words to take
    // separated by a delimiter

    // lets keep track of our current position
    currentPosition:=0
    words:=[]string{}

    // while our current position is less than the string length
    // we keep the loop running
    for (currentPosition<len(encoded)-1){
        // we need to get the next delimiter from our position
        idxOfDelim:=currentPosition

        // so far the current character in view is not our delim, keep looking
        for(string(encoded[idxOfDelim])!="#"){
            idxOfDelim++
        }

        // now we have the index of the delimiter

        // we need to get the amount of characters to take from the delimiter
        take:=encoded[currentPosition:idxOfDelim]

        // convert that to a number
        takeNumber,_:=strconv.Atoi(take)

        // now we have our index of delimiter and number to take from there
        start:=idxOfDelim+1
        end:=start+takeNumber

        // take from tat from the string
        word:=encoded[start:end]
        words=append(words,word)

        currentPosition=end
    }

    return words
}

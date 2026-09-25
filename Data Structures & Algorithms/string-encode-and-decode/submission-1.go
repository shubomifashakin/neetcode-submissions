type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    // for each string, we get the length and we append a delim to it
    encoded:=""

   for _,word:=range strs {
        // length of the string
        wordLength:= len(word)
		lenString:=strconv.Itoa(wordLength)
        newStr:=lenString+"#"+word

        encoded=encoded+newStr
    }

    return encoded
}

func (s *Solution) Decode(encoded string) []string {
    // track our position
    position:=0

    words:=[]string{}

    // so far we have not reached the end of the string, keep running
    for(position!=len(encoded)){
        // get the next occurence of the delimiter from our current position
        idxOfDelim:=position
        for(string(encoded[idxOfDelim])!= "#"){
            idxOfDelim++
        }

        // now that we have the index of delimiter and the position we stopped
        // we know the amount to take is between those
        numberOfCharsToTake:=encoded[position:idxOfDelim]
        converted,_:=strconv.Atoi(numberOfCharsToTake)

        // take that amount from the index of the delim
        start:=idxOfDelim+1
        end:=start+converted

        // take the word
        word:=encoded[start:end]
        // push it to the list
        words=append(words,word)

        position=end
    }

    return words
}

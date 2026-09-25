class Solution {
    /**
     * @param {string[]} strs
     * @returns {string}
     */
    encode(strs: string[]): string {
        //for each string, count it and append the delim
        let result=""

        for(const elem of strs){
            // find the count of the elements
           const strLength= elem.length

           // form the new string
           const newString= `${strLength}#${elem}`

           result=result+newString
        }
        
        return result
    }

    /**
     * @param {string} str
     * @returns {string[]}
     */
    decode(str: string): string[] {
        const words:string[]=[]

        let currentPosition=0

        // so far we havent reached the end of the str continue
        while(currentPosition<str.length){
            let idxOfDelim=currentPosition
            // get the index of the next delimite from our current position
            while(str[idxOfDelim]!== "#"){
                idxOfDelim++
            }

            // now that we have the index of the next delimiter and our current position, we know the amount to take is between those
            const amountOfCharsToTake= Number(str.substring(currentPosition,idxOfDelim))

            const start=idxOfDelim+1
            const end=start+amountOfCharsToTake

            const word= str.substring(start,end)
            words.push(word)

            currentPosition=end
        }

        return words
    }
}

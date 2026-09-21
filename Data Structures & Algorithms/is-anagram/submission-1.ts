class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s: string, t: string): boolean {
        // if they are not the same length then they are not anagrams, false
        if(s.length !== t.length)return false

        // split the characters of stringA, so it becomes an array, then sort it
        const splitA= s.split('').sort()

        // split the characters of stringB, so it becomes an array then sort it
        const splitB= t.split("").sort()
      
        let checked=0
        // so far we havent checked everything
      while(checked!=splitB.length){
        if(splitA[checked]!==splitB[checked]){
            return false
        }
        checked=checked+1
      }

        return true
    }
}

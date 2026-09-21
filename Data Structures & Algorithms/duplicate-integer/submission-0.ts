// using brute force, we loop through the arry elements
// comparing each value with other values in the array BESIDES ITSELF
// if the 2 values match and their index are not the same, we have found a match
// if we have compare everything and no match then theres no match
// class Solution {
    
//     /**
//      * @param {number[]} nums
//      * @return {boolean}
//      */
//     hasDuplicate(nums: number[]): boolean {
//     //    for(let x=0;x<=nums.length;x++){
//     //    // get the value at this position
//     //    const val=nums[x]
//     //     for(let y=0;y<=nums.length;y++){
//     //         const val2=nums[y]
//     //         // so far the index is not the same and the current values 
//     //         // are the same, we have a match
//     //         if(val === val2 && x!==y){
//     //             return true
//     //         }
//     //     }
//     //    }

//     //    return false
//     }
// }


// // Using two pointers solution
// class Solution {
    
//     /**
//      * @param {number[]} nums
//      * @return {boolean}
//      */
//     hasDuplicate(nums: number[]): boolean {
//     // sort the array first
//     // after sorting the array, use 2 pointers and compare side by side
//     const sortedArray=nums.sort((a,b)=>{return a-b})
//     // initialize 2 pointers (one that would be ahead and 1 that starts at the beginning)
//     let start=0
//     let end=1

//     // so far the start is not at the second to the last position
//     // and the end is not at the last position, run this
//     while(start!==nums.length-1 && end!==nums.length){
//         if(nums[start]==nums[end]){
//             return true
//         }
//         start=start+1
//         end=end+1
//     }

//     return false
//     }
// }



// HASH SET
class Solution {
    
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums: number[]): boolean {
        // create a set that would store seen values
        const seenValues= new Set()

        // loop through the array
        // for each iteration, check if we have the value on record, if we do, then we have a match
        // if we dont, record it in the array
        for(let x=0; x<=nums.length;x++){
            if (seenValues.has(nums[x])){
                return true
            }

            // if it doesnt contain it, record it
            seenValues.add(nums[x])
        }

        return false
    }
}

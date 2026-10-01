func findMin(nums []int) int {
    res := nums[0]
    l, r := 0, len(nums)-1

    for l <= r {
		// if we are already in a sorted range
        if nums[l] < nums[r] {
            if nums[l] < res {
                res = nums[l]
            }
            break
        }

        mid := l + (r-l)/2
		// if element at mid point is smaller than the current result, update result
        if nums[mid] < res {
            res = nums[mid]
        }

		// if element at mid point is less than val at start, then we need to search the elements at the right
        if nums[mid] >= nums[l] {
            l = mid + 1
        } else {
			// else search elements at left
            r = mid - 1
        }
    }
    return res
}
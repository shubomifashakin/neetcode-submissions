func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    combined := append(nums1, nums2...)
    sort.Ints(combined)

    total := len(combined)

	// if the array length is even
    if total %2 == 0 {
        return float64(combined[total/2-1] + combined[total/2]) / 2.0
    }
    return float64(combined[total/2])
}
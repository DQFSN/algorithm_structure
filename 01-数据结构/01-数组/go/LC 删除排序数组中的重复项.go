package main

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
    r, l := 0, 0
    for r < len(nums) {
        if nums[r] != nums[l] {
            l++
            nums[l], nums[r] = nums[r], nums[l]
        }
        r++
    }
    return l+1;
}
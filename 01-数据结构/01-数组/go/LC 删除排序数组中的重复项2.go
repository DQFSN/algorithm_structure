package main

func removeDuplicates2(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
    r, l,count := 0, 0, 0
    for r < len(nums) {
        if nums[r] != nums[l] {
            l++
            nums[l], nums[r] = nums[r], nums[l]
            count = 1;
        }else {
            count++
            if (count == 2) {
                l++
                nums[l], nums[r] = nums[r], nums[l]
            }

        }
        r++
    }
    return l+1;
}
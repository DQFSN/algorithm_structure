package main

func removeElement(nums []int, val int) int {
    l, r := 0,0
    for r < len(nums) {
        if nums[r] != val {
            nums[l], nums[r] = nums[r], nums[l]
            l++
        }
        r++
    }
    return l;
}
package main

func findKthLargest(nums []int, k int) int {
    return findKth(nums, 0, len(nums)-1, k)
}

func findKth (nums []int, l int, r int, k int) int {
    n := partion(nums,l,r)

    if n == len(nums) - k {
        return nums[n]
    }else if n < len(nums) - k {
        return findKth(nums, n+1, r, k)
    }else {
        return findKth(nums, l, n-1, k)
    }
}

func partion(nums []int, l int, r int) int {
    povt := nums[l]
    left, right := l, r
    for l < r {
        for l < right && nums[l] <= povt {
            l++
        }
        for left < r && nums[r] >= povt {
            r--
        }

        if l < r {
            nums[l], nums[r] = nums[r], nums[l]
            l++
            r--
        }
    }

    nums[left], nums[r] = nums[r], nums[left]
    // fmt.Println(nums,r)
    return r;
}
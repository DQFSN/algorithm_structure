package main

func moveZeroes(nums []int)  {
    for i, v := range nums {
        if v != 0 {
            continue
        }

        for j, v2 := range nums[i:] {
            if v2 == 0 {
                continue
            }
            nums[i], nums[j+i] = nums[j+i], nums[i]
            
            break
        }
    }
}
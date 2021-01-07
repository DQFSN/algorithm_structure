package main

func minSubArrayLen(s int, nums []int) int {
    l, r := 0, 0
    sum := 0
    min := len(nums) + 1
    for l < len(nums) {
        
        sum = getsum(nums[l:r])
        // fmt.Println(nums[l:r],sum,sum == s, min)

        if sum >= s {
            if r == l {
                return 1;
            }
            if r - l < min {
                min = r - l
            }
            l++
        }else {
            if r < len(nums) {
                r++                
            }else {
                break;
            }
        }

    }

    if min > len(nums) {
        return 0
    }
    return min
}

func getsum(nums []int) (sum int) {
    for _, v := range nums {
        sum += v
    }

    return
}




func minSubArrayLen2(s int, nums []int) int {
    l, r := 0, 0
    sum := 0
    min := len(nums) + 1

    for r < len(nums) + 1 {
        // fmt.Println(nums[l:r],sum)

        if sum < s {
            if r < len(nums) {
                sum += nums[r]
            }
            r++
        }else {
            if r - l < min {
                min = r - l
            }
            sum -= nums[l]
            l++
        }
    }

    if min > len(nums) {
        return 0
    }else {
        return min
    }

}

package main

func maxArea(height []int) int {
    l, r := 0, len(height)-1
    max := 0
    for l < r {
        h, w := 0, r - l
        if height[l] <= height[r] {
            h = height[l]
            l++
        }else {
            h = height[r]
            r--
        }
        if w * h > max {
            max = w * h
        }
    }

    return max    
}
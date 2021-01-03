package main


func sortColors(nums []int)  {

    dict := make(map[int]int)
    for _,v := range nums {
        dict[v]++
    }

    for i,_ := range nums {
        if i < dict[0] {
            nums[i] = 0
        }else if (i - dict[0]) < dict[1] {
            nums[i] = 1
        }else {
            nums[i] = 2
        }
    }
}



func sortColors1(nums []int)  {
    if len(nums) < 2 {
        return
    }else if len(nums) == 2 {
        if nums[0] > nums[1] {
            nums[0], nums[1] = nums[1], nums[0]
        }
        return
    }

    l, mid, r := 0, len(nums)/2, len(nums)
    sortColors(nums[l:mid])
    sortColors(nums[mid:r])
    // fmt.Printf("pre:%v\n",nums)
    merge(nums)
    // fmt.Printf("post:%v\n",nums)


}

func merge(nums []int) {
    A := nums[0:len(nums)/2]
    B := nums[len(nums)/2:]
    C := make([]int, len(nums), len(nums))
    a, b, c := 0,0,0
    for c < len(C) {
        if a < len(A) && b < len(B) {
            if A[a] <= B[b] {
                C[c] = A[a]
                a++
            }else {
                C[c] = B[b]
                b++
            }
            c++
        } else  {
            for a < len(A) {
                C[c] = A[a]
                a++
                c++
            }
            for b < len(B) {
                C[c] = B[b]
                b++
                c++
            }
        }
    }

    for i,v := range C {
        nums[i] = v
    }
}
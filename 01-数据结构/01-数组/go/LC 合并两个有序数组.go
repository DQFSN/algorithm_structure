package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
    a, b := m-1, n-1
    c := m+n-1

    for a >= 0 && b >= 0 {
        if nums1[a] <= nums2[b] {
            nums1[c] = nums2[b]
            b--
        }else {
            nums1[c] = nums1[a]
            a--
        }
        c--
    }

    for a >= 0 {
        nums1[c] = nums1[a]
        a--
        c--
    }

    for b >= 0 {
        nums1[c] = nums2[b]
        b--
        c--
    }

}
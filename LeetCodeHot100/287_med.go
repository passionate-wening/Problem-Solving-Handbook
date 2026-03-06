package LeetCodeHot100

func findDuplicate(nums []int) int {
    slow,fast:=nums[0],nums[nums[0]]
    for nums[slow] != nums[fast] {
        fmt.Println(slow,fast)
        slow = nums[slow]
        fast = nums[nums[fast]]
    }
    slow=0
    for nums[slow] != nums[fast] {
        slow=nums[slow]
        fast=nums[fast]
    }
    return nums[slow]
}

/*
【题解】
看了官解，这个环思想真的很棒！
*/

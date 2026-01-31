func searchInsert(nums []int, target int) int {
    return checkInsert(nums,target,0,len(nums)-1)
}

func checkInsert(nums []int, target,minI,maxI int) int {
    //二分法
    //中位数成立，输出
    med:=(maxI+minI)/2
    if nums[med] == target {
        return med
    } else if nums[med] < target {
    //不成立，大于区间为空，i+1
        if med == maxI {
            return maxI+1
        } else {
            return checkInsert(nums,target,med+1,maxI)
        }
    } else {
    //小于区间为空，i
        if med == minI {
            return minI
        } else {
            return checkInsert(nums,target,minI,med-1)
        }
    }
    //其他，递归
}

//官方二分法是用迭代写的，日后要练习

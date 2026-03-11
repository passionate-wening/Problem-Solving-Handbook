package LeetCodeHot100

func maxArea(height []int) int {
	maxHeight := 0
	maxVolume := 0
	for i := 0; i < len(height)-1; i++ {
		if maxHeight > height[i] {
			continue
		}
		maxHeight = height[i]
		for j := i + 1; j < len(height); j++ {
			h := maxHeight
			if h > height[j] {
				h = height[j]
			}
			volume := h * (j - i)
			if volume > maxVolume {
				maxVolume = volume
			}
		}
	}
	return maxVolume
}

/*
【题解】
双指针右移，长是两指针之差，高是较矮的值，那么在移动过程中，长是等量缩减的，那么当左指针的高度低于前面计算过的高度，就没有必要再算容积了。
官方题解更好，我既然想到了移动左指针舍弃低位数，就应该想到让右指针也有取舍。
官方给出的很重要的题目已知知识：容积 = 两个指针指向的数字中较小值x ∗ 指针之间的距离y ；
所以双指针是从两头到中间取值，移动过程y是恒定减少的，那么只有移动两指针中值较小的那一个，才有可能增加x的值。
*/

func maxArea1(height []int) int {
	i, j, maxVolume := 0, len(height)-1, 0
	for i < j {
		x, y := j-i, height[j]
		if height[i] < height[j] {
			y = height[i]
			i++
		} else {
			j--
		}
		volume := x * y
		if maxVolume < volume {
			maxVolume = volume
		}
	}
	return maxVolume
}

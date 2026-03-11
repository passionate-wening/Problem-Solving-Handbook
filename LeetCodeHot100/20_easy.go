package LeetCodeHot100

type stack struct {
	nums []byte
}

func isValid(s string) bool {
	m := make(map[byte]byte)
	m = map[byte]byte{byte(')'): byte('('), byte(']'): byte('['), byte('}'): byte('{')}
	st := &stack{nums: []byte{}}
	for _, val := range []byte(s) {
		v, ok := m[val]
		if ok {
			if st.pop() != v {
				return false
			}
		} else {
			st.push(val)
		}
	}
	return len(st.nums) == 0
}

func (s *stack) pop() byte {
	if len(s.nums) < 1 {
		return byte('0')
	}
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

func (s *stack) push(v byte) {
	s.nums = append(s.nums, v)
}

/*
【题解】栈简单题放松一下（要考虑清楚每一个边界情况哦）
*/

package base

func Add(first, second int) int {
	return first + second
}

func Max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func Count(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func Palindrome(nums []int) bool {
	size := len(nums)
	mid := len(nums) / 2
	for i := 0; i < mid; i++ {
		if nums[i] != nums[size-i-1] {
			return false
		}
	}
	return true
}

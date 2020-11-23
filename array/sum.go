package array

func Add(nums ...int) (result int) {
	for _, num := range nums {
		result += num
	}
	return
}

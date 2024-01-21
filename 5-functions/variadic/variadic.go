package variadic

/*
	Variadic function
		function that can take an arbitrary number of final arguments using "..." syntax
*/
func Sum(nums ...int) (num int) {
	// nums is just a "slice" (a reference of an array)
	for i := 0; i < len(nums); i++ {
		num = nums[i]
	}
	return num
}

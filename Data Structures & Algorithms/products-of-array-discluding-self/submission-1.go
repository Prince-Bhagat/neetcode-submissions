
func productExceptSelf(nums []int) []int {
	size := len(nums)
	

	left := make([]int, size)
	left[0] = nums[0]

	right:= make([]int, size)
	right[size-1] = nums[size-1]

	result := make([]int, size)

	counter := 1

	for counter < size{
		left[counter] =  nums[counter] * left[counter-1]
		counter++
	}
	

	counter = size-2

	for counter >= 0 {
		
		right[counter] = nums[counter] * right[counter+1]
		counter--
	}
	


	for i := 0; i < size; i++ {
		leftValue := func (i int)int{ if i<0 {return 1}else { return left[i]}} (i-1)
		rightValue := func (i int)int{ if i >= size {return 1}else { return right[i]}}(i+1)
		result[i] = leftValue * rightValue
	}
	return result
}

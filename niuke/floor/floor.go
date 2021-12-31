package floor

func jumpFloor(number int) int {
	// write code here
	if number <= 2 {
		return number
	}
	arr := make([]int, 0, number+1)
	arr = append(arr, 0, 1, 2)
	for i := 3; i <= number; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}
	return arr[number]
}

//  0 = 0
//  1 = 1
//  2 = 2
//  3 = 3  2 + 1
//  4 = 5
//  5 = 8
//  6 = 13
//  7 = 21

package stack2quque

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 {
		if len(stack1) == 0 {
			//  保证一定成功
			// return -1
		} else {
			for i := 0; i < len(stack1); i++ {
				stack2 = append(stack2, stack1[i])
			}
			stack1 = make([]int, 0)
		}
		result := stack2[0]
		stack2 = stack2[1:]
		return result

	}
	return -1
}

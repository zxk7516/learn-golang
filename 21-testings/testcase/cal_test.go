package testcase

import "testing"

// 1, 测试文件以 _test 为后缀
// 2, 测试方法以 Test 为开头
// 3, 测试方法参数为 t *testing.T
// 4, go test --run ^TestAddUpper  -- 运行单个函数 (TestAddUpper())
// 5, go test -v  -- 运行正确错误都有日志
func TestAddUpper(t *testing.T) {
	test_cases := [][]int{
		{10, 55},
		{1, 1},
	}
	for _, case_arr := range test_cases {
		res := AddUpper(case_arr[0])
		if res != case_arr[1] {
			t.Fatalf("错误, AddUpper(%v) should equal %v", case_arr[0], case_arr[1])
		}

	}

}

package main

/*func removeOuterParentheses(S string) string {
	result := make([]byte, 0, len(S))
	count := 0
	for i := range S {
		if S[i] == '(' {
			if count > 0 {
				result = append(result,S[i])
			}
			count++
		}else {
			count--
			if count > 0 {
				result = append(result,S[i])
			}
		}
	}
	return string(result)
}*/
/*func countNegatives(grid [][]int) int {
	M := len(grid)
	count := 0
	for i := range grid {
		N := len(grid[M-1])
		for j := range grid[i] {
			if grid[i][N-1] > 0 {
				break
			}else {
				count++
				N--
			}
			j++
		}
	}
	return count
}*/
/*func maxSubArray(nums []int) int {
	temp := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+temp < nums[i] {
			temp = nums[i]
		} else {
			temp += nums[i]
		}
	}
	if temp > max {
		max = temp
	}
	return max
}*/
/*package main

import (
"fmt"
"reflect"
)
type Person struct {
	name string
}

func (p *Person)Eat()  string{
	return p.name+"吃饭"
}
func main()  {
	p := &Person{
		"张三",
	}
	name := reflect.ValueOf(p).MethodByName("Eat").Call(make([]reflect.Value,0))
	fmt.Println(name)
}*/
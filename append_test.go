package main

import "testing"

func TestAppend(t *testing.T) {
	nums := []int{2, 3, 4}

	dnums := nums[0:2]
	dnums = append(dnums, 2)
	t.Log(dnums)
}

package main

import (
	"testing"
)

func TestSlice(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 3, 2, 2, 2, 2, 2}
	i := 5
	a = append(a[:i], append(a[i+1:], a[i])...)
	t.Log(a)

	var PubKey []int
	for _, value := range a {
		PubKey = append(PubKey,value)
	}
	t.Log(PubKey)
}

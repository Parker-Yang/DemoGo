package test

import (
	"fmt"
	"testing"
)

type Employee struct {
	id   string
	Name string
	age  int
}

func TestCreateEmployee(t *testing.T) {
	e := Employee{"0", "zhangsan", 2}
	e1 := Employee{id: "1", Name: "lisi", age: 2}
	e2 := new(Employee)
	e2.id = "2"
	e2.Name = "wangwu"
	e2.age = 20
	t.Log(e)
	t.Log(e1)
	t.Log(e2)
}

// 实现方法，使用指针避免内存拷贝
func (e *Employee) String() string {
	return fmt.Sprintf("id:%s\nName:%s\nage:%d", e.id, e.Name, e.age)
}
func TestToString(t *testing.T) {
	e := Employee{"1", "zhaoniu", 20}
	t.Log(e.String())
}

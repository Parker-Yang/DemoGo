package test

import (
	"fmt"
	"testing"
)

type User struct {
	Name string
}

func NewUser(s string) User {
	return User{Name: s}
}

type m1 interface {
	setname(s string)
}
type m2 interface {
	getname(s string)
}

func (u User) setname(s string) {
	fmt.Println("拷贝")
	u.Name = s
}
func (u *User) getname(s string) {
	fmt.Println("地址")
	u.Name = s
}
func TestGo(t *testing.T) {
	user := NewUser("zs")
	t.Log(user)
	// user.setname("lisi")
	User.setname(user, "lisi")
	t.Log(user)
	user.getname("lisi")
	t.Log(user)

}

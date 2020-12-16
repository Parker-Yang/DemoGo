package test

import (
	"encoding/json"
	"testing"
)

type user struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	data
}
type data struct {
	Time int `json:"time"`
}
func TestJson(t *testing.T) {
	u := &user{
		Name:   "zhansan",
		Age:    12,
		Gender: "male",
	}
	u.Time=1
	t.Log(u)
	marshal, err := json.Marshal(u)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(marshal))
	b := &user{}
	err = json.Unmarshal(marshal, b)
	if err != nil {
		t.Log(err)
	}
	t.Log(b)
}

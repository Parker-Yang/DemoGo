package test

import (
	"strings"
	"testing"

	"github.com/imdario/mergo"
)

// 用户模型
type TUser struct {
	// 用户ID
	UserID string `json:"userID" validate:"@string[1,]"`
	// 用户姓名username
	UserName string `json:"userName" validate:"@string[1,]"`
}

// Create & Update data
type UserData struct {
	// 用户信息
	User TUser `json:"user"`
	// 发起请求的用
	Sign string `json:"sign"`
	// 发起请求的用户id
	ReqUserID string `json:"reqUserID"`
	// 操作时间戳
	Timestamp int64 `json:"timestamp"`
}
func TestMergo(t *testing.T) {
	u1 := &TUser{UserName: "sd"}
	u2 := &TUser{UserID: "123"}
	// u3 := UserData{User: u2}
	contains := strings.Contains("123", "1")
	t.Log(contains)
	_ = mergo.Merge(u1,u2, mergo.WithOverride)
	t.Log(u1)
	t.Log(u2)

}

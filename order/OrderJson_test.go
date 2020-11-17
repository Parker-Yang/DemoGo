package order

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	data := UserData{
		UserID:         "123",
		UserName:       "张三",
		UserPub:        "abdsahduisa",
		Role:           1,
		OrgID:          "",
		DeptID:         "",
		Position:       "管理员",
		Description:    "认真工作",
		Tel:            "1234567890",
		Email:          "2341@qq.com",
		Birthday:       12312312,
		Gender:         1,
		ReceiverIDList: []string{"easdsad", "dasdas"},
		Sign:           "123123",
		Dept: DeptData{
			DeptName:        "123",
			DeptDescription: "123",
			OrgID:           "123",
		},
	}
	result := Order(data)
	fmt.Printf("%v",result)
}

type Gender int

func (g Gender) String() string {
	return "hello"
}

// 请求用户数据
type UserData struct {
	// 用户ID/UkeyId
	UserID string `json:"userID,omitempty" validate:"@string[0,]"`
	// 用户姓名username
	UserName string `json:"userName" validate:"@string[1,]"`
	// 用户公钥
	UserPub string `json:"userPub,omitempty" validate:"@string[0,]"`
	// 所属组织organizationId
	OrgID string `json:"orgID,omitempty" validate:"@string[0,]"`
	// 所属部门departmentId
	DeptID string `json:"deptID,omitempty" validate:"@string[0,]"`
	// 用户角色
	Role int `json:"role,omitempty"`
	// 用户职位
	Position string `json:"position,omitempty" validate:"@string[0,]"`
	// 用户描述
	Description string `json:"description,omitempty" validate:"@string[0,]"`
	// 用户电话telephone
	Tel string `json:"tel,omitempty" validate:"@string[0,]"`
	// 用户邮箱email
	Email string `json:"email" validate:"@string[1,]"`
	// 用户出生年月birthday
	Birthday int64 `json:"birthday,omitempty" validate:"@int64[0,]"`
	// 用户性别gender
	Gender         Gender   `json:"gender,omitempty"`
	ReceiverIDList []string `json:"receiverIDList" validate:"@slice<@string[1,]>"`
	Sign           string   `json:"sign" validate:"@string[1,]"`
	Dept           DeptData `json:"dept"`
}
type DeptData struct {
	// 部门名称
	DeptName string `json:"deptName" validate:"@string[1,]"`
	// 组织ID
	OrgID string `json:"orgID" validate:"@string[1,]"`
	// 部门描述
	DeptDescription string `json:"deptDescription,omitempty" validate:"@string[0,]"`
}

// package main
//
// import (
//	"fmt"
//	"math"
// )
//
// func foo(a int)  {
//	defer fmt.Println("foot退出")
//	defer func() {
//		if r := recover(); r!=nil{
//			fmt.Printf("捕获到的错误：%s\n", r)
//		}
//	}()
//	if a < 0 {
//
//		panic("必须输入大于0的数")
//	}
//
//	fmt.Println("该数的方根为：", math.Sqrt(float64(a)))
// }
// func main() {
//	a,b := 10,-10
//	foo(a)
//	foo(b)
// }
// package main
//
// import "fmt"
//
// func main0501() {
//
//	//枚举定义和使用
//	const (
//		//在定义枚举时默认值为0 每换一行值增加1
//		//在定义枚举是可以不动写  =iota
//		a = iota //0
//		b = iota //1
//		c = iota //2
//
//		d
//		e
//	)
//	//a=100//err
//	fmt.Println(a, b, c, d, e)
// }
// func main() {
//	const (
//		//iota枚举里面数据可以赋值  不建议赋值
//		//枚举里面数据都是整型数据
//		a = iota
//		b = 10
//		c = iota
//		d
//		e
//	)
//	fmt.Println(a, b, c, d, e)
// }
// func main11() {
//	const (
//		//枚举值里面数据如果不换行 一行内的值相同
//		//枚举定义时如果没有赋值 会重复上一行格式
//		a    = iota
//		b, c = iota, iota
//		d, e
//	)
//	fmt.Println(a, b, c, d, e)
// }
// package main
//
// import "time"
//
// type stu struct {
//	id    int
//	name  string
//	sex   string
//	age   int
//	score int
// }

// //为结构体 数据类型绑定方法
// //结构体名可以直接作为方法的接收者
// func (s *stu) PrintInfo() {
//
//	s.name="大法师"
//	fmt.Println("大家好我叫", s.name, "我今年", s.age)
//
//
// }
//
// func main() {
//	s1 := stu{101, "张强伟", "男", 20, 0}
//
//
//	//对象.方法
//	s1.PrintInfo()
//
//	fmt.Println(s1)
//
// }

/*func main()  {
	ch := make(chan int)		// 用来进行数据通信
	quit := make(chan bool)		// 用来判断是否结束go程的channel

	go func() {
		for i:=0; i<5; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 300)
		}
		close(ch)
		fmt.Println("子go程打印完毕")
		quit <- true			// 通知主go程可以终止
	}()
	for {
		select {
		case num := <-ch :
			fmt.Println("num = ", num)
		case <-quit :
			//time.Sleep(time.Second)
			fmt.Println("主go程 收到子消息， 结束运行")
			//break		// 只能跳出当前 的 select
			return
			//runtime.Goexit()	// 退出主 go 程 ———— 不会清理进程地址空间。
			//os.Exit(0)
		}
		fmt.Println("==========================")
	}
}
func main() {
	a := 1
	b := 1
	c := []int{1,2,3,4,5}
	for _, i2 := range c {
		a := 2
		b := 2
		i2 += a+b
		fmt.Println(a,b)
		fmt.Println(i2)
	}
	fmt.Println(a,b)
	fmt.Println(c)
}
*/
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
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
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
func Order(data interface{}) string {
	text, _ := json.Marshal(data)
	temp := make(map[string]interface{})
	_ = json.Unmarshal(text, &temp)
	Filter(temp)
	result, _ := json.Marshal(temp)
	return string(result)
}

func Filter(data map[string]interface{}) {
	if _, ok := data["sign"]; ok {
		delete(data, "sign")
	}
	for _, value := range data {
		if elem, ok := value.(map[string]interface{}); ok {
			Filter(elem)
		}
	}
}
func serilizationSign(data interface{}) string {
	result := "{"
	typeOf := reflect.TypeOf(data)
	slice := make([]string, 0, typeOf.NumField())
	for i := 0; i < typeOf.NumField(); i++ {
		name := reflect.TypeOf(data).Field(i).Tag.Get("json")
		name = strings.Split(name, ",")[0]
		if name == "sign" {
			continue
		}
		t := reflect.TypeOf(data).Field(i).Type.Kind().String()
		value := fmt.Sprintf("%v", reflect.ValueOf(data).Field(i))
		if value == "" {
			continue
		}
		if reflect.ValueOf(data).Field(i).MethodByName("String").IsValid() {
			value = "\""
			value += fmt.Sprintf("%v", reflect.ValueOf(data).Field(i))
			value += "\""
		}
		if t == "string" {
			value = "\""
			value += fmt.Sprintf("%v", reflect.ValueOf(data).Field(i))
			value += "\""
		}
		if t == "struct" {
			value = serilizationSign(reflect.ValueOf(data).Field(i).Interface())
		}
		name = "\"" + name
		name += "\":" + value
		slice = append(slice, name)
	}
	sort.Strings(slice)
	result += strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", ",", -1) + "}"
	return result
}

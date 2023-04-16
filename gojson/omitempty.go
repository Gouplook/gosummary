package gojson

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserId   string
	UserName string
	age      int
	sex      string
}

type User2 struct {
	UserId   string `json:"id"`
	UserName string `json:"name"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
}

type User3 struct {
	UserId   string `json:"id"`
	UserName string `json:"name,omitempty"`
	Age      int    `json:"age,omitempty"`
	Sex      string `json:"sex,omitempty"`
}

type User4 struct {
	UserId   string `json:"id"`
	UserName string `json:"name,omitempty"`
	Age      int    `json:"-"`
	Sex      string `json:"sex,omitempty"`
}

func JO() {
	u := User{
		UserId:   "1",
		UserName: "张三",
		age:      20,
		sex:      "男",
	}
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err.Error())
	}

	u2 := User2{
		"1",
		"张三",
		20,
		"男",
	}
	data2, _ := json.Marshal(u2)

	u3 := User3{
		UserId:   "1",
		UserName: "om",
	}
	data3, _ := json.Marshal(u3)

	u4 := User2{
		UserId: "1",
	}
	data4, _ := json.Marshal(u4)

	u5 := User4{
		UserId:   "1",
		UserName: "张三",
		Sex:      "男",
	}
	data5, _ := json.Marshal(u5)
	fmt.Printf("%s ：只打印大写，小写自动忽略，只允许内部使用，json没有标记，因此默认使用变量名\n", string(data))
	fmt.Printf("%s :全部大写，并且使用json标记，因此序列化后自动使用json标记名称\n", string(data2))
	fmt.Printf("%s :全部大写，并且使用json标记，增加omitempty标记，带有该标记的不赋值的情况\n", string(data3))
	fmt.Printf("%s :全部大写，并且使用json标记，没有omitempty标记，上次字段依旧不赋值的情况\n", string(data4))
	fmt.Printf("%s :全部大写，并且使用json的-标记，没有omitempty标记，上次字段依旧不赋值的情况\n", string(data5))
}

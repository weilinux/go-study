package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string   `json:"name,omitempty"` //omitempty  为空则忽略此字段
	Age   int      `json:"age"`
	Hobby []string `json:"hobby"`
}

func main() {
	//TODO json序列化

	//结构体
	user1 := User{
		Name:  "小明",
		Age:   18,
		Hobby: []string{"篮球", "足球", "羽毛球"},
	}
	fmt.Printf("用户信息1：%v\n", user1)

	marshal, err := json.Marshal(&user1)
	if err != nil {
		return
	}
	fmt.Printf("序列化后结果：%v\n", string(marshal))

	//map
	user2 := make(map[string]interface{})
	user2["name"] = "小明"
	user2["age"] = 18
	user2["hobby"] = []string{"篮球", "足球", "羽毛球"}
	fmt.Printf("用户信息2：%v\n", user2)

	marshal2, err := json.Marshal(user2)
	if err != nil {
		return
	}
	fmt.Printf("序列化后结果：%v\n", string(marshal2))

	//slice
	user3 := make([]map[string]interface{}, 2)
	user3[0] = make(map[string]interface{})
	user3[0]["name"] = "小明"
	user3[0]["age"] = 18
	user3[0]["hobby"] = []string{"篮球", "足球", "羽毛球"}
	user3[1] = make(map[string]interface{})
	user3[1]["name"] = "小红"
	user3[1]["age"] = 17
	user3[1]["hobby"] = []string{"排球", "兵乓球", "网球"}
	fmt.Printf("用户信息3：%v\n", user3)

	marshal3, err := json.Marshal(user3)
	if err != nil {
		return
	}
	fmt.Printf("序列化后结果：%v\n", string(marshal3))

	//TODO 反序列化
	var user4 []map[string]interface{}
	err = json.Unmarshal(marshal3, &user4)
	if err != nil {
		return
	}
	fmt.Printf("反序列化后结果：%v\n", user4)

}

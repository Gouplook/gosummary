/**
 * @Author: yinjinlin
 * @File:  advMap
 * @Description:
 * @Date: 2021/6/17 上午10:12
 */

package gomap

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
	"sort"
)

type CardBase struct {
	Name     string  // 名称
	Price    float64 // 价格
	CardId   int     `mapstructure:"card_id"` // 如需要加下划线的，必须自己定义tag
	Clicks   int     // 点击量
	SalesNum int     // 销量 mapstructure 默认映射是小写的
}

type Card struct {
	CardBase
	Sales int // 销量
}

type CardMap struct {
	Name    string  // 名称
	Price   float64 // 价格
	CardId  int     `mapstructure:"card_id"` // 如需要加下划线的，必须自己定义tag 保持与之前的map字段一样
	Clicks  int     // 点击量
	SalesNu int     `mapstructure:"salesnum"` // 添加tag 保持与需要转到的map字段一致
}

// 定义一个map
func AdvanceMap() {
	maps := make([]map[string]interface{}, 1)
	maps[0] = make(map[string]interface{})
	maps[0]["name"] = "综合0"
	maps[0]["card_id"] = 11
	maps[0]["price"] = 100.0
	maps[0]["salesnum"] = 1003

	fmt.Println("maps=====", maps)
	var outStruct []CardMap // 测试map 与结构体字段不一致，转换不了

	_ = mapstructure.WeakDecode(maps, &outStruct)
	for _, v := range outStruct {
		fmt.Println("========key value =======")
		fmt.Println("Name  = ", v.Name)
		fmt.Println("CardId  = ", v.CardId)
		fmt.Println("SalesNu  = ", v.SalesNu)
		fmt.Println("Clicks  = ", v.Clicks)
		fmt.Println("Price  = ", v.Price)
	}

	// 追加（map形式追加）
	maps = append(maps, map[string]interface{}{
		"name":     "RcardId",
		"card_id":  14,
		"salesnum": 1006,
		"Clicks":   22,
	})

	// 追加后的数据
	fmt.Println("Append maps====", maps)

	maps2 := make([]map[string]interface{}, 2)
	// 必须make，否则panic: assignment to entry in nil map
	// 原因：未初始化的的value 是nil，
	maps2[0] = make(map[string]interface{})
	maps2[0]["name"] = "signle1"
	maps2[0]["card_id"] = 101
	maps2[0]["price"] = 101.0
	maps2[0]["salesnum"] = 1003101

	maps2[1] = make(map[string]interface{})
	maps2[1]["name"] = "signle2"
	maps2[1]["card_id"] = 201
	maps2[1]["price"] = 201.0
	maps2[1]["salesnum"] = 2003101

	// 切片追加切片
	maps = append(maps, maps2...)
	fmt.Println("maps =", maps)

	// map 转换成结构体 字段
	fmt.Println("====map 转换成结构体====")
	_ = mapstructure.WeakDecode(maps, &outStruct)

	for k, v := range outStruct {
		fmt.Println("========key value =======")
		fmt.Println("Name  = ", k, v.Name)
		fmt.Println("CardId  = ", k, v.CardId)
		fmt.Println("SalesNu  = ", k, v.SalesNu)
		fmt.Println("Clicks  = ", k, v.Clicks)
	}

}

// 双map
func AdvMapMap() {
	// 第一种初始化
	// maps := map[string]map[string]interface{}{}
	// 第二种初始化
	var maps map[string]map[string]interface{}
	maps = make(map[string]map[string]interface{})
	maps["001"] = make(map[string]interface{})
	maps["001"] = map[string]interface{}{
		"name": "Linux",
		"sex":  "M",
	}
	maps["002"] = make(map[string]interface{})
	maps["002"] = map[string]interface{}{
		"cardId": 20,
	}
	maps["003"] = make(map[string]interface{})
	maps["003"] = map[string]interface{}{
		"Id": 4,
	}
	maps["004"] = make(map[string]interface{})
	maps["003"] = map[string]interface{}{
		"Id": 6,
	}

	fmt.Println(maps)
}

// 切片
func SileIn() {
	cardId := make([]int, 0)
	cardId = append(cardId, 19)
	cardId = append(cardId, 12)
	cardId = append(cardId, 18)
	cardId = append(cardId, 98)
	cardId = append(cardId, 100)

	for _, card := range cardId {
		fmt.Println(card)
	}

	var slice []int
	fmt.Println("slice =", slice)
	slice1 := make([]int, 1)
	fmt.Println("slice1 =", slice1) // [0]
	slice2 := make([]int, 0)
	fmt.Println("slice2 =", slice2) // []

}

// 不用append 就覆盖
func SlieIn2() {
	cardId := make([]int, 0)
	cardId = []int{2}
	cardId = []int{4, 5}

	fmt.Println(cardId)

}

func MapMap() {
	level := map[int]int{}
	//mp := make(map[int]map[int]int)
	mp := map[int]map[int]int{}
	mp[22] = make(map[int]int)
	mp[23] = make(map[int]int)
	mp[24] = make(map[int]int)

	mp[22][1] = 2
	mp[22][0] = 1
	mp[22][2] = 0

	mp[23][0] = 1
	mp[23][2] = 5
	mp[23][1] = 2

	mp[24][0] = 1
	mp[24][1] = 2
	mp[24][2] = 0

	fmt.Println("before-", mp)
	if value, ok := mp[23][1]; ok {
		fmt.Println("value", value)
		mp[23][1] += 1
	} else {
		mp[23][1] = 1
	}
	fmt.Println(mp)
	type temp struct {
		key   int
		value int
	}

	tempList := make([]temp, 0)
	// map[22:map[0:1 1:2 2:0] 23:map[0:1 1:3 2:0] 24:map[0:1 1:2 2:0]]
	// 找出22 里面最大可以
	for k, v := range mp {
		for k2, va := range v {
			tempList = append(tempList, temp{
				key:   k2,
				value: va,
			})
		}
		sort.Slice(tempList, func(i, j int) bool {
			return tempList[i].value > tempList[j].value
		})
		level[k] = tempList[0].key
	}

	fmt.Println("level= ", level)

}

func MapMap2() {
	// i, itemId
	mp := make([]map[int]int, 10)
	for i := 0; i < 4; i++ {

		itemId := 160 + i
		sspId := 200 + 2*i
		//mp[i] = make(map[int]int)
		mp = append(mp, map[int]int{
			i:      i,
			sspId:  sspId,
			itemId: itemId,
		})

	}

	fmt.Println(mp)

}

type CardIcad struct {
	CardId int
	CardSn string
	Name   string
}

//
func MapSplitToStruct2(reple CardIcad) {
	cMak := map[string]interface{}{}
	cMak["CardId"] = 1002
	cMak["CardSn"] = "JS0003"

	fmt.Println("打印前：----", reple)
	_ = mapstructure.WeakDecode(cMak, &reple)
	fmt.Println("打印后：----", reple.CardSn)
	fmt.Println("打印后：----", reple.CardId)
}

type Person struct {
	Name string `mapstructure:"name"` // 可以映射，默认情况下，mapstructure 自动映射
	Age  int
	// Job  string

}

type Cat struct {
	Name  string
	Age   int
	Breed string
}

// 用法不常用
func MapToStruct() {
	datas := []string{`
    { 
      "type": "person",
      "name":"dj",
      "age":18,
      "job": "programmer"
    }
  `,
		`
    {
      "type": "cat",
      "name": "kitty",
      "age": 1,
      "breed": "Ragdoll"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		//	fmt.Println("m=",m)
		switch m["type"].(string) {
		case "person":
			var p Person
			err = mapstructure.Decode(m, &p)

			if err != nil {

			}
			fmt.Println("person", p)

		case "cat":
			var cat Cat
			// weakdecode 是Decode简化版
			_ = mapstructure.WeakDecode(m, &cat)
			fmt.Println("cat", cat)
		}
	}

}

// 正向转换
func MapToStruct1() {
	mapdata := map[string]interface{}{
		"name": "jack",
		"age":  19,
		"job":  "goland",
	}
	var structdata Person
	_ = mapstructure.WeakDecode(mapdata, &structdata)
	fmt.Println(structdata)
}

type Person2 struct {
	Name string
	Age  int
	Job  string `mapstructure:"omitempty"`
}

// 反向转换
func StructToMap() {
	p := &Person2{
		Name: "dj",
		Age:  18,
	}
	var m map[string]interface{}
	_ = mapstructure.WeakDecode(p, &m)
	fmt.Println("Main := ", p)
	fmt.Println("m : ", m)
	data, _ := json.Marshal(m)
	fmt.Println(string(data))
}

func DoubleMap() {
	var student map[string]map[string]string

	student = make(map[string]map[string]string)
	student["001"] = make(map[string]string)
	student["001"]["name"] = "tom"
	student["001"]["sex"] = "M"
	student["001"]["sex2"] = "M"
	student["001"]["sex3"] = "M"

	fmt.Println(student)
	student = make(map[string]map[string]string)
	fmt.Println(student)
}

type Cars struct {
	Name string
	Age  int
}

func MapSlice() {
	var sliceMap map[string][]string
	sliceMap = make(map[string][]string, 3)
	key := "1"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

func MapSliceStruct() {
	var sliceStruct map[string][]Cars
	sliceStruct = make(map[string][]Cars)
	key := "123"
	value, ok := sliceStruct[key]
	if !ok {
		value = make([]Cars, 0)
	}
	value = append(value, Cars{
		Name: "ZS",
		Age:  18,
	})
	value = append(value, Cars{
		Name: "WS",
		Age:  48,
	})
	value = append(value, Cars{
		Name: "YS",
		Age:  28,
	})
	sort.Slice(value, func(i, j int) bool {
		if value[i].Age > value[j].Age {
			return true
		}
		return false
	})
	sliceStruct[key] = value
	fmt.Println(sliceStruct)
}

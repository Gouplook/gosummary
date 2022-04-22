/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/4/19 13:05
@Description:

*********************************************/
package gojson

import (
	"encoding/json"
	"fmt"
)

// 场景1：
// 传过来的是结构体，需要序列化一下，存到数据库中
// specInfo, _ := json.Marshal(specIds)
// mSS.Field.F_spec_info: string(specInfo),

// 场景2：
// rpc添加描述 struct ---> string 前端传过来的数据 存储到数据库中 marshl
// rpc获取描述 string ---> struct 从数据库中取出数据给前端 Unmarshl（以结构体形式返回）

// Product 商品信息
type Product struct {
	Name      string `json:"name"`
	ProductID int64  `json:"product_id,omitempty"` // omitempty，tag里面加上omitempy，可以在序列化的时候忽略0值或者空值
	Number    int    `json:"-"`                    // 表示不进行序列化
	Price     float64
	IsOnSale  bool
}

// Marshal
func JsonMarshal() {
	p := &Product{}
	p.Name = "Xiao mi 6"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 0
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
}

func JsonMarshalBySile() {
	var splice []map[string]interface{}
	//创建切片需要装载的第一个map
	var m1 map[string]interface{}
	//make一些第一个map
	m1 = make(map[string]interface{})
	//第一个map的字段
	m1["name"] = "柯旗星"
	m1["age"] = 23
	m1["birthday"] = "2020-12-12"
	m1["sal"] = 12.2
	//把第一个map装进多层切片中
	splice = append(splice, m1)

	//进行第二个map的处理和装载
	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "柯旗星"
	m2["age"] = 23
	m2["birthday"] = "2020-12-12"
	m2["sal"] = 12.2
	splice = append(splice, m2)

	fmt.Println(splice)
	//对多层切片进行处理
	data, err := json.Marshal(splice)
	if err != nil {
		fmt.Printf("序列化失败")
	}
	//打印出来
	fmt.Printf("切片序列化后结果=%v\n", string(data))

}

// type，有些时候，我们在序列化或者反序列化的时候，可能结构体类型和需要的类型不一致，
// 这个时候可以指定,支持string,number和boolean
type Good struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,string"`
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

func JsonUnmarshal() {
	var data = `{"name":"Xiao mi 6","product_id":"10","number":"10000","price":"2499","is_on_sale":"true"}`
	// 一般从数据查到的数据或网络之间传送过来的。
	g := &Good{}
	err := json.Unmarshal([]byte(data), g)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("")
	fmt.Println(*g)
	fmt.Println(g.Name)
	fmt.Println(g.ProductID)
	fmt.Println(g.Number)
	fmt.Println(g.Price)
}

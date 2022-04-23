package goorm

import (
	"encoding/json"
	"fmt"
	"gosummary/goorm/gomode"
	"gosummary/goorm/models"
	"testing"
)

func TestCreateUser(t *testing.T) {
	initDb()
	// 新增用户
	//CreateUser()

	// 新增卡项
	CreateCards()
}

func TestFindUser(t *testing.T) {
	initDb()
	//FindUser()
	FindAllUser()
}

func TestAotiSql(t *testing.T) {
	initDb()
	AotiSql()
}

func TestGoModel_Insert(t *testing.T) {

	models.Init()
	cModel := new(gomode.GoModel).Init()
	cmDataAll := make([]map[string]interface{}, 0)
	//cmData := map[string]interface{}{
	//	cModel.Field.F_name:        "季度优惠券",
	//	cModel.Field.F_amount:      19.9,
	//	cModel.Field.F_num:         25,
	//	cModel.Field.F_time_type:   1,
	//	cModel.Field.F_receive_num: 0,
	//	cModel.Field.F_used_num:    0,
	//	cModel.Field.F_start_time:  0,
	//	cModel.Field.F_end_time:    0,
	//	cModel.Field.F_status:      1,
	//	cModel.Field.F_description: "有效期2022年第二季度",
	//}

	cmDataAll = append(cmDataAll, map[string]interface{}{
		cModel.Field.F_name:        "5月份度优惠券",
		cModel.Field.F_amount:      9.9,
		cModel.Field.F_num:         25,
		cModel.Field.F_time_type:   1,
		cModel.Field.F_receive_num: 0,
		cModel.Field.F_used_num:    0,
		cModel.Field.F_start_time:  0,
		cModel.Field.F_end_time:    0,
		cModel.Field.F_status:      1,
		cModel.Field.F_description: "有效期2022年5月份",
	})
	cmDataAll = append(cmDataAll, map[string]interface{}{
		cModel.Field.F_name:        "6月份度优惠券",
		cModel.Field.F_amount:      19.9,
		cModel.Field.F_num:         25,
		cModel.Field.F_time_type:   1,
		cModel.Field.F_receive_num: 0,
		cModel.Field.F_used_num:    0,
		cModel.Field.F_start_time:  0,
		cModel.Field.F_end_time:    0,
		cModel.Field.F_status:      1,
		cModel.Field.F_description: "有效期2022年6月份",
	})
	//n, err := cModel.Insert(cmData)
	//if err != nil {
	//	fmt.Println("Insert err")
	//}

	n, err := cModel.InsertAll(cmDataAll)
	if err != nil {
		fmt.Println("Insert err")
	}
	fmt.Println(n)

}

func TestGoModel_Update(t *testing.T) {

	models.Init()
	cModel := new(gomode.GoModel).Init()
	cmDataAll := make([]map[string]interface{}, 0)
	//cmData := map[string]interface{}{
	//	//cModel.Field.F_name:        "季度优惠券",1
	//	//cModel.Field.F_amount:      33.9,
	//	cModel.Field.F_num:         102,
	//	cModel.Field.F_time_type:   1,
	//	//cModel.Field.F_receive_num: 0,
	//	//cModel.Field.F_used_num:    0,
	//	//cModel.Field.F_start_time:  0,
	//	//cModel.Field.F_end_time:    0,
	//	//cModel.Field.F_status:      1,
	//	//cModel.Field.F_description: "有效期2022年第二季度",
	//}

	cmDataAll = append(cmDataAll, map[string]interface{}{
		//cModel.Field.F_name:        "5月份度优惠券",
		//cModel.Field.F_amount:      9.9,
		cModel.Field.F_num:  251,
		cModel.Field.F_type: 1,
		//cModel.Field.F_time_type: 1,
		//cModel.Field.F_receive_num: 0,
		//cModel.Field.F_used_num:    0,
		//cModel.Field.F_start_time:  0,
		//cModel.Field.F_end_time:    0,
		//cModel.Field.F_status:      1,
		//cModel.Field.F_description: "有效期2022年5月份",
	})
	cmDataAll = append(cmDataAll, map[string]interface{}{
		//cModel.Field.F_name:        "6月份度优惠券",
		//cModel.Field.F_amount:      19.9,
		cModel.Field.F_num: 252,
		//cModel.Field.F_time_type: 1,
		cModel.Field.F_type: 1,
		//cModel.Field.F_receive_num: 0,
		//cModel.Field.F_used_num:    0,
		//cModel.Field.F_start_time:  0,
		//cModel.Field.F_end_time:    0,
		//cModel.Field.F_status:      1,
		//cModel.Field.F_description: "有效期2022年6月份",
	})
	//n, err := cModel.Insert(cmData)
	//if err != nil {
	//	fmt.Println("Insert err")
	//}

	n, err := cModel.UpdateAllByWhere(1, cmDataAll, cModel.Field.F_num)
	if err != nil {
		fmt.Println("Insert err")
	}
	fmt.Println(n)

}

func TestGoModel_UpdateByIds(t *testing.T) {
	models.Init()
	cModel := new(gomode.GoModel).Init()
	// status 更新为已结束
	n, err := cModel.UpdateByIds([]int{1, 3}, 3)
	if err != nil {
		fmt.Println("UpdateByids err")
	}
	fmt.Println(n)

}

func TestGoModel_Find(t *testing.T) {
	models.Init()
	cModel := new(gomode.GoModel).Init()

	rs := cModel.Find("3C优惠券")

	//for _, v := range rs {
	//	//fmt.Println(k, v)
	//	rsByte, _ := json.Marshal(v)
	//	t.Log(string(rsByte))
	//}
	fmt.Println(rs)
	rsByte, _ := json.Marshal(rs)
	fmt.Printf("%v", string(rsByte))

}
func TestGoModel_SelectsByTypeId(t *testing.T) {
	models.Init()
	cModel := new(gomode.GoModel).Init()
	status := 3
	rs := cModel.SelectsByStatus(status)
	var s []map[string]interface{}
	s = make([]map[string]interface{}, 0)
	for _, v := range rs {
		s = append(s, v)
	}
	fmt.Println(s)
	rb, _ := json.Marshal(s)
	fmt.Printf("%s\n", string(rb))

}
func TestGoModel_SelectsByIds(t *testing.T) {
	models.Init()
	cModel := new(gomode.GoModel).Init()
	rs := cModel.SelectsByIds([]int{1, 2, 5})
	rb, _ := json.Marshal(rs)
	fmt.Println(string(rb))
}

func TestGoModeFindBetween(t *testing.T) {
	models.Init()
	cModel := new(gomode.GoModel).Init()
	rs := cModel.SelectsBetween(15.5, 35.5)
	fmt.Println(rs)
}

func TestGoModelSelectsLike(t *testing.T) {
	models.Init()
	cModel := new(gomode.GoModel).Init()
	rs := cModel.SelectsLike("月份")
	//rr, _ := json.Marshal(rs)

	fmt.Println(rs)
	//logs.Info(string(rr))
}

package gomode

import (
	"github.com/astaxie/beego/orm"
	"gosummary/goorm/models"
	"gosummary/tool"
)

//创建事务
//couponModel.Model.Begin()
//defer func() {
//	if err != nil {
//		couponModel.Model.RollBack()
//	}
//}()
//couponModel.Model.Commit()

// 累计
//func (g *GoModel) Count(coupon_id []int, where []base.WhereItem) (num int) {
//	if len(coupon_id) != 0 {
//		where = append(where, base.WhereItem{e.Field.F_id, []interface{}{"IN", coupon_id}})
//	}
//	return e.Model.Where(where).Count()
//}

//表结构体
type GoModel struct {
	Model *models.Model
	Field GoModelField
}

type GoModelField struct {
	T_table     string `default:"account_coupons"`
	F_acc_id    string `default:"acc_id"`
	F_coupon_id string `default:"coupon_id"`
}

func (g *GoModel) Init(ormer ...orm.Ormer) *GoModel {
	tool.ReflectModel(&g.Field)
	g.Model = models.NewModel(g.Field.T_table, ormer...)
	return g
}

//新增数据
func (g *GoModel) Insert(data map[string]interface{}) (result int, err error) {
	return g.Model.Data(data).Insert()
}

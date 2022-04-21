package gomode

//创建事务
//couponModel.Model.Begin()
//defer func() {
//	if err != nil {
//		couponModel.Model.RollBack()
//	}
//}()
//couponModel.Model.Commit()

type GoModel struct {
}

// 累计
//func (g *GoModel) Count(coupon_id []int, where []base.WhereItem) (num int) {
//	if len(coupon_id) != 0 {
//		where = append(where, base.WhereItem{e.Field.F_id, []interface{}{"IN", coupon_id}})
//	}
//	return e.Model.Where(where).Count()
//}

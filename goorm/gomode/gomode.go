package gomode

import (
	"github.com/astaxie/beego/orm"
	"gosummary/goorm/models"
	"gosummary/tool"
	"time"
)

//表结构体
type GoModel struct {
	Model *models.Model
	Field GoModelField
}

type GoModelField struct {
	T_table       string `default:"coupons"`
	F_id          string `default:"id"`
	F_name        string `default:"name"`
	F_amount      string `default:"amount"`
	F_num         string `default:"num"`
	F_receive_num string `default:"receive_num"`
	F_used_num    string `default:"used_num"`
	F_time_type   string `default:"time_type"`
	F_start_time  string `default:"start_time"`
	F_end_time    string `default:"end_time"`
	F_limit_num   string `default:"limit_num"`
	F_status      string `default:"status"`
	F_type        string `default:"type"`
	F_scope_type  string `default:"scope_type"`
	F_description string `default:"description"`
	F_create_time string `default:"create_time"`
	F_update_time string `default:"update_time"`
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

// 批量插入
func (g *GoModel) InsertAll(data []map[string]interface{}) (result int, err error) {
	return g.Model.InsertAll(data)
}

// 更新数据
func (g *GoModel) Update(id int, data map[string]interface{}) (result int, err error) {
	if len(data) == 0 {
		return 0, nil
	}
	where := []models.WhereItem{
		{g.Field.F_id, id},
	}
	return g.Model.Where(where).Data(data).Update()
}

// 批量更新状态和时间(更新某个字段）
func (g *GoModel) UpdateByIds(ids []int, status int) (result int, err error) {
	return g.Model.Where([]models.WhereItem{
		{g.Field.F_id, []interface{}{"IN", ids}},
	}).Data(map[string]interface{}{
		g.Field.F_status:      status,
		g.Field.F_update_time: time.Now().Unix(),
	}).Update()
}

// 批量更新
func (g *GoModel) UpdateAllByWhere(id int, datas []map[string]interface{}, field string) (result int, err error) {
	where := []models.WhereItem{
		{g.Field.F_type, id},
	}
	return g.Model.Where(where).UpdateCase(datas, field)
}

//根据条件查找一条数据
func (g *GoModel) Find(name string) (result map[string]interface{}) {
	return g.Model.Where([]models.WhereItem{
		{g.Field.F_name, name},
	}).Find()
}
func (g *GoModel) SelectsBetween(amountStar, amountEnd float64) []map[string]interface{} {
	return g.Model.Where([]models.WhereItem{
		{g.Field.F_amount, []interface{}{"BETWEEN", []float64{amountStar, amountEnd}}},
	}).Select()
}

// 通配符查询
func (g *GoModel) SelectsLike(name string) []map[string]interface{} {
	return g.Model.Where([]models.WhereItem{
		{g.Field.F_name, []interface{}{"like", "%" + name + "%"}},
	}).Select()
}

// 批量查询
func (g *GoModel) SelectsByStatus(status int) (result []map[string]interface{}) {
	rs := g.Model.Where([]models.WhereItem{
		{g.Field.F_status, status},
	}).Select()
	return rs
}

// 根据多个条件批量查询
func (g *GoModel) SelectsByIds(ids []int) (result []map[string]interface{}) {
	return g.Model.Where([]models.WhereItem{
		{g.Field.F_id, []interface{}{"IN", ids}},
	}).Select()
}

// 分页批量查询-列表
func (g *GoModel) SelectsByLimit(status int, start int, limit int) []map[string]interface{} {
	return g.Model.Where([]models.WhereItem{
		{g.Field.F_status, status},
	}).OrderBy(g.Field.F_id+" DESC ").Limit(start, limit).Select()
}

// 获取总数量-列表
func (g *GoModel) GetNum(status int) (count int) {
	return g.Model.Where([]models.WhereItem{
		{g.Field.F_status, status},
	}).Count(g.Field.F_id)
}

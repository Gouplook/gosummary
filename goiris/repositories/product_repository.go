package repositories

import (
	"database/sql"
	"gosummary/goiris/common"
	"gosummary/goiris/datamodels"
	"strconv"
)

//第一步，先开发对应的接口
//第二步，实现定义的接口
type IProduct interface {
	Conn() error
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewProductManger(table string, db *sql.DB) IProduct {
	return &ProductManager{
		table:     table,
		mysqlConn: db,
	}
}

func (p *ProductManager) Conn() (err error) {
	if p.mysqlConn == nil {
		db, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = db
	}
	if p.table == "" {
		p.table = "product"
	}
	return
}

//插入
func (p *ProductManager) Insert(product *datamodels.Product) (productId int64, err error) {
	// 1. 判断数据库是否连接
	if err = p.Conn(); err != nil {
		return
	}
	// 2. 准备sql
	sql := "INSERT product SET productName=?,productNum=?,productImage=?,productUrl=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	// 3. 传入参数
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

//商品的删除
func (p *ProductManager) Delete(productID int64) bool {
	if err := p.Conn(); err != nil {
		return false
	}
	sql := "DELETE FROM product WHERE ID=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return false
	}
	_, err = stmt.Exec(strconv.FormatInt(productID, 10))
	if err != nil {
		return false
	}
	return true
}

//商品的更新
func (p *ProductManager) Update(product *datamodels.Product) (err error) {
	if err = p.Conn(); err != nil {
		return
	}
	sql := "UPDATE " + p.table + " SET productName=?,productNum=?,productImage=?,productUrl=? where ID=" + strconv.FormatInt(product.ID, 10)
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return
	}
	return nil
}

//根据商品ID查询商品
func (p *ProductManager) SelectByKey(productID int64) (productResult *datamodels.Product, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	sql := "SELECT * FROM " + p.table + " WHERE ID = " + strconv.FormatInt(productID, 10)
	row, err := p.mysqlConn.Query(sql)
	defer row.Close()
	if err != nil {
		return &datamodels.Product{}, err
	}
	rs := common.GetResultRow(row)
	if len(rs) == 0 {
		return &datamodels.Product{}, err
	}
	productResult = &datamodels.Product{}
	common.DataToStructByTagSql(rs, productResult)
	return
}

//获取所有商品
func (p *ProductManager) SelectAll() (productArray []*datamodels.Product, errProduct error) {
	if err := p.Conn(); err != nil {
		return nil, err
	}
	sql := "Select * from " + p.table
	rows, err := p.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}

	for _, v := range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		productArray = append(productArray, product)
	}
	return
}

package repositories

import (
	"database/sql"
	"gosummary/goiris/common"
	"gosummary/goiris/datamodels"
)

type IOrderRepository interface {
	Conn() error
	Insert(order *datamodels.Order) (int64, error)
	Update(*datamodels.Order) error
	Delete(int64) bool
	SelectByKey(int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}

type OrderManagerRepository struct {
	table string
	db    *sql.DB
}

func NewOrderManagerRepository(table string, db *sql.DB) IOrderRepository {
	return &OrderManagerRepository{
		table: table,
		db:    db,
	}
}

func (o *OrderManagerRepository) Conn() error {
	if o.db == nil {
		db, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		o.db = db
	}
	if o.table == "" {
		o.table = "order"
	}
	return nil
}
func (o *OrderManagerRepository) Insert(order *datamodels.Order) (productID int64, err error) {
	if err := o.Conn(); err != nil {
		return 0, err
	}
	sql := "INSERT " + o.table + " SET userID=?,productID=?,orderStatus=?"
	stmt, errStmt := o.db.Prepare(sql)
	if errStmt != nil {
		return productID, errStmt
	}
	result, errResult := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	if errResult != nil {
		return productID, errResult
	}
	return result.LastInsertId()
}
func (o *OrderManagerRepository) Delete(orderID int64) (isOk bool) {
	// todo
	return
}
func (o *OrderManagerRepository) Update(order *datamodels.Order) (err error) {

	// todo

	return

}
func (o *OrderManagerRepository) SelectByKey(orderID int64) (order *datamodels.Order, err error) {
	// todo

	return

}
func (o *OrderManagerRepository) SelectAll() (orderArray []*datamodels.Order, err error) {

	return
}

func (o *OrderManagerRepository) SelectAllWithInfo() (OrderMap map[int]map[string]string, err error) {

	return
}
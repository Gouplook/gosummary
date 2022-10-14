package repositories

import (
	"database/sql"
	"gosummary/goiris/common"
	"gosummary/goiris/datamodels"
)

type IUserRepository interface {
	Conn() error
	Insert(user *datamodels.User) (userId int64, err error)
	Select(userName string) (user *datamodels.User, err error)
}

type UserManageRepository struct {
	table string  // 表名
	db    *sql.DB // 连接数据库
}

func NewUserManageRepository(table string, db *sql.DB) IUserRepository {
	return &UserManageRepository{
		table: table,
		db:    db,
	}
}

func (u *UserManageRepository) Conn() (err error) {
	if u.db == nil {
		db, errMysql := common.NewMysqlConn()
		if errMysql != nil {
			return errMysql
		}
		u.db = db
	}
	if u.table == "" {
		u.table = "user"
	}
	return
}

func (u *UserManageRepository) Select(userName string) (user *datamodels.User, err error) {

	return
}

func (u *UserManageRepository) Insert(user *datamodels.User) (userId int64, err error) {

	return
}

func (u *UserManageRepository) SelectByID(userId int64) (user *datamodels.User, err error) {

	return
}

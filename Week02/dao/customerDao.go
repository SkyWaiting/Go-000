package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type CustomerDao struct {
}

func NewCustomerDao() *CustomerDao {
	return &CustomerDao{}
}

type Customer struct {
	Id   string
	Name string
	age  int
}

var db *sql.DB

func (cusDao *CustomerDao) GetCustomerById(id string) (customer *Customer, err error) {
	err = sql.ErrNoRows
	return customer, errors.Wrap(err, "Can't find thie customer")
}

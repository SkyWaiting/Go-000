package service

import (
	"Week02/dao"
	"github.com/pkg/errors"
)

type CustomerService struct {
	customerDao *dao.CustomerDao
}

func NewCustomerService() *CustomerService {
	return &CustomerService{dao.NewCustomerDao()}
}

func (cs *CustomerService) GetCustomerById(id string) (c *dao.Customer, err error) {
	cs = NewCustomerService()
	c, err = cs.customerDao.GetCustomerById(id)
	return c, errors.Wrapf(err, "CustomerService: customer %s can't find", id)
}

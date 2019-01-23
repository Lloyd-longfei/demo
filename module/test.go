package module

import (
	"demo/database"
	"demo/model"
)

type Param struct {
	OrderId string
}

func (p *Param) Insert() error {
	orm, _ := database.OrmLink().Create()
	var test model.Test
	test.Name = "ceshi"
	test.OrderId = p.OrderId
	_, err := orm.Insert(&test)

	return err
}

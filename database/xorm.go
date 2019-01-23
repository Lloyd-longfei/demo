package database

import (
	"demo/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type DatabaseConf struct {
	Conn []string
}

func (env *DatabaseConf) Create() (*xorm.EngineGroup, error) {
	eg, err := xorm.NewEngineGroup("mysql", env.Conn)
	return eg, err
}

//抽象

func OrmLink() *DatabaseConf {
	v := conf.ReadConfFile()
	return &DatabaseConf{
		Conn: []string{
			v.GetString("database.read.username") + ":" + v.GetString("database.read.password") + "@tcp(" + v.GetString("database.read.host") + ":" + v.GetString("database.read.port") + ")/" + v.GetString("database.read.db") + "?charset=utf8&parseTime=True&loc=Local",
		},
	}
}

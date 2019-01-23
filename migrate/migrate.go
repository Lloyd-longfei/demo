package migrate

import (
	"demo/database"
	"demo/model"

	_ "github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

//创建表
func Migrate() {
	engine, err := database.OrmLink().Create()
	defer engine.Close()
	if err != nil {
		log.Error(err.Error())

	}
	err = engine.Sync(new(model.Test))
	if err != nil {
		log.Error(err.Error())

	}
	if err == nil {
		log.Info("run migrate success!")
	}
}

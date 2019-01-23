package handler

import (
	"demo/module"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type TestParam struct {
	OrderId string `json:"order_id" validate:"required"`
	Name    string `json:"name" validate:"required"`
}

// Name will print hello name
// @Summary Print
// @Accept json
// @Tags Name
// @Security Bearer
// @Produce  json
// @Param name path string true "name"
// @Resource Name
// @Router /insert [post]
func Test(c *gin.Context) {
	var json TestParam
	c.BindJSON(&json)
	validate := validator.New()
	if err := validate.Struct(json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":         err.Error(),
			"status_code": 412,
			"data":        map[string]string{},
		})
		return
	}
	mod := module.Param{OrderId: json.OrderId}
	if err := mod.Insert(); err != nil {
		//c.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg":         err.Error(),
			"status_code": 500,
			"data":        map[string]string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":         "success",
		"status_code": 200,
		"data":        map[string]string{},
	})
}

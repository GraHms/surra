package main

import (
	"github.com/surra"
	"net/http"
)

type Extra struct {
	Id     *string `json:"id" binding:"required"`
	Name   *string `json:"name" binding:"required"`
	Msisdn *string `json:"msisdn" binding:"required"`
}
type BodyModel struct {
	Id     *string `json:"id" binding:"required"`
	Name   *string `json:"name" binding:"required"`
	Msisdn *string `json:"msisdn" binding:"required"`
	Extra  Extra   `json:"extra"`
}

func main() {
	r := surra.Principal()
	r.GET("/hello", func(c *surra.Context) {
		c.JSON(http.StatusOK, surra.KeyVal{
			"message": "pong",
		})
	})
	r.POST("/hello", func(c *surra.Context) {
		var model BodyModel
		err := c.BindTMFJson(&model)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, model)
	})
	r.Run()

}

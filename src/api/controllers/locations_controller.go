package controllers

import (
	"net/http"

	"../services"
	"github.com/gin-gonic/gin"
)

func GetCountry(c *gin.Context) {
	country, err := services.GetCountry(c.Param("country_id"))

	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, country)
}

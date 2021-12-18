package controllers

import (
	"github.com/gin-gonic/gin"
	"merk_microservices/common"
	"merk_microservices/daos"
	"merk_microservices/models"
	"net/http"
)

type Merk struct {
	daos daos.Merk
}

func (u *Merk) MerkCreate(c *gin.Context) {

	params := models.CreateMerk{}

	response := models.Response{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.ApiMessage = "validation " + mess
		c.JSON(400, response)
	} else {

		data, err := u.daos.MerkCreate(params)

		if err != nil {
			response.ApiStatus = 0
			response.ApiMessage = err.Error()
			c.JSON(500, response)
		} else {
			response.ApiStatus = 1
			response.Data = data
			response.ApiMessage = common.StatusSukses
			c.JSON(http.StatusOK, response)

		}

	}

}

func (u *Merk) GetDataMerk(c *gin.Context) {

	response := models.Response{}


	data, err := u.daos.MerkGet()

	if err != nil {
		response.ApiStatus = 0
		response.ApiMessage = err.Error()
		c.JSON(500, response)
	} else {
		response.ApiStatus = 1
		response.Data = data
		response.ApiMessage = common.StatusSukses
		c.JSON(http.StatusOK, response)

	}


}

func (u *Merk) MerkUpdate(c *gin.Context) {

	params := models.UpdateMerk{}

	response := models.Response{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.ApiMessage = "validation " + mess
		c.JSON(400, response)
	} else {

		data, err := u.daos.MerkUpdate(params)

		if err != nil {
			response.ApiStatus = 0
			response.ApiMessage = err.Error()
			c.JSON(500, response)
		} else {
			response.ApiStatus = 1
			response.Data = data
			response.ApiMessage = common.StatusSukses
			c.JSON(http.StatusOK, response)

		}

	}

}

func (u *Merk) MerkDelete(c *gin.Context) {

	params := models.DeleteMerk{}

	response := models.Response{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.ApiMessage = "validation " + mess
		c.JSON(400, response)
	} else {

		data, err := u.daos.MerkDelete(params)

		if err != nil {
			response.ApiStatus = 0
			response.ApiMessage = err.Error()
			c.JSON(500, response)
		} else {
			response.ApiStatus = 1
			response.Data = data
			response.ApiMessage = common.StatusSukses
			c.JSON(http.StatusOK, response)

		}

	}

}
package controllers

import (
	"github.com/gin-gonic/gin"
	"kategori_microservice/common"
	"kategori_microservice/daos"
	"kategori_microservice/models"
	"net/http"
)

type Kategori struct {
	daos daos.Kategori
}

func (u *Kategori) KategoriCreate(c *gin.Context) {

	params := models.CreateKategori{}

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

		data, err := u.daos.KategoriCreate(params)

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

func (u *Kategori) GetDataKategori(c *gin.Context) {

	response := models.Response{}


	data, err := u.daos.KategoriGet()

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

func (u *Kategori) KategoriUpdate(c *gin.Context) {

	params := models.UpdateKategori{}

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

		data, err := u.daos.KategoriUpdate(params)

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

func (u *Kategori) KategoriDelete(c *gin.Context) {

	params := models.DeleteKategori{}

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

		data, err := u.daos.KategoriDelete(params)

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
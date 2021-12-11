package controllers

import (
	"github.com/gin-gonic/gin"
	"keranjang_microservice/common"
	"keranjang_microservice/daos"
	"keranjang_microservice/models"
	"net/http"
)

type Keranjang struct {
	daos daos.Keranjang
}

func (u *Keranjang) KeranjangCreate(c *gin.Context) {

	params := models.CreateKeranjang{}

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

		data, err := u.daos.KeranjangCreate(params)

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

func (u *Keranjang) GetDataKeranjang(c *gin.Context) {

	response := models.Response{}
	params := models.KeranjangGet{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.ApiMessage = "validation " + mess
		c.JSON(400, response)
	} else {

		data, err := u.daos.KeranjangGet(params)

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

func (u *Keranjang) KeranjangDelete(c *gin.Context) {

	params := models.DeleteKeranjang{}

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

		data, err := u.daos.KeranjangDelete(params)

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

//
//func (u *Produk) ProdukUpdate(c *gin.Context) {
//
//	params := models.UpdateProduk{}
//
//	response := models.Response{}
//
//	err := c.ShouldBind(&params)
//
//	if err != nil {
//		var mess string
//		if err != nil {
//			mess = mess + err.Error()
//		}
//
//		response.ApiMessage = "validation " + mess
//		c.JSON(400, response)
//	} else {
//
//		data, err := u.daos.ProdukUpdate(params)
//
//		if err != nil {
//			response.ApiStatus = 0
//			response.ApiMessage = err.Error()
//			c.JSON(500, response)
//		} else {
//			response.ApiStatus = 1
//			response.Data = data
//			response.ApiMessage = common.StatusSukses
//			c.JSON(http.StatusOK, response)
//
//		}
//
//	}
//
//}
//
package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user_microservices/common"
	"user_microservices/daos"
	"user_microservices/models"
)

type User struct {
	daos daos.User
}

func (u *User) UserCreate(c *gin.Context) {

	params := models.CreateUser{}

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

		data, err := u.daos.UserCreate(params)

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

func (u *User) GetDataUser(c *gin.Context) {

	response := models.Response{}
	params := models.GetUser{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.ApiMessage = "validation " + mess
		c.JSON(400, response)
	} else {

		data, err := u.daos.UserGet(params)

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

func (u *User) UserUpdate(c *gin.Context) {

	params := models.UpdateUser{}

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

		data, err := u.daos.UserUpdate(params)

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

func (u *User) UserDelete(c *gin.Context) {

	params := models.DeleteUser{}

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

		data, err := u.daos.UserDelete(params)

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
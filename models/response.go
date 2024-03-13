package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseModelWithData struct {
	Data interface{} `json:"data"`
	ResponseModel
}

func ResponseSuccessWithData(c *gin.Context, data interface{}) {

	response := ResponseModelWithData{}
	response.Code = http.StatusOK
	response.Data = data
	response.Message = "Success"

	c.JSON(http.StatusOK, response)

}

func ResponseSuccessCreated(c *gin.Context, data interface{}) {

	response := ResponseModelWithData{}
	response.Code = http.StatusCreated
	response.Data = data
	response.Message = "Created"

	c.JSON(http.StatusCreated, response)

}

func ResponseSuccess(c *gin.Context) {

	response := ResponseModel{}
	response.Code = http.StatusOK
	response.Message = "Success"

	c.JSON(http.StatusOK, response)

}

func ResponseError(c *gin.Context, err string, code int) {

	response := ResponseModel{}
	response.Code = code
	response.Message = err

	c.JSON(code, response)

}

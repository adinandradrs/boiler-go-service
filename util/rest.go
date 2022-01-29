package util

import (
	"net/http"

	commodel "github.com/adinandradrs/codefun-go-common/model"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func ThrowBadError(err string, context *gin.Context) {
	out := commodel.GetStatusCodeResponse(http.StatusBadRequest, commodel.RestResponse{
		Data:    nil,
		Result:  false,
		Message: err,
	})
	context.JSON(out.Code, out.Response)
}

func ThrowAnyError(inp interface{}, context *gin.Context) {
	out := commodel.StatusCodeResponse{}
	mapstructure.Decode(inp, &out)
	context.JSON(out.Code, out.Response)
}

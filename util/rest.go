package util

import (
	commodel "github.com/adinandradrs/codefun-go-common/model"
	"github.com/gin-gonic/gin"
)

func ThrowError(code int, err string, ctx *gin.Context) {
	inp := commodel.StatusCodeResponse{
		Code: code,
		Response: commodel.RestResponse{
			Message: err,
			Result:  false,
		},
	}
	ctx.JSON(inp.Code, inp.Response)
}

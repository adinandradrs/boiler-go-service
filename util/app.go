package util

import (
	"crypto/rand"
	"math/big"
	"net/http"

	combase "github.com/adinandradrs/boiler-go-common"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func ThrowBadError(err string, context *gin.Context) {
	out := combase.GetStatusCodeResponse(http.StatusBadRequest, combase.RestResponse{
		Data:    nil,
		Result:  false,
		Message: err,
	})
	context.JSON(out.Code, out.Response)
}

func ThrowAnyError(inp interface{}, context *gin.Context) {
	out := combase.StatusCodeResponse{}
	mapstructure.Decode(inp, &out)
	context.JSON(out.Code, out.Response)
}

func GenerateOtp(digit int) (string, error) {
	nums := "012345679"
	bytes := make([]byte, digit)
	for i := 0; i < digit; i++ {
		num, err := rand.Int(rand.Reader,
			big.NewInt(int64(len(nums))),
		)
		if err != nil {
			return "", err
		}
		bytes[i] = nums[num.Int64()]
	}
	return string(bytes), nil
}

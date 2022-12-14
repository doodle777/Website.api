package bilibili

import (
	"duanxu.tech/website/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetVideoList(context *gin.Context) {
	request := common.PageRequest[int]{
		Param:     3039861,
		PageIndex: 1,
		PageSize:  30,
	}
	response := getVideoListDomain(request)
	context.SecureJSON(http.StatusOK, response)
}

func GetVideoDetail(context *gin.Context) {
	aid, _ := strconv.Atoi(context.Query("aid"))

	request := common.CommonRequest[int]{
		Param: aid,
	}
	response := getVideoDetailDomain(request)
	context.SecureJSON(http.StatusOK, response)
}

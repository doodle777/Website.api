package bilibili

import (
	"duanxu.tech/website/common"
)

func getVideoListDomain(param common.PageRequest[int]) common.PageResponse[BilibiliVideoDetail] {

	request := BilibiliVideoListRequest{
		Mid: param.Param,
		Ps:  param.PageSize,
		Pn:  param.PageIndex,
	}

	response := getVideoList(request)

	result := common.PageResponse[BilibiliVideoDetail]{
		Data:      response.Data.List.Vlist,
		PageSize:  response.Data.Page.Ps,
		PageIndex: response.Data.Page.Pn,
		TotalRow:  response.Data.Page.Count,
	}
	return result
}

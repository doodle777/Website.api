package bilibili

import (
	"duanxu.tech/website/common"
)

func getVideoListDomain(param common.PageRequest[int]) common.PageResponse[VideoDetail] {

	request := VideoListRequest{
		Mid: param.Param,
		Ps:  param.PageSize,
		Pn:  param.PageIndex,
	}

	response := getVideoList(request)

	result := common.PageResponse[VideoDetail]{
		Data:      response.Data.List.Vlist,
		PageSize:  response.Data.Page.Ps,
		PageIndex: response.Data.Page.Pn,
		TotalRow:  response.Data.Page.Count,
	}
	return result
}

func getVideoDetailDomain(param common.CommonRequest[int]) common.CommonResponse[VideoDetail] {

	request := VideoDetailRequest{
		Aid: param.Param,
	}

	response := getVideoDetail(request)

	result := common.CommonResponse[VideoDetail]{
		Data: VideoDetail{
			Aid:      param.Param,
			Cid:      response.Data[0].Cid,
			Duration: response.Data[0].Duration,
		},
	}
	return result
}

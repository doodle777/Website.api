package bilibili

import (
	"duanxu.tech/website/common"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func getVideoList(param BilibiliVideoListRequest) BilibiliVideoListResponse {
	request, _ := http.NewRequest("GET", common.ApiBilibiliVideoList, nil)
	query := request.URL.Query()
	query.Add("mid", strconv.Itoa(param.Mid))
	query.Add("ps", strconv.Itoa(param.Ps))
	query.Add("pn", strconv.Itoa(param.Pn))
	query.Add("tid", strconv.Itoa(0))
	query.Add("order", "pubdate")
	query.Add("order_avoided", "true")
	query.Add("jsonp", string("jsonp"))
	request.URL.RawQuery = query.Encode()

	response, _ := http.DefaultClient.Do(request)
	body, _ := io.ReadAll(response.Body)

	result := BilibiliVideoListResponse{}
	_ = json.Unmarshal(body, &result)

	return result
}

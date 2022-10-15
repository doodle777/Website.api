package bilibili

// BilibiliVideoListRequest Bilibili视频列表查询请求体
type BilibiliVideoListRequest struct {
	Mid          int
	Ps           int
	Tid          int
	Pn           int
	Keyword      string
	Order        string
	OrderAvoided bool
	Jsonp        string
}

// BilibiliVideoListResponse Bilibili视频列表查询响应体
type BilibiliVideoListResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		List struct {
			Tlist struct {
			} `json:"tlist"`
			Vlist []BilibiliVideoDetail `json:"vlist"`
		} `json:"list"`
		Page struct {
			Pn    int `json:"pn"`
			Ps    int `json:"ps"`
			Count int `json:"count"`
		} `json:"page"`
		EpisodicButton struct {
			Text string `json:"text"`
			Uri  string `json:"uri"`
		} `json:"episodic_button"`
		IsRisk      bool        `json:"is_risk"`
		GaiaResType int         `json:"gaia_res_type"`
		GaiaData    interface{} `json:"gaia_data"`
	} `json:"data"`
}

// BilibiliVideoDetail Bilibili视频详情
type BilibiliVideoDetail struct {
	Comment        int    `json:"comment"`
	Typeid         int    `json:"typeid"`
	Play           int    `json:"play"`
	Pic            string `json:"pic"`
	Subtitle       string `json:"subtitle"`
	Description    string `json:"description"`
	Copyright      string `json:"copyright"`
	Title          string `json:"title"`
	Review         int    `json:"review"`
	Author         string `json:"author"`
	Mid            int    `json:"mid"`
	Created        int    `json:"created"`
	Length         string `json:"length"`
	Duration       int    `json:"duration"`
	VideoReview    int    `json:"video_review"`
	Aid            int    `json:"aid"`
	Bvid           string `json:"bvid"`
	Cid            int    `json:"cid"`
	HideClick      bool   `json:"hide_click"`
	IsPay          int    `json:"is_pay"`
	IsUnionVideo   int    `json:"is_union_video"`
	IsSteinsGate   int    `json:"is_steins_gate"`
	IsLivePlayback int    `json:"is_live_playback"`
}

// BilibiliVideoDetailRequest Bilibili视频详情查询请求体
type BilibiliVideoDetailRequest struct {
	Aid int
}

// BilibiliVideoDetailResponse Bilibili视频详情查询响应体
type BilibiliVideoDetailResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    []struct {
		Cid       int    `json:"cid"`
		Page      int    `json:"page"`
		From      string `json:"from"`
		Part      string `json:"part"`
		Duration  int    `json:"duration"`
		Vid       string `json:"vid"`
		Weblink   string `json:"weblink"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
	} `json:"data"`
}

package statistics

import "go-okx/rest/api"

const GetCOIVolumeLimitNumPerSec = 2.5
const GetCOIVolumeLimitRule = "IP"

func NewGetCOIVolume(param *GetCOIVolumeParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/rubik/stat/contracts/open-interest-volume",
		Method: api.MethodGet,
		Param:  param,
	}, &GetCOIVolumeResponse{}
}

type GetCOIVolumeParam struct {
	Ccy    string `url:"ccy"`
	Begin  string `url:"begin,omitempty"`
	End    string `url:"end,omitempty"`
	Period string `url:"period,omitempty"`
}

type COIVolumeData struct {
	Item [3]string
}

// 获取产品深度
// 获取产品深度列表(订单列表)
type GetCOIVolumeResponse struct {
	api.Response
	Data []COIVolumeData `json:"data"`
}

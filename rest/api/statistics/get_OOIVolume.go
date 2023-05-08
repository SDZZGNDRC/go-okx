package statistics

import "go-okx/rest/api"

const GetOOIVolumeLimitNumPerSec = 2.5
const GetOOIVolumeLimitRule = "IP"

func NewGetOOIVolume(param *GetOOIVolumeParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/rubik/stat/option/open-interest-volume",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOOIVolumeResponse{}
}

type GetOOIVolumeParam struct {
	Ccy    string `url:"ccy"`
	Begin  string `url:"begin,omitempty"`
	End    string `url:"end,omitempty"`
	Period string `url:"period,omitempty"`
}

type OOIVolumeData struct {
	Item [3]string
}

// 获取产品深度
// 获取产品深度列表(订单列表)
type GetOOIVolumeResponse struct {
	api.Response
	Data []OOIVolumeData `json:"data"`
}

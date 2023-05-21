package status

import "github.com/SDZZGNDRC/go-okx/rest/api"

const GetStatusLimitNumPerSec = 0.2
const GetStatusLimitRule = "IP"

func NewGetStatus(param *GetStatusParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/system/status",
		Method: api.MethodGet,
		Param:  param,
	}, &GetStatusResponse{}
}

type GetStatusParam struct {
	Status string `json:"state"`
}

type GetStatusResponse struct {
	api.Response
	Data []struct {
		Title        string `json:"title"`
		State        string `json:"state"`
		Begin        string `json:"begin"`
		End          string `json:"end"`
		PreOpenBegin string `json:"preOpenBegin"`
		Href         string `json:"href"`
		ServiceType  string `json:"serviceType"`
		System       string `json:"system"`
		ScheDesc     string `json:"scheDesc"`
	} `json:"data"`
}

package market

import "go-okx/rest/api"

func NewGetP24Volume(param *GetP24VolumeParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/platform-24-volume",
		Method: api.MethodGet,
		Param:  param,
	}, &GetP24VolumeResponse{}
}

type GetP24VolumeParam struct {
}

type GetP24VolumeResponse struct {
	api.Response
	Data []P24VolumeDataItem `json:"data"`
}

type P24VolumeDataItem struct {
	VolUsd      string `json:"volUsd,omitempty"`
	VolCny      string `json:"volCny,omitempty"`
	BlockVolUsd string `json:"blockVolUsd,omitempty"`
	BlockVolCny string `json:"blockVolCny,omitempty"`
}

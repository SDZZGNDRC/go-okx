package custom

import "github.com/SDZZGNDRC/go-okx/rest/api"

func NewGetFileContent(param *GetFileContentParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/cdn/okex/traderecords/trades/daily/" + param.Date + "/" + param.FileName,
		Method: api.MethodGet,
		Param:  param,
	}, nil
}

type GetFileContentParam struct {
	Date     string
	FileName string
}

package custom

import "github.com/SDZZGNDRC/go-okx/rest/api"

func NewGetFileList(param *GetFileListParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/priapi/v5/broker/public/v2/orderRecord",
		Method: api.MethodGet,
		Param:  param,
	}, &GetFileListResponse{}
}

type GetFileListParam struct {
	Path       string `url:"path,omitempty"`
	NextMarker string `url:"nextMarker,omitempty"`
	Size       string `url:"size,omitempty"`
}

type GetFileListResponse struct {
	api.Response
	Data struct {
		IsTruncate     bool           `json:"isTruncate,omitempty"` // Indicates whether the list was truncated.
		NextMarker     string         `json:"nextMarker,omitempty"`
		RecordFileList []FileListItem `json:"recordFileList,omitempty"`
	} `json:"data"`
}

type FileListItem struct {
	FileName       string `json:"fileName,omitempty"`
	FileUploadDate string `json:"fileUploadDate,omitempty"`
}

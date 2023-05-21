package ws

import "github.com/SDZZGNDRC/go-okx/common"

type Args struct {
	Channel    string `json:"channel"`
	InstId     string `json:"instId,omitempty"`
	InstType   string `json:"instType,omitempty"`
	InstFamily string `json:"InstFamily,omitempty"`
	Uly        string `json:"uly,omitempty"`
}

// 登录参数, 用于身份验证
type ArgsLogin struct {
	ApiKey     string `json:"apiKey"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}

func NewArgsLoginFromAuth(auth common.Auth) *ArgsLogin {
	signature := auth.Signature("GET", "/users/self/verify", "", true)
	return &ArgsLogin{
		ApiKey:     auth.ApiKey,
		Passphrase: auth.Passphrase,
		Sign:       signature.Build(),
		Timestamp:  signature.Timestamp,
	}
}

package finance

import (
	"encoding/json"
	"fmt"
	"go-okx/rest/api"
)

func NewGetProjectCurrency(param *GetProjectCurrencyParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/v2/asset/balance/project-currency",
		Method: api.MethodGet,
		Param:  param,
	}, &GetProjectCurrencyResponse{}
}

type GetProjectCurrencyParam struct {
	T string `url:"t,omitempty"`
}

type GetProjectCurrencyResponse struct {
	api.Response2
	Data []ProjectCurrencyData `json:"data"`
}

type ProjectCurrencyData struct {
	CardBgColor         string            `json:"cardBgColor,omitempty"`
	CurrencyFullName    string            `json:"currencyFullName,omitempty"`
	CurrencyIcon        string            `json:"currencyIcon,omitempty"`
	CurrencyID          int               `json:"currencyId,omitempty"`
	CurrencyName        string            `json:"currencyName,omitempty"`
	CurrencyNum         string            `json:"currencyNum,omitempty"`
	EarnMinHold         int               `json:"earnMinHold,omitempty"`
	EarnPercentage      string            `json:"earnPercentage,omitempty"`
	EarnPercentageNum   interface{}       `json:"earnPercentageNum,omitempty"`
	IsProfitableLoan    bool              `json:"isProfitableLoan,omitempty"`
	Label               int               `json:"label,omitempty"`
	Labels              []int             `json:"labels,omitempty"`
	MarketCap           int               `json:"marketCap,omitempty"`
	NeedRegionalDisplay bool              `json:"needRegionalDisplay,omitempty"`
	ProjectList         []ProjectListItem `json:"projectList,omitempty"`
	RateNum             struct {
		Value RateNumValueWrapper `json:"value,omitempty"`
		Type  string              `json:"type,omitempty"`
	} `json:"rateNum,omitempty"`
	RateRangeMax                         string  `json:"rateRangeMax,omitempty"`
	RateRangeMaxCompar                   float64 `json:"rateRangeMaxCompar,omitempty"`
	RateRangeMaxComparatorForDefiDisplay float64 `json:"rateRangeMaxComparatorForDefiDisplay,omitempty"`
	RateRangeMin                         string  `json:"rateRangeMin,omitempty"`
	RateRangeSelectedMax                 string  `json:"rateRangeSelectedMax,omitempty"`
	RateRangeSelectedMaxNum              struct {
		Value string `json:"value,omitempty"`
		Type  string `json:"type,omitempty"`
	} `json:"rateRangeSelectedMaxNum,omitempty"`
	StockStatus int           `json:"stockStatus,omitempty"`
	SymbolSign  []interface{} `json:"symbolSign,omitempty"`
	TermAll     string        `json:"termAll,omitempty"`
	Valuation   float64       `json:"valuation,omitempty"`
}

type ProjectListItem struct {
	AdjustmentRate               int         `json:"adjustmentRate,omitempty"`
	BackTime                     interface{} `json:"backTime,omitempty"`
	Bonus                        string      `json:"bonus,omitempty"`
	BonusRate                    string      `json:"bonusRate,omitempty"`
	BonusRateNum                 interface{} `json:"bonusRateNum,omitempty"`
	BonusUnit                    string      `json:"bonusUnit,omitempty"`
	BonusUnitFullName            string      `json:"bonusUnitFullName,omitempty"`
	BonusUnitIcon                string      `json:"bonusUnitIcon,omitempty"`
	BuyingStatus                 int         `json:"buyingStatus,omitempty"`
	CapitalLinkDirect            int         `json:"capitalLinkDirect,omitempty"`
	CapitalType                  int         `json:"capitalType,omitempty"`
	CapitalTypeMsg               string      `json:"capitalTypeMsg,omitempty"`
	CardBgColor                  string      `json:"cardBgColor,omitempty"`
	CompleteHold                 float64     `json:"completeHold,omitempty"`
	CurrencyFullName             string      `json:"currencyFullName,omitempty"`
	CurrencyIcon                 string      `json:"currencyIcon,omitempty"`
	CurrencyID                   int         `json:"currencyId,omitempty"`
	DcdOptionType                string      `json:"dcdOptionType,omitempty"`
	DcdType                      string      `json:"dcdType,omitempty"`
	Description                  string      `json:"description,omitempty"`
	DisplayMultiplier            int         `json:"displayMultiplier,omitempty"`
	DisplayRate                  float64     `json:"displayRate,omitempty"`
	EndTime                      int64       `json:"endTime,omitempty"`
	EventFlag                    int         `json:"eventFlag,omitempty"`
	ExchangeRate                 float64     `json:"exchangeRate,omitempty"`
	ExpiryTime                   int         `json:"expiryTime,omitempty"`
	IsKycNeededForLite           bool        `json:"isKycNeededForLite,omitempty"`
	IsRecommendedHighRateProject int         `json:"isRecommendedHighRateProject,omitempty"`
	Issue                        int         `json:"issue,omitempty"`
	KycLevel                     int         `json:"kycLevel,omitempty"`
	KycProjectCountryCode        string      `json:"kycProjectCountryCode,omitempty"`
	KycProjectCountryName        string      `json:"kycProjectCountryName,omitempty"`
	Label                        int         `json:"label,omitempty"`
	Labels                       []int       `json:"labels,omitempty"`
	LimitCountryCode             string      `json:"limitCountryCode,omitempty"`
	LimitCountryType             int         `json:"limitCountryType,omitempty"`
	LocaleURL                    string      `json:"localeUrl,omitempty"`
	MatchCapitalLinkDirect       int         `json:"matchCapitalLinkDirect,omitempty"`
	MatchCapitalName             string      `json:"matchCapitalName,omitempty"`
	MatchCapitalType             int         `json:"matchCapitalType,omitempty"`
	MatchCurrencyFullName        string      `json:"matchCurrencyFullName,omitempty"`
	MatchCurrencyIcon            string      `json:"matchCurrencyIcon,omitempty"`
	MatchCurrencyID              int         `json:"matchCurrencyId,omitempty"`
	MaxHold                      float64     `json:"maxHold,omitempty"`
	MinHold                      float64     `json:"minHold,omitempty"`
	NeedRegionalDisplay          bool        `json:"needRegionalDisplay,omitempty"`
	NewProjectFlag               bool        `json:"newProjectFlag,omitempty"`
	NextCycleTime                interface{} `json:"nextCycleTime,omitempty"`
	OnChainTime                  interface{} `json:"onChainTime,omitempty"`
	Opened                       int         `json:"opened,omitempty"`
	PassedPurchaseCondition      int         `json:"passedPurchaseCondition,omitempty"`
	Period                       string      `json:"period,omitempty"`
	ProductStatus                bool        `json:"productStatus,omitempty"`
	ProductTypeEnum              string      `json:"productTypeEnum,omitempty"`
	ProductsType                 int         `json:"productsType,omitempty"`
	ProfitStartTime              interface{} `json:"profitStartTime,omitempty"`
	ProjectDisplayName           string      `json:"projectDisplayName,omitempty"`
	ProjectIcon                  string      `json:"projectIcon,omitempty"`
	ProjectName                  string      `json:"projectName,omitempty"`
	ProjectType                  int         `json:"projectType,omitempty"`
	Purchasable                  int         `json:"purchasable,omitempty"`
	PurchaseCondition            int         `json:"purchaseCondition,omitempty"`
	PurchaseConditionRule        string      `json:"purchaseConditionRule,omitempty"`
	Purchased                    int         `json:"purchased,omitempty"`
	Rate                         string      `json:"rate,omitempty"`
	RateNum                      struct {
		Value RateNumValueWrapper `json:"value,omitempty"`
		Type  string              `json:"type,omitempty"`
	} `json:"rateNum,omitempty"`
	RateRangeMax           string                       `json:"rateRangeMax,omitempty"`
	RateRangeMaxCompar     float64                      `json:"rateRangeMaxCompar,omitempty"`
	RateRangeMin           string                       `json:"rateRangeMin,omitempty"`
	RateThreshold          string                       `json:"rateThreshold,omitempty"`
	RecommendedProjectList []interface{}                `json:"recommendedProjectList,omitempty"`
	RedeemDay              int                          `json:"redeemDay,omitempty"`
	RedeemTime             interface{}                  `json:"redeemTime,omitempty"`
	Risk                   int                          `json:"risk,omitempty"`
	RiskMsg                string                       `json:"riskMsg,omitempty"`
	SavingBonusLimit       SavingBonusLimitValueWrapper `json:"savingBonusLimit,omitempty"`
	SavingBonusRate        string                       `json:"savingBonusRate,omitempty"`
	SavingBonusRateNum     interface{}                  `json:"savingBonusRateNum,omitempty"`
	SavingType             int                          `json:"savingType,omitempty"`
	StartTime              int64                        `json:"startTime,omitempty"`
	Status                 int                          `json:"status,omitempty"`
	Strike                 string                       `json:"strike,omitempty"`
	StxDay                 string                       `json:"stxDay,omitempty"`
	StxWeek                string                       `json:"stxWeek,omitempty"`
	SubaccountSupport      int                          `json:"subaccountSupport,omitempty"`
	SubscriptionEndTime    int64                        `json:"subscriptionEndTime,omitempty"`
	SubscriptionStartTime  int64                        `json:"subscriptionStartTime,omitempty"`
	SumHold                float64                      `json:"sumHold,omitempty"`
	SupportEntity          []interface{}                `json:"supportEntity,omitempty"`
	SupportHoldFlag        int                          `json:"supportHoldFlag,omitempty"`
	Symbol                 string                       `json:"symbol,omitempty"`
	Tags                   []interface{}                `json:"tags,omitempty"`
	TermAll                string                       `json:"termAll,omitempty"`
	Times                  int                          `json:"times,omitempty"`
	Type                   int                          `json:"type,omitempty"`
	UnderlyingIndex        string                       `json:"underlyingIndex,omitempty"`
	Unit                   string                       `json:"unit,omitempty"`
	UnitFullName           string                       `json:"unitFullName,omitempty"`
	UnitIcon               string                       `json:"unitIcon,omitempty"`
	URL                    string                       `json:"url,omitempty"`
	UserAvailableAmount    int                          `json:"userAvailableAmount,omitempty"`
	UserMaxHold            float64                      `json:"userMaxHold,omitempty"`
}

type RateNumValue []string
type RateNumValueWrapper struct {
	RateNumValue
	Text string
}

func (e *RateNumValueWrapper) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}
	// fmt.Println(string(data))
	if data[0] == '"' && data[len(data)-1] == '"' { // string?
		return json.Unmarshal(data, &e.Text)
	}
	if data[0] == '[' && data[len(data)-1] == ']' { // object?
		return json.Unmarshal(data, &e.RateNumValue)
	}
	return fmt.Errorf("unsupported error message type")
}

type SavingBonusLimitValue string
type SavingBonusLimitValueWrapper struct {
	SavingBonusLimitValue
	Data int
}

func (e *SavingBonusLimitValueWrapper) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}
	if data[0] == '"' && data[len(data)-1] == '"' { // string?
		return json.Unmarshal(data, &e.SavingBonusLimitValue)
	}
	return json.Unmarshal(data, &e.Data)
}

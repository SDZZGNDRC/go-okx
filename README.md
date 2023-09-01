# go-okx

A golang sdk for okx api.  

## RESTFul API

## WebSocket API

### Public Channel

#### Books

> 深度频道
> 参数: 全体instId  
> 数量: 1.6K  
> 爬取: √

#### EstimatedPrice

> 预估交割/行权价格频道  
> 获取永续合约，交割合约和期权预估交割/行权价。  
> 交割/行权预估价只有交割/行权前一小时开始推送预估交割/行权价，有价格变化就推送  
> 参数: instType{FUTURES|SWAP|OPTION} + instId  
> 数量: ~300  
> 爬取: √

#### FundingRate

> 资金费率频道  
> 获取永续合约资金费率，30秒到90秒内推送一次数据  
> 参数: instId(SWAP)  
> 数量: ~167  
> 爬取: √

#### IndexKline

> 指数K线频道  
> 获取指数的K线数据，推送频率最快是间隔1秒推送一次数据。  
> 参数: channel{index-candle1m|index-candle6Hutc} + 指数ID{BTC-USD|ETH-USD}  
> 数量: 2 * 2  
> 爬取: ×

#### IndexTickers

> 指数行情频道  
> 获取指数的行情数据。每100ms有变化就推送一次数据，否则一分钟推一次。  
> 参数: 指数ID  
> 数量: 2  
> 爬取: √

#### Kline

> K线频道  
> 获取K线数据，推送频率最快是间隔1秒推送一次数据。
> 参数: channel{candle1m|candle6Hutc} + instId  
> 数量: 2 * 1.6K  
> 爬取: ×

#### LiquidationOrders

> 平台公共爆仓单频道  
> 获取爆仓单信息。产品ID维度最多一秒推一条爆仓单信息。
> 参数: instType{MARGIN|SWAP|FUTURES|OPTION}  
> 数量: 4  
> 爬取: √

#### MarkPriceKline

> 标记价格K线频道  
> 获取标记价格的K线数据，推送频率最快是间隔1秒推送一次数据。
> 参数: channel{mark-price-candle1m|mark-price-candle6Hutc}  
> 数量: 2  
> 爬取: ×

#### MarkPrice

> 标记价格频道  
> 获取标记价格，标记价格有变化时，每200ms推送一次数据，标记价格没变化时，每10s推送一次数据
> 参数: instId  
> 数量: ~479  
> 爬取: √

#### OpenInterest

> 持仓总量频道  
> 获取持仓总量，每3s有数据更新推送一次数据
> 参数: {FUTURES|SWAP|OPTION} + instId  
> 数量: 1095  
> 爬取: √

#### OptDeal

> 期权公共成交频道  
> 获取最近的期权成交数据，有成交数据就推送，每次推送仅包含一条成交数据。
> 参数: instType{OPTION} + instId
> 数量: 868  
> 爬取: ×

#### OptSummary

> 期权定价频道  
> 获取所有期权合约详细定价信息，一次性推送所有
> 参数: instFamily{BTC-USD|ETH-USD} ???  
> 数量: 2 ???  
> 爬取: √

#### PriceLimit

> 限价频道  
> 获取衍生品交易的最高买价和最低卖价。限价有变化时，每5秒推送一次数据，限价没变化时，不推送数据  
> 参数: instId{SWAP|FUTURES|OPTION}
> 数量: 1065  
> 爬取: ×

#### Status

> 获取系统维护的状态，当系统维护状态改变时推送。首次订阅：”推送最新一条的变化数据“；后续每次有状态变化，推送变化的内容。

#### Tickers

> 行情频道  
> 获取产品的最新成交价、买一价、卖一价和24小时交易量等信息。  
> 最快100ms推送一次，没有触发事件时不推送，触发推送的事件有：成交、买一卖一发生变动。
> 参数: instId
> 数量: 1.6K  
> 爬取: √

#### Trades

> 交易频道  
> 获取最近的成交数据，有成交数据就推送，每次推送仅包含一条成交数据。
> 参数: instId  
> 数量: 1.6K
> 爬取: √

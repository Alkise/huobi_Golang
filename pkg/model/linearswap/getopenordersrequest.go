package linearswap

type GetOpenOrdersRequest struct {
	ContractCode string `json:"contract_code"`
	Pair         string `json:"pair"`
	PageIndex    int32  `json:"page_index"`
	PageSize     int32  `json:"page_size"`
	SortBy       string `json:"sort_by"`
	TradeType    string `json:"trade_type"`
}

package linearswap

type CrossLinearOpenOrdersResponse struct {
	Status string                             `json:"status"`
	Ts     int64                              `json:"ts"`
	Data   *CrossLinearOpenOrdersResponseData `json:"data"`
}

type CrossLinearOpenOrdersResponseData struct {
	Orders      []*Order `json:"orders"`
	TotalPage   int      `json:"total_page"`
	CurrentPage int      `json:"current_page"`
	TotalSize   int      `json:"total_size"`
}

package sqlDomain

type SummaryStockLog struct {
	Stock string `json:"stock"`
	BasicSQL
}

func (s SummaryStockLog) GetId() interface{} {
	return s.Stock
}

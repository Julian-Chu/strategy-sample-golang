package strategy_sample_go

type Shipper interface {
	CalculateFee(product Product) (float64, error)
}

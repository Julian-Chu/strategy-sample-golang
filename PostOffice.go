package strategy_sample_go

var _ Shipper = PostOffice{}

type PostOffice struct {
}

func (p PostOffice) CalculateFee(product Product) (float64, error) {
	feeByWeight := 80 + product.Weight*10
	size := product.Length * product.Width * product.Height
	feeBySize := size * 0.00002 * 1100
	return min(feeByWeight, feeBySize), nil
}

func min(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

package strategy_sample_go

type PostOffice struct {
}

func (p PostOffice) CalculateFee(product Product) (float64, error) {
	feeByWeight := 80 + product.Weight*10
	size := product.Length * product.Width * product.Height
	feeBySize := size * 0.00002 * 1100
	if feeByWeight < feeBySize {
		return feeByWeight, nil
	}
	return feeBySize, nil
}

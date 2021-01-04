package strategy_sample_go

type FedEx struct {
}

func (f FedEx) CalculateFee(product Product) (float64, error) {
	size := product.Length * product.Width * product.Height
	if product.Length > 100 || product.Width > 100 || product.Weight > 100 {
		return size*0.00002*1100 + 500, nil
	}
	return size * 0.00002 * 1200, nil
}

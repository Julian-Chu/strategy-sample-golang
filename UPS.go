package strategy_sample_go

type UPS struct {
}

func (u UPS) CalculateFee(product Product) (float64, error) {
	if product.Weight > 20 {
		return 500, nil
	}
	return 100 + product.Weight*10, nil
}

package strategy_sample_go

func Post(p *Product) (float64, error) {
	feeByWeight := 80 + p.Weight*10
	feeBySize := p.Size() * 0.00002 * 1100
	if feeByWeight < feeBySize {
		return feeByWeight, nil
	}
	return feeBySize, nil
}

func FedEx(p *Product) (float64, error) {
	if p.Length > 100 || p.Width > 100 || p.Height > 100 {
		return p.Size()*0.00002*1100 + 500, nil
	}
	return p.Size() * 0.00002 * 1200, nil
}

func UPS(p *Product) (float64, error) {
	if p.Weight > 20 {
		return 500, nil
	}
	return 100 + p.Weight*10, nil
}

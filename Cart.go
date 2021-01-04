package strategy_sample_go

import "errors"

type Cart struct{}

var (
	ErrShipperNotExist = errors.New("shipper not exist")
)

type CalcShippingFeeFunc func(*Product) (float64, error)

func (c Cart) ShippingFee(shipper string, p *Product) (float64, error) {
	switch shipper {
	case "UPS":
		return UPS(p)
	case "FedEx":
		return FedEx(p)
	case "Post office":
		return Post(p)
	default:
		return -1, ErrShipperNotExist
	}
}

func Post(p *Product) (float64, error) {
	feeByWeight := 80 + p.Weight*10
	size := p.Length * p.Width * p.Height
	feeBySize := size * 0.00002 * 1100
	if feeByWeight < feeBySize {
		return feeByWeight, nil
	}
	return feeBySize, nil
}

func FedEx(p *Product) (float64, error) {
	size := p.Length * p.Width * p.Height
	if p.Length > 100 || p.Width > 100 || p.Height > 100 {
		return size*0.00002*1100 + 500, nil
	}
	return size * 0.00002 * 1200, nil
}

func UPS(p *Product) (float64, error) {
	if p.Weight > 20 {
		return 500, nil
	}
	return 100 + p.Weight*10, nil
}

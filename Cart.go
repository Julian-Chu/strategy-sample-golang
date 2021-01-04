package strategy_sample_go

import "errors"

type Cart struct{}

var (
	ErrShipperNotExist = errors.New("shipper not exist")
)

type Product struct {
	Length float64
	Width  float64
	Height float64
	Weight float64
}

func NewProduct(length float64, width float64, height float64, weight float64) *Product {
	return &Product{Length: length, Width: width, Height: height, Weight: weight}
}

func (c Cart) ShippingFee(shipper string, p *Product) (float64, error) {
	switch shipper {
	case "UPS":
		if p.Weight > 20 {
			return 500, nil
		}
		return 100 + p.Weight*10, nil
	case "FedEx":
		size := p.Length * p.Width * p.Height
		if p.Length > 100 || p.Width > 100 || p.Height > 100 {
			return size*0.00002*1100 + 500, nil
		}
		return size * 0.00002 * 1200, nil
	case "Post office":
		feeByWeight := 80 + p.Weight*10
		size := p.Length * p.Width * p.Height
		feeBySize := size * 0.00002 * 1100
		if feeByWeight < feeBySize {
			return feeByWeight, nil
		}
		return feeBySize, nil
	default:
		return -1, ErrShipperNotExist
	}
}

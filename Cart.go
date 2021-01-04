package strategy_sample_go

import "errors"

type Cart struct{}

var (
	ErrShipperNotExist = errors.New("shipper not exist")
)

func (c Cart) ShippingFee(shipper string, length, width, height, weight float64) (float64, error) {
	switch shipper {
	case "UPS":
		if weight > 20 {
			return 500, nil
		}
		return 100 + weight*10, nil
	case "FedEx":
		size := length * width * height
		if length > 100 || width > 100 || height > 100 {
			return size*0.00002*1100 + 500, nil
		}
		return size * 0.00002 * 1200, nil
	case "Post office":
		feeByWeight := 80 + weight*10
		size := length * width * height
		feeBySize := size * 0.00002 * 1100
		if feeByWeight < feeBySize {
			return feeByWeight, nil
		}
		return feeBySize, nil
	default:
		return -1, ErrShipperNotExist
	}
}

package strategy_sample_go

import (
	"errors"
)

type Cart struct{}

var (
	ErrShipperNotExist = errors.New("shipper not exist")
)

func (c Cart) ShippingFee(shipper string, product Product) (float64, error) {
	switch shipper {
	case "UPS":
		ups := &UPS{}
		return ups.CalculateFee(product)
	case "FedEx":
		fedex := &FedEx{}
		return fedex.CalculateFee(product)
	case "Post office":
		post := &PostOffice{}
		return post.CalculateFee(product)
	default:
		return -1, ErrShipperNotExist
	}
}

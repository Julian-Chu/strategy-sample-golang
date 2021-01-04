package strategy_sample_go

import (
	"errors"
	"strings"
)

type Cart struct {
	shipperMap map[string]Shipper
}

func NewCart() *Cart {
	m := map[string]Shipper{
		"ups":         &UPS{},
		"fedex":       &FedEx{},
		"post office": &PostOffice{},
	}
	return &Cart{shipperMap: m}
}

var (
	ErrShipperNotExist = errors.New("shipper not exist")
)

func (c Cart) ShippingFee(shipper string, product Product) (float64, error) {
	shipper = strings.ToLower(shipper)
	if shipper, ok := c.shipperMap[shipper]; ok {
		return shipper.CalculateFee(product)
	}
	return -1, ErrShipperNotExist
}

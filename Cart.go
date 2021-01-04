package strategy_sample_go

import (
	"errors"
	"strings"
)

type Cart struct {
	shippingFeeMap map[string]CalcShippingFeeFunc
}

func NewCart() *Cart {
	m := map[string]CalcShippingFeeFunc{
		"ups":         UPS,
		"fedex":       FedEx,
		"post office": Post,
	}
	return &Cart{shippingFeeMap: m}
}

var (
	ErrShipperNotExist = errors.New("shipper not exist")
)

type CalcShippingFeeFunc func(*Product) (float64, error)

func (c Cart) ShippingFee(shipper string, p *Product) (float64, error) {
	if calcFee, ok := c.shippingFeeMap[strings.ToLower(shipper)]; ok {
		return calcFee(p)
	}
	return -1, ErrShipperNotExist
}

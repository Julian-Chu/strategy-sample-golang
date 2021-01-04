package strategy_sample_go

type Product struct {
	Length float64
	Width  float64
	Height float64
	Weight float64
}

func NewProduct(length, width, height, weight float64) *Product {
	return &Product{Length: length, Width: width, Height: height, Weight: weight}
}

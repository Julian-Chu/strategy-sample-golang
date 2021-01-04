package strategy_sample_go

type Product struct {
	Length float64
	Width  float64
	Height float64
	Weight float64
}

func NewProduct(length float64, width float64, height float64, weight float64) *Product {
	return &Product{Length: length, Width: width, Height: height, Weight: weight}
}

func (p Product) Size() float64 {
	return p.Length * p.Width * p.Height
}

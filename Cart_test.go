package strategy_sample_go

import "testing"

func TestCart_ShippingFee(t *testing.T) {
	ups := "UPS cat"
	fedex := "FedEx"
	postOffice := "post office"
	type args struct {
		shipper string
		length  float64
		width   float64
		height  float64
		weight  float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "ups with light weight",
			args: args{
				shipper: ups,
				length:  30,
				width:   20,
				height:  10,
				weight:  5,
			},
			want:    150,
			wantErr: false,
		},
		{
			name: "ups with heavy weight",
			args: args{
				shipper: ups,
				length:  30,
				width:   20,
				height:  10,
				weight:  50,
			},
			want:    500,
			wantErr: false,
		},
		{
			name: "fedex with small size",
			args: args{
				shipper: fedex,
				length:  30,
				width:   20,
				height:  10,
				weight:  50,
			},
			want:    144,
			wantErr: false,
		},
		{
			name: "fedex with huge size",
			args: args{
				shipper: fedex,
				length:  100,
				width:   20,
				height:  10,
				weight:  50,
			},
			want:    480,
			wantErr: false,
		},
		{
			name: "post office by weight",
			args: args{
				shipper: postOffice,
				length:  100,
				width:   20,
				height:  10,
				weight:  3,
			},
			want:    110,
			wantErr: false,
		},
		{
			name: "unknown shipper",
			args: args{
				shipper: "unknown",
				length:  0,
				width:   0,
				height:  0,
				weight:  0,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCart()
			p := Product{
				Length: tt.args.length,
				Width:  tt.args.width,
				Height: tt.args.height,
				Weight: tt.args.weight,
			}
			got, err := c.ShippingFee(tt.args.shipper, p)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShippingFee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ShippingFee() got = %v, want %v", got, tt.want)
			}
		})
	}
}

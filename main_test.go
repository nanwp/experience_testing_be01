package main

import "testing"

func TestPembayaranBarang(t *testing.T) {
	type args struct {
		ht    float64
		m     string
		cicil bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Harga Total=20.000, COD, cicil=false, error=false", args{20000, "cod", false}, false},
		{"Harga Total=20.000, COD, cicil=true, error=true", args{20000, "cod", true}, true},
		{"Harga Total=100.000, Credit, cicil=false, error=true", args{100000, "credit", false}, true},
		{"Harga Total=100.000, Gerai, cicil=false, error=false", args{100000, "gerai", false}, false},
		{"Harga Total=100.000, Gerai, cicil=true, error=true", args{100000, "gerai", true}, true},
		{"Harga Total=100.000, Credit, cicil=true, error=true", args{100000, "credit", true}, true},
		{"Harga Total=500.000, Credit, cicil=true, error=false", args{500000, "credit", true}, false},
		{"Harga Total=500.000, Credit, cicil=false, error=true", args{500000, "credit", false}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PembayaranBarang(tt.args.ht, tt.args.m, tt.args.cicil); (err != nil) != tt.wantErr {
				t.Errorf("PembayaranBarang() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

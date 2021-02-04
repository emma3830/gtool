package gtool

import "testing"

func TestRound(t *testing.T) {
	type args struct {
		f float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"四捨五入小數第2位-進位",
			args{12.256, 2},
			12.26,
		},
		{
			"四捨五入小數第2位-捨去",
			args{12.252, 2},
			12.25,
		},
		{
			"四捨五入小數第0位-進位",
			args{12.72, 0},
			13.0,
		},
		{
			"四捨五入小數第0位-捨去",
			args{12.42, 0},
			12.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.f, tt.args.n); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Round(12.256, 2)
	}
}

func TestFloorFloatNum(t *testing.T) {
	type args struct {
		f    float64
		para int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"無條件捨去到小數第2位 12.256",
			args{12.256, 2},
			12.25,
		},
		{
			"無條件捨去到小數第2位 12.252",
			args{12.252, 2},
			12.25,
		},
		{
			"無條件捨去到小數第0位 12.72",
			args{12.72, 0},
			12.0,
		},
		{
			"無條件捨去到小數第0位 12.42",
			args{12.42, 0},
			12.0,
		},
		{
			"無條件捨去到小數第2位 -12.256",
			args{-12.256, 2},
			-12.25,
		},
		{
			"無條件捨去到小數第0位 -12.42",
			args{-12.42, 0},
			-12.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloorFloatNum(tt.args.f, tt.args.para); got != tt.want {
				t.Errorf("FloorFloatNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFloorFloatNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FloorFloatNum(12.256, 2)
	}
}

func TestFloatAdd(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"小數六位計算",
			args{1.127999, 4.879333},
			6.007332,
		},
		{
			"小數十二位計算",
			args{1.333333333333, 4.333999999999},
			5.667333333332,
		},
		{
			"負數小數六位計算",
			args{-1.127999, -4.879333},
			-6.007332,
		},
		{
			"負數小數十二位計算",
			args{-1.333333333333, -4.333999999999},
			-5.667333333332,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatAdd(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("FloatAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatSub(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"小數六位計算",
			args{4.879333, 1.127999},
			3.751334,
		},
		{
			"小數十二位計算",
			args{4.333999999999, 1.333333333333},
			3.000666666666,
		},
		{
			"負數小數六位計算",
			args{-4.879333, -1.127999},
			-3.751334,
		},
		{
			"負數小數十二位計算",
			args{-4.333999999999, -1.333333333333},
			-3.000666666666,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatSub(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("FloatSub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatMul(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"小數六位計算",
			args{4.879333, 1.127999},
			5.503882744667,
		},
		{
			"小數十二位計算",
			args{4.333999999999, 1.333333333333},
			5.778666666663889,
		},
		{
			"負數小數六位計算",
			args{-4.879333, -1.127999},
			5.503882744667,
		},
		{
			"負數小數十二位計算",
			args{-4.333999999999, -1.333333333333},
			5.778666666663889,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatMul(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("FloatMul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatDiv(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"小數六位計算",
			args{4.879333, 1.127999},
			4.325653657494377,
		},
		{
			"小數十二位計算",
			args{4.333999999999, 1.333333333333},
			3.250500000000063,
		},
		{
			"負數小數六位計算",
			args{-4.879333, -1.127999},
			4.325653657494377,
		},
		{
			"負數小數十二位計算",
			args{-4.333999999999, -1.333333333333},
			3.250500000000063,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatDiv(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("FloatDiv() = %v, want %v", got, tt.want)
			}
		})
	}
}

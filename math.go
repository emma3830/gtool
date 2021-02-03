package gtool

import (
	"math"

	"github.com/shopspring/decimal"
)

// Round 四捨五入小數第n位
// ex. f = 12.256, n = 2
// 12.256 * (10 的 2 次方) >> 1225.6
// 1225.6 + 0.5 >> 1226.1
// 無條件捨去至整數 >> 1226
// 1226 / (10 的 2 次方) >> 12.26
func Round(f float64, n int) float64 {
	pow := math.Pow10(n)

	roundInt := int(FloatAdd(FloatMul(f, pow), 0.5))
	return FloatDiv(float64(roundInt), pow)
}

// FloorFloatNum 指定無條件捨去到小數第Ｘ位
func FloorFloatNum(f float64, para int) float64 {
	// 將指定的位數轉成int32
	digits := int32(para)
	sourceNum := decimal.NewFromFloat(f)

	// 位移X個位數(10進位)，取絕對值(因為可能是負數)，捨去整數後的數字，再位移回來
	resultNum := sourceNum.Shift(digits).Abs().Floor().Shift(-1 * digits)

	// 如果是負數，需要轉回來(取絕對值的時候變成了正數)
	if sourceNum.IsNegative() {
		resultNum = resultNum.Neg()
	}

	// 不需要管精準度
	resultFloat, _ := resultNum.Float64()

	return resultFloat
}

// FloatAdd float相加 x+y
func FloatAdd(x, y float64) float64 {
	sourceX := decimal.NewFromFloat(x)
	sourceY := decimal.NewFromFloat(y)

	result := sourceX.Add(sourceY)
	f, _ := result.Float64()
	return f
}

// FloatSub float相減 x-y
func FloatSub(x, y float64) float64 {
	sourceX := decimal.NewFromFloat(x)
	sourceY := decimal.NewFromFloat(y)

	result := sourceX.Sub(sourceY)
	f, _ := result.Float64()
	return f
}

// FloatMul float相乘 x*y
func FloatMul(x, y float64) float64 {
	sourceX := decimal.NewFromFloat(x)
	sourceY := decimal.NewFromFloat(y)

	result := sourceX.Mul(sourceY)
	f, _ := result.Float64()
	return f
}

// FloatDiv float相除 x/y
func FloatDiv(x, y float64) float64 {
	sourceX := decimal.NewFromFloat(x)
	sourceY := decimal.NewFromFloat(y)

	result := sourceX.Div(sourceY)
	f, _ := result.Float64()
	return f
}

package converter

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func StringValueIntoFloatWithPrecision(value string, precision int) float64 {
	parsedValue := 0.0

	if value != "" {
		f, ok := new(big.Float).SetString(value)
		if ok {
			parsedValue, _ = f.Quo(f, big.NewFloat(math.Pow10(precision))).Float64()
		}
	}

	return parsedValue
}

func Int64ValueIntoFloatWithPrecision(value int64, precision int) float64 {
	parsedValue := 0.0

	if value > 0 {
		f := new(big.Float).SetInt64(value)
		parsedValue, _ = f.Quo(f, big.NewFloat(math.Pow10(precision))).Float64()
	}

	return parsedValue
}

func StringValueIntoInt64(value string) int64 {
	parsedValue := int64(0)

	if value != "" {
		i, ok := new(big.Int).SetString(value, 10)
		if ok {
			parsedValue = i.Int64()
		}
	}

	return parsedValue
}

func StringValueIntoInt(value string) int {
	parsedValue := 0

	if value != "" {
		value, err := strconv.Atoi(value)
		if err == nil {
			parsedValue = value
		}
	}

	return parsedValue
}

func StringWithPrecisionToRawInt64Value(value string, precision int) int64 {
	parsedValue := int64(0)
	floatAmount, _, err := big.ParseFloat(value, 10, uint(precision), big.ToZero)

	if err == nil {
		scaledAmount := new(big.Float).Mul(floatAmount, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(precision)), nil)))
		intAmount := new(big.Int)
		scaledAmount.Int(intAmount)

		parsedValue = intAmount.Int64()
	}

	return parsedValue
}

func StringValueIntoFloatWithoutPrecision(value string) float64 {
	parsedValue := 0.0

	if value != "" {
		value, err := strconv.ParseFloat(value, 64)
		if err == nil {
			parsedValue = value
		}
	}

	return parsedValue
}

func StringValueIntoRawWithPrecision(value string, precision int) string {
	parsedValue := "0"

	if value != "" {
		f, ok := new(big.Float).SetString(value)
		if ok {
			parsedValueFloat, _ := f.Quo(f, big.NewFloat(math.Pow10(precision))).Float64()
			parsedValue = fmt.Sprintf("%f", parsedValueFloat)
		}
	}

	return parsedValue
}

package converter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestConvertStringValueIntoFloatWithPrecision(t *testing.T) {
	result := StringValueIntoFloatWithPrecision("2805", 6)
	assert.True(t, result > 0)
	assert.Equal(t, result, 0.002805)
}

func TestConvertInt64ValueIntoFloatWithPrecision(t *testing.T) {
	result := Int64ValueIntoFloatWithPrecision(int64(12345678901234567), 10)
	assert.True(t, result > 0)
	assert.Equal(t, result, 1234567.8901234567)
}

func TestConvertStringValueIntoint64(t *testing.T) {
	result := StringValueIntoInt64("12345678901234567")
	assert.True(t, result > 0)
	assert.Equal(t, result, int64(12345678901234567))
}

func TestConvertStringValueIntoInt(t *testing.T) {
	result := StringValueIntoInt("12345")
	assert.True(t, result > 0)
	assert.Equal(t, result, 12345)
}

func TestConvertStringWithPrecisionToRawInt64Value(t *testing.T) {
	result := StringWithPrecisionToRawInt64Value("1.8997228", 18)

	assert.NotEmpty(t, result)
	assert.Equal(t, result, int64(1899719238281250000))
}

func TestConvertStringValueIntoFloatWithoutPrecision(t *testing.T) {
	result := StringValueIntoFloatWithoutPrecision("150.25")
	assert.True(t, result > 0)
	assert.Equal(t, result, 150.25)
}

func TestStringValueIntoRawWithPrecision(t *testing.T) {
	result := StringValueIntoRawWithPrecision("440000000", 6)
	assert.True(t, result != "")
	assert.Equal(t, result, "440.000000")
}

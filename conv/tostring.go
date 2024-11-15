package conv

import (
	"strconv"
)

func Float64ToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', -1, 64)
}

func Float32ToString(input_num float32) string {
	return strconv.FormatFloat(float64(input_num), 'f', -1, 32)
}

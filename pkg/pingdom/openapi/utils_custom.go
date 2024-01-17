package openapi

// convert []int64 array to []int32
func convertInt64ArrayToInt32Array(int64Array []int64) []int32 {
	var int32Array []int32
	for _, v := range int64Array {
		int32Array = append(int32Array, int32(v))
	}
	return int32Array
}
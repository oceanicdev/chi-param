package param

import (
	"errors"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"strings"
)

// ErrInvalidParam is an error for not presented or invalid parameter
var ErrInvalidParam = errors.New("Failed to get parameter")

// String returns a path parameter as a string type
func String(r *http.Request, key string) (string, error) {
	v := chi.URLParam(r, key)
	if len(v) == 0 {
		return "", ErrInvalidParam
	}
	return v, nil
}

// Int returns a path parameter as an int type
func Int(r *http.Request, key string) (int, error) {
	return strconv.Atoi(chi.URLParam(r, key))
}

// Int8 returns a path parameter as an int8 type
func Int8(r *http.Request, key string) (int8, error) {
	value, err := strconv.ParseInt(chi.URLParam(r, key), 10, 8)
	return int8(value), err
}

// Int16 returns a path parameter as an int16 type
func Int16(r *http.Request, key string) (int16, error) {
	value, err := strconv.ParseInt(chi.URLParam(r, key), 10, 16)
	return int16(value), err
}

// Int32 returns a path parameter as an int32 type
func Int32(r *http.Request, key string) (int32, error) {
	value, err := strconv.ParseInt(chi.URLParam(r, key), 10, 32)
	return int32(value), err
}

// Int64 returns a path parameter as an int64 type
func Int64(r *http.Request, key string) (int64, error) {
	return strconv.ParseInt(chi.URLParam(r, key), 10, 64)
}

// Uint returns a path parameter as an uint type
func Uint(r *http.Request, key string) (uint, error) {
	value, err := strconv.ParseUint(chi.URLParam(r, key), 10, 32)
	return uint(value), err
}

// Uint8 returns a path parameter as an uint8 type
func Uint8(r *http.Request, key string) (uint8, error) {
	value, err := strconv.ParseUint(chi.URLParam(r, key), 10, 8)
	return uint8(value), err
}

// Uint16 returns a path parameter as an uint16 type
func Uint16(r *http.Request, key string) (uint16, error) {
	value, err := strconv.ParseUint(chi.URLParam(r, key), 10, 16)
	return uint16(value), err
}

// Uint32 returns a path parameter as an uint32 type
func Uint32(r *http.Request, key string) (uint32, error) {
	value, err := strconv.ParseUint(chi.URLParam(r, key), 10, 32)
	return uint32(value), err
}

// Uint64 returns a path parameter as an uint64 type
func Uint64(r *http.Request, key string) (uint64, error) {
	return strconv.ParseUint(chi.URLParam(r, key), 10, 64)
}

// Bool returns a path parameter as a boolean type
func Bool(r *http.Request, key string) (bool, error) {
	return strconv.ParseBool(chi.URLParam(r, key))
}

// Float32 returns a path parameter as a float32 type
func Float32(r *http.Request, key string) (float32, error) {
	value, err := strconv.ParseFloat(chi.URLParam(r, key), 32)
	return float32(value), err
}

// Float64 returns a path parameter as a float64 type
func Float64(r *http.Request, key string) (float64, error) {
	return strconv.ParseFloat(chi.URLParam(r, key), 64)
}

// QueryStringArray returns a slice of query parameters with string type
func QueryStringArray(r *http.Request, key string) ([]string, error) {
	values, ok := r.URL.Query()[key]
	if !ok {
		return nil, ErrInvalidParam
	}
	return values, nil
}

// QueryIntArray returns a slice of query parameters with int type
func QueryIntArray(r *http.Request, key string) ([]int, error) {
	values, err := QueryStringArray(r, key)
	if err != nil {
		return nil, err
	}
	out := make([]int, len(values))
	for index, value := range values {
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		out[index] = v
	}
	return out, nil
}

// QueryInt8Array returns a slice of query parameters with int8 type
func QueryInt8Array(r *http.Request, key string) ([]int8, error) {
	values, err := QueryStringArray(r, key)
	if err != nil {
		return nil, err
	}
	out := make([]int8, len(values))
	for index, value := range values {
		v, err := strconv.ParseInt(value, 10, 8)
		if err != nil {
			return nil, err
		}
		out[index] = int8(v)
	}
	return out, nil
}

// QueryInt16Array returns a slice of query parameters with int16 type
func QueryInt16Array(r *http.Request, key string) ([]int16, error) {
	values, err := QueryStringArray(r, key)
	if err != nil {
		return nil, err
	}
	out := make([]int16, len(values))
	for index, value := range values {
		v, err := strconv.ParseInt(value, 10, 16)
		if err != nil {
			return nil, err
		}
		out[index] = int16(v)
	}
	return out, nil
}

// QueryInt32Array returns a slice of query parameters with int32 type
func QueryInt32Array(r *http.Request, key string) ([]int32, error) {
	values, err := QueryStringArray(r, key)
	if err != nil {
		return nil, err
	}
	out := make([]int32, len(values))
	for index, value := range values {
		v, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return nil, err
		}
		out[index] = int32(v)
	}
	return out, nil
}

// QueryInt64Array returns a slice of query parameters with int64 type
func QueryInt64Array(r *http.Request, key string) ([]int64, error) {
	values, err := QueryStringArray(r, key)
	if err != nil {
		return nil, err
	}
	out := make([]int64, len(values))
	for index, value := range values {
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		out[index] = v
	}
	return out, nil
}

// QueryUintArray returns a slice of query parameters with uint type
func QueryUintArray(r *http.Request, key string) ([]uint, error) {
	values, err := QueryStringArray(r, key)
	if err != nil {
		return nil, err
	}
	out := make([]uint, len(values))
	for index, value := range values {
		v, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return nil, err
		}
		out[index] = uint(v)
	}
	return out, nil
}

// QueryUint8Array returns a slice of query parameters with uint8 type
func QueryUint8Array(r *http.Request, key string) ([]uint8, error) {
	values, err := QueryStringArray(r, key)
	if err != nil {
		return nil, err
	}
	out := make([]uint8, len(values))
	for index, value := range values {
		v, err := strconv.ParseUint(value, 10, 8)
		if err != nil {
			return nil, err
		}
		out[index] = uint8(v)
	}
	return out, nil
}

// QueryUint16Array returns a slice of query parameters with uint16 type
func QueryUint16Array(r *http.Request, key string) ([]uint16, error) {
	values, err := QueryStringArray(r, key)
	if err != nil {
		return nil, err
	}
	out := make([]uint16, len(values))
	for index, value := range values {
		v, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return nil, err
		}
		out[index] = uint16(v)
	}
	return out, nil
}

// QueryUint32Array returns a slice of query parameters with uint32 type
func QueryUint32Array(r *http.Request, key string) ([]uint32, error) {
	values, err := QueryStringArray(r, key)
	if err != nil {
		return nil, err
	}
	out := make([]uint32, len(values))
	for index, value := range values {
		v, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return nil, err
		}
		out[index] = uint32(v)
	}
	return out, nil
}

// QueryUint64Array returns a slice of query parameters with uint64 type
func QueryUint64Array(r *http.Request, key string) ([]uint64, error) {
	values, err := QueryStringArray(r, key)
	if err != nil {
		return nil, err
	}
	out := make([]uint64, len(values))
	for index, value := range values {
		v, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return nil, err
		}
		out[index] = v
	}
	return out, nil
}

// QueryBoolArray returns a slice of query parameters with boolean type
func QueryBoolArray(r *http.Request, key string) ([]bool, error) {
	values, err := QueryStringArray(r, key)
	if err != nil {
		return nil, err
	}
	out := make([]bool, len(values))
	for index, value := range values {
		v, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		out[index] = v
	}
	return out, nil
}

// QueryFloat32Array returns a slice of query parameters with float32 type
func QueryFloat32Array(r *http.Request, key string) ([]float32, error) {
	values, err := QueryStringArray(r, key)
	if err != nil {
		return nil, err
	}
	out := make([]float32, len(values))
	for index, value := range values {
		// replace + stripped out during url parse stage
		if strings.Contains(value, " ") {
			value = strings.Replace(value, " ", "+", 1)
		}

		v, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return nil, err
		}
		out[index] = float32(v)
	}
	return out, nil
}

// QueryFloat64Array returns a slice of query parameters with float64 type
func QueryFloat64Array(r *http.Request, key string) ([]float64, error) {
	values, err := QueryStringArray(r, key)
	if err != nil {
		return nil, err
	}
	out := make([]float64, len(values))
	for index, value := range values {
		// replace + stripped out during url parse stage
		if strings.Contains(value, " ") {
			value = strings.Replace(value, " ", "+", 1)
		}

		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		out[index] = v
	}
	return out, nil
}

// QueryString returns a query parameter with string type
func QueryString(r *http.Request, key string) (string, error) {
	values, err := QueryStringArray(r, key)
	if err != nil {
		return "", err
	}
	return values[0], nil
}

// QueryInt returns a query parameter with int type
func QueryInt(r *http.Request, key string) (int, error) {
	values, err := QueryIntArray(r, key)
	if err != nil {
		return 0, err
	}
	return values[0], nil
}

// QueryInt8 returns a query parameter with int8 type
func QueryInt8(r *http.Request, key string) (int8, error) {
	values, err := QueryInt8Array(r, key)
	if err != nil {
		return 0, err
	}
	return values[0], nil
}

// QueryInt16 returns a query parameter with int16 type
func QueryInt16(r *http.Request, key string) (int16, error) {
	values, err := QueryInt16Array(r, key)
	if err != nil {
		return 0, err
	}
	return values[0], nil
}

// QueryInt32 returns a query parameter with int32 type
func QueryInt32(r *http.Request, key string) (int32, error) {
	values, err := QueryInt32Array(r, key)
	if err != nil {
		return 0, err
	}
	return values[0], nil
}

// QueryInt64 returns a query parameter with int64 type
func QueryInt64(r *http.Request, key string) (int64, error) {
	values, err := QueryInt64Array(r, key)
	if err != nil {
		return 0, err
	}
	return values[0], nil
}

// QueryUint returns a query parameter with uint type
func QueryUint(r *http.Request, key string) (uint, error) {
	values, err := QueryUintArray(r, key)
	if err != nil {
		return 0, err
	}
	return values[0], nil
}

// QueryUint8 returns a query parameter with uint8 type
func QueryUint8(r *http.Request, key string) (uint8, error) {
	values, err := QueryUint8Array(r, key)
	if err != nil {
		return 0, err
	}
	return values[0], nil
}

// QueryUint16 returns a query parameter with uint16 type
func QueryUint16(r *http.Request, key string) (uint16, error) {
	values, err := QueryUint16Array(r, key)
	if err != nil {
		return 0, err
	}
	return values[0], nil
}

// QueryUint32 returns a query parameter with uint32 type
func QueryUint32(r *http.Request, key string) (uint32, error) {
	values, err := QueryUint32Array(r, key)
	if err != nil {
		return 0, err
	}
	return values[0], nil
}

// QueryUint64 returns a query parameter with uint64 type
func QueryUint64(r *http.Request, key string) (uint64, error) {
	values, err := QueryUint64Array(r, key)
	if err != nil {
		return 0, err
	}
	return values[0], nil
}

// QueryBool returns a query parameter with boolean type
func QueryBool(r *http.Request, key string) (bool, error) {
	values, err := QueryBoolArray(r, key)
	if err != nil {
		return false, err
	}
	return values[0], nil
}

// QueryFloat32 returns a query parameter with float32 type
func QueryFloat32(r *http.Request, key string) (float32, error) {
	values, err := QueryFloat32Array(r, key)
	if err != nil {
		return 0, err
	}
	return values[0], nil
}

// QueryFloat64 returns a query parameter with float64 type
func QueryFloat64(r *http.Request, key string) (float64, error) {
	values, err := QueryFloat64Array(r, key)
	if err != nil {
		return 0, err
	}
	return values[0], nil
}

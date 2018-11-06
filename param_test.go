package param

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"math"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

func newParamRequest(t *testing.T, value string) (*http.Request, string) {
	t.Helper()

	paramKey := "chiRocks"
	// create new chi route context
	c := chi.NewRouteContext()

	c.URLParams.Add(paramKey, value)

	// create new request and append chi context
	r := httptest.NewRequest("GET", "/", nil)
	ctx := context.WithValue(r.Context(), chi.RouteCtxKey, c)

	return r.WithContext(ctx), paramKey
}

func newQueryRequest(t *testing.T, queryString string) *http.Request {
	t.Helper()

	r := httptest.NewRequest("GET", "/?"+queryString, nil)

	return r
}

func TestString(t *testing.T) {
	want := "bar"
	req, key := newParamRequest(t, want)

	got, err := String(req, key)

	if err != nil {
		t.Fatal("failed to parse string argument: ", err)
	}

	if got != want {
		t.Fatalf("got '%s', want '%s'", got, want)
	}
}

func TestStringErr(t *testing.T) {
	req, _ := newParamRequest(t, "whoops")

	_, err := String(req, "fail")

	if err == nil {
		t.Fatal("String failed to return error for non existent param")
	}
}

func TestInt(t *testing.T) {
	want := 2301342
	maxValue := strconv.Itoa(want)

	req, key := newParamRequest(t, maxValue)
	got, err := Int(req, key)

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestIntErr(t *testing.T) {
	req, key := newParamRequest(t, "ten")
	_, err := Int(req, key)

	if err == nil {
		t.Fatal("Int('ten') should throw an error")
	}
}

func TestInt8(t *testing.T) {
	var want int8 = math.MaxInt8
	maxValue := strconv.Itoa(math.MaxInt8)

	req, key := newParamRequest(t, maxValue)
	got, err := Int8(req, key)

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestInt8Err(t *testing.T) {
	req, key := newParamRequest(t, "ten")
	_, err := Int8(req, key)

	if err == nil {
		t.Fatal("Int8('ten') should throw an error")
	}
}

func TestInt16(t *testing.T) {
	var want int16 = math.MaxInt16
	maxValue := strconv.Itoa(math.MaxInt16)

	req, key := newParamRequest(t, maxValue)
	got, err := Int16(req, key)

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestInt16Err(t *testing.T) {
	req, key := newParamRequest(t, "ten")
	_, err := Int16(req, key)

	if err == nil {
		t.Fatal("Int16('ten') should throw an error")
	}
}

func TestInt32(t *testing.T) {
	var want int32 = math.MaxInt32
	maxValue := strconv.Itoa(math.MaxInt32)

	req, key := newParamRequest(t, maxValue)
	got, err := Int32(req, key)

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestInt32Err(t *testing.T) {
	req, key:= newParamRequest(t, "ten")
	_, err := Int32(req, key)

	if err == nil {
		t.Fatal("Int32('ten') should throw an error")
	}
}

func TestInt64(t *testing.T) {
	var want int64 = math.MaxInt64
	maxValue := strconv.Itoa(math.MaxInt64)

	req, key := newParamRequest(t, maxValue)
	got, err := Int64(req, key)

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestInt64Err(t *testing.T) {
	req, key := newParamRequest(t, "ten")
	_, err := Int64(req, key)

	if err == nil {
		t.Fatal("Int64('ten') should throw an error")
	}
}

func TestUInt(t *testing.T) {
	var want uint = math.MaxUint32
	maxValue := strconv.FormatUint(uint64(want), 10)

	req, key := newParamRequest(t, maxValue)
	got, err := Uint(req, key)

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestUIntErr(t *testing.T) {
	req, key := newParamRequest(t, "ten")
	_, err := Uint(req, key)

	if err == nil {
		t.Fatal("Uint('ten') should throw an error")
	}
}

func TestUInt8(t *testing.T) {
	var want uint8 = math.MaxUint8
	maxValue := strconv.FormatUint(uint64(want), 10)

	req, key := newParamRequest(t, maxValue)
	got, err := Uint8(req, key)

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestUInt8Err(t *testing.T) {
	req, key := newParamRequest(t, "ten")
	_, err := Uint8(req, key)

	if err == nil {
		t.Fatal("Uint8('ten') should throw an error")
	}
}

func TestUInt16(t *testing.T) {
	var want uint16 = math.MaxUint16
	maxValue := strconv.FormatUint(uint64(want), 10)

	req, key := newParamRequest(t, maxValue)
	got, err := Uint16(req, key)

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestUInt16Err(t *testing.T) {
	req, key := newParamRequest(t, "ten")
	_, err := Uint16(req, key)

	if err == nil {
		t.Fatal("Uint16('ten') should throw an error")
	}
}

func TestUInt32(t *testing.T) {
	var want uint32 = math.MaxUint32
	maxValue := strconv.FormatUint(uint64(want), 10)

	req, key := newParamRequest(t, maxValue)
	got, err := Uint32(req, key)

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestUInt32Err(t *testing.T) {
	req, key := newParamRequest(t, "ten")
	_, err := Uint32(req, key)

	if err == nil {
		t.Fatal("Uint32('ten') should throw an error")
	}
}

func TestUInt64(t *testing.T) {
	var want uint64 = math.MaxUint64
	maxValue := strconv.FormatUint(uint64(want), 10)

	req, key := newParamRequest(t, maxValue)
	got, err := Uint64(req, key)

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestUInt64Err(t *testing.T) {
	req, key := newParamRequest(t, "ten")
	_, err := Uint64(req, key)

	if err == nil {
		t.Fatal("Uint64('ten') should throw an error")
	}
}

func TestBool(t *testing.T) {
	var want bool = true

	req, key := newParamRequest(t, "true")
	got, err := Bool(req, key)

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestBoolErr(t *testing.T) {
	req, key := newParamRequest(t, "truth!")
	_, err := Bool(req, key)

	if err == nil {
		t.Fatal("Bool('truth!') should throw an error")
	}
}

func TestFloat32(t *testing.T) {
	var want float32 = math.MaxFloat32
	maxValue := strconv.FormatFloat(math.MaxFloat32, 'E', -1, 32)

	req, key := newParamRequest(t, maxValue)
	got, err := Float32(req, key)

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestFloat32Err(t *testing.T) {
	req, key := newParamRequest(t, "pi")
	_, err := Float32(req, key)

	if err == nil {
		t.Fatal("Float32('pi') should throw an error")
	}
}

func TestFloat64(t *testing.T) {
	var want = math.MaxFloat64
	maxValue := strconv.FormatFloat(math.MaxFloat64, 'E', -1, 64)

	req, key := newParamRequest(t, maxValue)
	got, err := Float64(req, key)

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestFloat64Err(t *testing.T) {
	req, key := newParamRequest(t, "pi")
	_, err := Float64(req, key)

	if err == nil {
		t.Fatal("Float64('pi') should throw an error")
	}
}

func TestQueryStringArray(t *testing.T) {
	req := newQueryRequest(t, "fruit=apple&fruit=orange&veggie=pepper")

	got, err := QueryStringArray(req, "fruit")
	if err != nil {
		t.Fatal(err)
	}

	want := []string{"apple", "orange"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryStringArrayErr(t *testing.T) {
	req := newQueryRequest(t, "fruit=apple")

	_, err := QueryStringArray(req, "veggie")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

}

func TestQueryIntArray(t *testing.T) {
	req := newQueryRequest(t, "age=23&age=42&name=oliver")

	got, err := QueryIntArray(req, "age")
	if err != nil {
		t.Fatal(err)
	}

	want := []int{23, 42}
	if  !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryIntArrayErr(t *testing.T) {
	req := newQueryRequest(t, "age=twelve")

	_, err := QueryIntArray(req, "fruit")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

	_, err = QueryIntArray(req, "age")
	if err == nil {
		t.Fatal("expected error trying to parse string to int")
	}
}

func TestQueryInt8Array(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("num=%d&num=%d", math.MinInt8, math.MaxInt8))

	got, err := QueryInt8Array(req, "num")
	if err != nil {
		t.Fatal(err)
	}

	want := []int8{math.MinInt8, math.MaxInt8}
	if  !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryInt8ArrayErr(t *testing.T) {

	req := newQueryRequest(t, fmt.Sprintf("age=%d", math.MaxInt16))

	_, err := QueryInt8Array(req, "height")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

	_, err = QueryInt8Array(req, "age")
	if err == nil {
		t.Fatal("expected error trying to parse int8")
	}
}

func TestQueryInt16Array(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("num=%d&num=%d", math.MinInt16, math.MaxInt16))

	got, err := QueryInt16Array(req, "num")
	if err != nil {
		t.Fatal(err)
	}

	want := []int16{math.MinInt16, math.MaxInt16}
	if  !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryInt16ArrayErr(t *testing.T) {

	req := newQueryRequest(t, fmt.Sprintf("age=%d", math.MaxInt32))

	_, err := QueryInt16Array(req, "height")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

	_, err = QueryInt16Array(req, "age")
	if err == nil {
		t.Fatal("expected error trying to parse int16")
	}
}

func TestQueryInt32Array(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("num=%d&num=%d", math.MinInt32, math.MaxInt32))

	got, err := QueryInt32Array(req, "num")
	if err != nil {
		t.Fatal(err)
	}

	want := []int32{math.MinInt32, math.MaxInt32}
	if  !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryInt32ArrayErr(t *testing.T) {

	req := newQueryRequest(t, fmt.Sprintf("age=%d", math.MaxInt64))

	_, err := QueryInt32Array(req, "height")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

	_, err = QueryInt32Array(req, "age")
	if err == nil {
		t.Fatal("expected error trying to parse int32")
	}
}

func TestQueryInt64Array(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("num=%d&num=%d", math.MinInt64, math.MaxInt64))

	got, err := QueryInt64Array(req, "num")
	if err != nil {
		t.Fatal(err)
	}

	want := []int64{math.MinInt64, math.MaxInt64}
	if  !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryInt64ArrayErr(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%s", "notanint"))

	_, err := QueryInt64Array(req, "height")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

	_, err = QueryInt64Array(req, "age")
	if err == nil {
		t.Fatal("expected error trying to parse int64")
	}
}

func TestQueryUIntArray(t *testing.T) {
	var max uint = math.MaxUint32
	req := newQueryRequest(t, fmt.Sprintf("num=%d&num=%d", 32, max))

	got, err := QueryUintArray(req, "num")
	if err != nil {
		t.Fatal(err)
	}

	want := []uint{32, math.MaxUint32}
	if  !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryUIntArrayErr(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%s", "notanint"))

	_, err := QueryUintArray(req, "height")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

	_, err = QueryUintArray(req, "age")
	if err == nil {
		t.Fatal("expected error trying to parse uint")
	}
}

func TestQueryUInt8Array(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("num=%d&num=%d", 32, math.MaxUint8))

	got, err := QueryUint8Array(req, "num")
	if err != nil {
		t.Fatal(err)
	}

	want := []uint8{32, math.MaxUint8}
	if  !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryUInt8ArrayErr(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%s", "notanint"))

	_, err := QueryUint8Array(req, "height")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

	_, err = QueryUint8Array(req, "age")
	if err == nil {
		t.Fatal("expected error trying to parse uint8")
	}
}

func TestQueryUInt16Array(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("num=%d&num=%d", 32, math.MaxUint16))

	got, err := QueryUint16Array(req, "num")
	if err != nil {
		t.Fatal(err)
	}

	want := []uint16{32, math.MaxUint16}
	if  !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryUInt16ArrayErr(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%s", "notanint"))

	_, err := QueryUint16Array(req, "height")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

	_, err = QueryUint16Array(req, "age")
	if err == nil {
		t.Fatal("expected error trying to parse uint16")
	}
}

func TestQueryUInt32Array(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("num=%d&num=%d", 32, math.MaxUint32))

	got, err := QueryUint32Array(req, "num")
	if err != nil {
		t.Fatal(err)
	}

	want := []uint32{32, math.MaxUint32}
	if  !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryUInt32ArrayErr(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%s", "notanint"))

	_, err := QueryUint32Array(req, "height")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

	_, err = QueryUint32Array(req, "age")
	if err == nil {
		t.Fatal("expected error trying to parse uint64")
	}
}

func TestQueryUInt64Array(t *testing.T) {
	var max uint64 = math.MaxUint64
	req := newQueryRequest(t, fmt.Sprintf("num=%d&num=%d", 32, max))

	got, err := QueryUint64Array(req, "num")
	if err != nil {
		t.Fatal(err)
	}

	want := []uint64{32, math.MaxUint64}
	if  !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryUInt64ArrayErr(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%s", "notanint"))

	_, err := QueryUint64Array(req, "height")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

	_, err = QueryUint64Array(req, "age")
	if err == nil {
		t.Fatal("expected error trying to parse uint64")
	}
}

func TestQueryBoolArray(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("options=%v&options=%v", true, false))

	got, err := QueryBoolArray(req, "options")
	if err != nil {
		t.Fatal(err)
	}

	want := []bool{true, false}
	if  !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryBoolArrayErr(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("options=%s", "truth"))

	_, err := QueryBoolArray(req, "opts")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

	_, err = QueryBoolArray(req, "options")
	if err == nil {
		t.Fatal("expected error trying to parse bool")
	}
}

func TestQueryFloat32Array(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("result=%G&result=%G", math.SmallestNonzeroFloat32, math.MaxFloat32))

	got, err := QueryFloat32Array(req, "result")
	if err != nil {
		t.Fatal(err)
	}

	want := []float32{math.SmallestNonzeroFloat32, math.MaxFloat32}
	if  !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryFloat32ArrayErr(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("result=%s", "notafloat"))

	_, err := QueryFloat32Array(req, "opts")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

	_, err = QueryFloat32Array(req, "result")
	if err == nil {
		t.Fatal("expected error trying to parse float32")
	}
}

func TestQueryFloat64Array(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("result=%G&result=%G", math.SmallestNonzeroFloat64, math.MaxFloat64))

	got, err := QueryFloat64Array(req, "result")
	if err != nil {
		t.Fatal(err)
	}

	want := []float64{math.SmallestNonzeroFloat64, math.MaxFloat64}
	if  !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryFloat64ArrayErr(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("result=%s", "notafloat"))

	_, err := QueryFloat64Array(req, "opts")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

	_, err = QueryFloat64Array(req, "result")
	if err == nil {
		t.Fatal("expected error trying to parse float64")
	}
}

func TestQueryString(t *testing.T) {
	req := newQueryRequest(t, "fruit=apple")

	got, err := QueryString(req, "fruit")
	if err != nil {
		t.Fatal(err)
	}

	want := "apple"
	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryStringErr(t *testing.T) {
	req := newQueryRequest(t, "fruit=apple")

	_, err := QueryString(req, "veggie")
	if err == nil {
		t.Fatal("expected error for invalid param")
	}

}

func TestQueryInt(t *testing.T) {
	want := 24234234
	req := newQueryRequest(t, fmt.Sprintf("age=%d", want))

	got, err := QueryInt(req, "age")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryIntErr(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%e", math.MaxFloat32))

	_, err := QueryInt(req, "age")
	if err == nil {
		t.Fatal("expected error parsing invalid int")
	}

}

func TestQueryInt8(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%d", math.MaxInt8))

	got, err := QueryInt8(req, "age")
	if err != nil {
		t.Fatal(err)
	}

	want := int8(math.MaxInt8)
	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryInt8Err(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%d", math.MaxInt16))

	_, err := QueryInt8(req, "age")
	if err == nil {
		t.Fatal("expected error parsing invalid int8")
	}

}

func TestQueryInt16(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%d", math.MaxInt16))

	got, err := QueryInt16(req, "age")
	if err != nil {
		t.Fatal(err)
	}

	want := int16(math.MaxInt16)
	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryInt16Err(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%d", math.MaxInt32))

	_, err := QueryInt16(req, "age")
	if err == nil {
		t.Fatal("expected error parsing invalid int16")
	}
}

func TestQueryInt32(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%d", math.MaxInt32))

	got, err := QueryInt32(req, "age")
	if err != nil {
		t.Fatal(err)
	}

	want := int32(math.MaxInt32)
	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryInt32Err(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%d", math.MaxInt64))

	_, err := QueryInt32(req, "age")
	if err == nil {
		t.Fatal("expected error parsing invalid int32")
	}
}

func TestQueryInt64(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%d", math.MaxInt64))

	got, err := QueryInt64(req, "age")
	if err != nil {
		t.Fatal(err)
	}

	want := int64(math.MaxInt64)
	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryInt64Err(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%d1", math.MaxInt64))

	_, err := QueryInt64(req, "age")
	if err == nil {
		t.Fatal("expected error parsing invalid int64")
	}
}

func TestQueryUint(t *testing.T) {
	want := uint(math.MaxUint32)
	req := newQueryRequest(t, fmt.Sprintf("age=%d", want))

	got, err := QueryUint(req, "age")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryUintErr(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%d", uint(math.MaxUint64)))

	_, err := QueryUint(req, "age")
	if err == nil {
		t.Fatal("expected error parsing invalid uint")
	}
}

func TestQueryUint8(t *testing.T) {
	want := uint8(math.MaxUint8)
	req := newQueryRequest(t, fmt.Sprintf("age=%d", want))

	got, err := QueryUint8(req, "age")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryUint8Err(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%d", uint(math.MaxUint64)))

	_, err := QueryUint8(req, "age")
	if err == nil {
		t.Fatal("expected error parsing invalid uint8")
	}
}

func TestQueryUint16(t *testing.T) {
	want := uint16(math.MaxUint16)
	req := newQueryRequest(t, fmt.Sprintf("age=%d", want))

	got, err := QueryUint16(req, "age")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryUint16Err(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%d", uint32(math.MaxUint32)))

	_, err := QueryUint16(req, "age")
	if err == nil {
		t.Fatal("expected error parsing invalid uint16")
	}
}

func TestQueryUint32(t *testing.T) {
	want := uint32(math.MaxUint32)
	req := newQueryRequest(t, fmt.Sprintf("age=%d", want))

	got, err := QueryUint32(req, "age")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryUint32Err(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%d", uint64(math.MaxUint64)))

	_, err := QueryUint32(req, "age")
	if err == nil {
		t.Fatal("expected error parsing invalid uin32")
	}
}

func TestQueryUint64(t *testing.T) {
	want := uint64(math.MaxUint64)
	req := newQueryRequest(t, fmt.Sprintf("age=%d", want))

	got, err := QueryUint64(req, "age")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryUint64Err(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("age=%d", -1))

	_, err := QueryUint64(req, "age")
	if err == nil {
		t.Fatal("expected error parsing invalid uint64")
	}
}

func TestQueryBool(t *testing.T) {
	want := true
	req := newQueryRequest(t, fmt.Sprintf("active=%s", "true"))

	got, err := QueryBool(req, "active")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryBoolErr(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("active=%s", "truth"))

	_, err := QueryBool(req, "active")
	if err == nil {
		t.Fatal("expected error parsing invalid bool")
	}
}

func TestQueryFloat32(t *testing.T) {
	want := float32(math.MaxFloat32)
	req := newQueryRequest(t, fmt.Sprintf("value=%G", want))

	got, err := QueryFloat32(req, "value")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryFloat32Err(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("value=%s", "notafloat"))

	_, err := QueryFloat32(req, "value")
	if err == nil {
		t.Fatal("expected error parsing invalid float32")
	}
}

func TestQueryFloat64(t *testing.T) {
	want := float64(math.MaxFloat64)
	req := newQueryRequest(t, fmt.Sprintf("value=%G", want))

	got, err := QueryFloat64(req, "value")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueryFloat64Err(t *testing.T) {
	req := newQueryRequest(t, fmt.Sprintf("value=%s", "notafloat"))

	_, err := QueryFloat64(req, "value")
	if err == nil {
		t.Fatal("expected error parsing invalid float64")
	}
}
